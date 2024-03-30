package signer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"telegraph/config"
	"telegraph/log"
	"telegraph/tools"

	"telegraph/pkg/wallet"

	"github.com/ethereum/go-ethereum/common"

	networkmodule "telegraph/modules/network"
	transactionmodule "telegraph/modules/transaction"
	validatormodule "telegraph/modules/validator"
	"telegraph/pkg/db/validator"
	"telegraph/pkg/signer/auth"
)

type CrossChainData struct {
	Addresses []common.Address
	Numbers   []*big.Int
	Strings   []string
	Bools     []bool
}

type Signer interface {
	SignTransaction(t transactionmodule.Transaction)
	signEVM(t transactionmodule.Transaction, publicKey common.Address)
	evmThresholdCheck(hashedMesg common.Hash, signatureBytes []byte, networkName string) *big.Int
}

type SignerImpl struct {
	auth               auth.IAuth
	wallet             wallet.Wallet
	transactionService transactionmodule.TransactionService
	networkService     networkmodule.NetworkService
}

func (signerInstance *SignerImpl) SignTransaction(t transactionmodule.Transaction) {
	pubAddress, err := signerInstance.wallet.GetPublicAddress()
	tools.Check(err, "wallet.GetPublicKey")
	signerInstance.signEVM(t, pubAddress)
}

func signWithKeyServer(t transactionmodule.Transaction, EVMHTTPURL string, contractAddress string) ([]byte, common.Hash, error) {
	// put tx in request body
	requestUrl := "http://localhost:7045/signer/request-signature"

	// Convert the transaction object to a JSON string
	tJson, err := json.Marshal(t)
	if err != nil {
		return nil, common.Hash{}, err
	}

	// the expected request body is a JSON object with the following fields:
	// type requestSignatureHandlerRequest struct {
	// 	Transaction     Transaction `json:"transaction"`
	// 	EVMHTTPURL      string      `json:"EVMHTTPURL"`
	// 	ContractAddress string      `json:"contractAddress"`
	// }
	requestBody := []byte(fmt.Sprintf(`{"transaction":%s,"EVMHTTPURL":"%s","contractAddress":"%s"}`, string(tJson), EVMHTTPURL, contractAddress))

	// Create a new request using http
	request, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, common.Hash{}, err
	}
	request.Header.Add("Content-Type", "application/json")

	// Send the request using an http.Client
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, common.Hash{}, err
	}
	defer response.Body.Close()

	// Check if the HTTP request was successful
	if response.StatusCode != http.StatusOK {
		return nil, common.Hash{}, fmt.Errorf("server returned non-200 status: %d", response.StatusCode)
	}

	// Read the response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, common.Hash{}, err
	}

	// this is the response data: json.NewEncoder(w).Encode(map[string]string{"signature": fmt.Sprintf("0x%x", signature), "hashedMessage": hashedMesg.String()})
	// it has the signature and the hashed message(common.Hash)
	// it should return both the signature and the hashed message
	var responseMap map[string]string
	err = json.Unmarshal(responseBody, &responseMap)
	if err != nil {
		return nil, common.Hash{}, err
	}

	// convert the signature to bytes
	signature := responseMap["signature"]
	signatureBytes := common.FromHex(signature)

	// convert the hashed message to common.Hash
	hashedMessage := responseMap["hashedMessage"]
	hashedMesg := common.HexToHash(hashedMessage)

	return signatureBytes, hashedMesg, nil
}

