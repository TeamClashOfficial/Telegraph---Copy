// TODO: Make the wallet more secure
// This is the first version of the wallet with insecure storage of wallet
package wallet

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"telegraph/log"
	validatormodule "telegraph/modules/validator"
	"telegraph/tools"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	dataDir    = "data"
	walletFile = "wallet.data"
)

var ckaID = []byte("telegraph")

type Wallet interface {
	// GetWallet() Wallet
	GetPublicAddress() (common.Address, error)
	SignMessageWithPrivateKey(t common.Hash) string
	VerifySignature(signature string) bool
}

type WalletImpl struct {
	validatorService validatormodule.ValidatorService
}

// http POST call to /signer/request-signature on localhost:7045
func requestSignature(t common.Hash) []byte {
	requestUrl := "http://localhost:7045/signer/request-signature"
	requestBody := t.Bytes() // Correct way to convert to []byte

	// Create a new request using http
	request, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err) // Handle the error appropriately
	}
	request.Header.Add("Content-Type", "application/json")

	// Send the request using an http.Client
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err) // Handle the error appropriately
	}
	defer response.Body.Close()

	// Here you would typically read the response and return the appropriate result
	// Assuming the server is sending back the signature as bytes in the response body.
	// But since we're not really processing the response here, it returns nil.
	// You should process the response body and return the actual result needed.

	return nil
}

// http POST call to /signer/initialize on localhost:7045
func initializeWallet() {
	requestUrl := "http://localhost:7045/signer/initialize"
	requestBody := []byte("")
	// Create a new request using http
	request, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err) // Handle the error appropriately
	}
	request.Header.Add("Content-Type", "application/json")

	// Send the request using an http.Client
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err) // Handle the error appropriately
	}
	defer response.Body.Close()

	// Here you would typically read the response and return the appropriate result
}

func base64Decode(data string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, fmt.Errorf("failed to base64 decode: %v", err)
	}

	return decoded, nil
}

// getPublicKey retrieves the public key associated with the wallet from the server and unmarshals it into an ecdsa.PublicKey.
func getPublicAddress() (common.Address, error) {
	requestUrl := "http://localhost:7045/signer/public-key"
	request, err := http.NewRequest("GET", requestUrl, nil) // No need for a body in a GET request
	if err != nil {
		return common.Address{}, err
	}
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return common.Address{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return common.Address{}, errors.New("server responded with an error")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return common.Address{}, err
	}

	var result struct {
		Address string `json:"address"`
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return common.Address{}, err
	}

	// Convert the address string back to address format
	address := common.HexToAddress(result.Address)

	return address, nil
}

// func (wallet *WalletImpl) GetWallet() Wallet {
// 	signer, err := hsm.HSMCtx.FindKeyPair(ckaID, nil)

// 	if signer == nil || err != nil {
// 		log.Error("Error loading wallet private key:", err)
// 		return nil
// 	}

// 	return &WalletImpl{
// 		validatorService: wallet.validatorService,
// 	}
// }

func (wallet *WalletImpl) SignMessageWithPrivateKey(t common.Hash) string {
	// get signature from request-signature
	signature := requestSignature(t)
	return string(signature)
}

func (wallet *WalletImpl) VerifySignature(signature string) bool {
	hashBytes := []byte("Authorization")
	hash := crypto.Keccak256(hashBytes)
	pubKey, err := crypto.SigToPub(hash, []byte(signature))
	if err != nil {
		log.Fatal(err)
	}
	tools.Check(err)
	// pubKey to address
	address := crypto.PubkeyToAddress(*pubKey)
	// check if public key exists in validator collection
	validator, err := wallet.validatorService.FindValidatorByAddress(address)
	if err != nil {
		return false
	}
	if validator.PublicETHAddress == address {
		return true
	}
	return false
}

func (wallet *WalletImpl) GetPublicAddress() (common.Address, error) {
	// get public key from key server
	pubAddress, err := getPublicAddress()
	if err != nil {
		log.Error(err)
		return common.Address{}, err
	}
	return pubAddress, nil
}

func NewWallet(validatorService validatormodule.ValidatorService) Wallet {

	return &WalletImpl{
		validatorService: validatorService,
	}
}