func (signerInstance *SignerImpl) signEVM(t transactionmodule.Transaction, publicAddress common.Address) {
	endChainNetwork, err := signerInstance.networkService.GetNetworkByName(t.EndChain)
	tools.Check(err, "getNetworkByName")
	signatureBytes, hashedMesg, err := signWithKeyServer(t, endChainNetwork.EVMHTTPURL, endChainNetwork.ContractAddress)
	tools.Check(err, "signatureBytes")

	r := [32]byte{}
	s := [32]byte{}
	copy(r[:], signatureBytes[:32])
	copy(s[:], signatureBytes[32:64])
	v := uint8(int(signatureBytes[64])) + 27

	trx := transactionmodule.Transaction{}
	pubAddress, err := signerInstance.wallet.GetPublicAddress()
	tools.Check(err, "wallet.GetPublicKey")
	trx.Signers = append(trx.Signers, pubAddress.Bytes())
	trx.H = append(trx.H, hashedMesg)
	trx.R = append(trx.R, r)
	trx.S = append(trx.S, s)
	trx.V = append(trx.V, v)
	fmt.Println("About to Sign", trx.SignerTime)
	if trx.SignerTime == nil {
		// set to 0
		trx.SignerTime = new(big.Int)
	}
	fmt.Println("About to Sign", trx.SignerTime)
	trx.SignerTime.Add(trx.SignerTime, signerInstance.evmThresholdCheck(hashedMesg, signatureBytes, t.EndChain))
	fmt.Println("NEW SIGNER TIME", trx.SignerTime)
	trx.Hash = t.Hash

	// save signed hash and (r,s,v) values
	err = signerInstance.transactionService.UpdateTransaction(trx)
	tools.Check(err)

	// After a successful signing, the node should send its R, S,
	// and corresponding Hash to all nodes within its Validator list
	updateBytes, err := json.Marshal(trx)
	tools.Check(err)
	validators := validator.GetValidators()
	// sign auth message
	for _, v := range validators {
		ip := v.Domain
		if ip != config.Conf.IP { // skip self and send to others
			log.Debug("Sending transaction update to ip: ", ip)

			resp, err := tools.API("PUT", ip+"/transaction/update", updateBytes)
			if tools.Check(err) {
				log.Debug("Status ", resp.StatusCode, " from /transaction/update request to ", ip)
			}
		}
	}
}

func (signerInstance *SignerImpl) evmThresholdCheck(hashedMesg common.Hash, signatureBytes []byte, networkName string) *big.Int {
	endChainNetwork, err := signerInstance.networkService.GetNetworkByName(networkName)
	tools.Check(err, "getNetworkByName")

	// call the key-server to get the signer time
	requestUrl := "http://localhost:7045/signer/get-signer-time"
	// the expected request body is a JSON object with the following fields:
	// type getSignerTimeHandlerRequest struct {
	// 	HashedMesg      common.Hash `json:"hashedMesg"`
	// 	SignatureBytes  []byte      `json:"signatureBytes"`
	// 	EVMHTTPURL      string      `json:"EVMHTTPURL"`
	// 	ContractAddress string      `json:"contractAddress"`
	// }
	requestBody := []byte(fmt.Sprintf(`{"hashedMesg":"%s","signatureBytes":"%s","EVMHTTPURL":"%s","contractAddress":"%s"}`, hashedMesg, signatureBytes, endChainNetwork.EVMHTTPURL, endChainNetwork.ContractAddress))

	// Create a new request using http
	request, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil
	}
	request.Header.Add("Content-Type", "application/json")

	// Send the request using an http.Client
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil
	}
	defer response.Body.Close()

	// Check if the HTTP request was successful
	if response.StatusCode != http.StatusOK {
		return nil
	}

	// Read the response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil
	}

	// this is the response data: json.NewEncoder(w).Encode(map[string]string{"signerTime": fmt.Sprintf("%d", signerTime)})
	// it has the signer time
	// it should return the signer time
	var responseMap map[string]string
	err = json.Unmarshal(responseBody, &responseMap)
	if err != nil {
		return nil
	}

	// convert the signer time to big.Int
	signerTime := responseMap["signerTime"]
	signerTimeBigInt := new(big.Int)
	signerTimeBigInt.SetString(signerTime, 10)

	return signerTimeBigInt
}

func NewSigner(auth auth.IAuth, transactionService transactionmodule.TransactionService, networkService networkmodule.NetworkService, validatorService validatormodule.ValidatorService) Signer {
	wallet := wallet.NewWallet(validatorService)
	return &SignerImpl{
		auth:               auth,
		wallet:             wallet,
		transactionService: transactionService,
		networkService:     networkService,
	}
}
