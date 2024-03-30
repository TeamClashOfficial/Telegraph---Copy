package transactionmodule

import (
	"context"
	"encoding/json"
	"fmt"
	"telegraph/config"
	"telegraph/log"
	networkmodule "telegraph/modules/network"
	validatormodule "telegraph/modules/validator"
	"telegraph/pkg/signer/auth"
	"telegraph/tools"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TransactionService interface {
	GetTransactions() ([]Transaction, error)
	CreateFromBytes(data []byte) error
	UpdateFromBytes(data []byte) error
	FinalizeFromBytes(data []byte) error
	AddTransaction(transaction Transaction) error
	UpdateTransaction(transaction Transaction) error
	sendTransaction(trx Transaction) error
}

type TransactionServiceImpl struct {
	transactionRepo TransactionRepository
	auth            auth.IAuth
	// wallet           wallet.Wallet
	networkService   networkmodule.NetworkService
	validatorService validatormodule.ValidatorService
}

func (service *TransactionServiceImpl) GetTransactions() ([]Transaction, error) {
	transactions, err := service.transactionRepo.FindTransactions()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (service *TransactionServiceImpl) CreateFromBytes(data []byte) error {
	trx := Transaction{}
	err := json.Unmarshal(data, &trx)
	if err != nil {
		return err
	}
	return service.AddTransaction(trx)
}

func (service *TransactionServiceImpl) UpdateFromBytes(data []byte) error {
	trx := Transaction{}
	err := json.Unmarshal(data, &trx)
	if err != nil {
		return err
	}
	return service.UpdateTransaction(trx)
}

type FinalizeTransactionHandlerRequest struct {
	StartHash string `json:"startHash"`
	EndHash   string `json:"endHash"`
}

func (service *TransactionServiceImpl) FinalizeFromBytes(data []byte) error {
	// data should be a JSON object with the following fields:
	// type sendTransactionHandlerRequest struct {
	// 	startHash string `json:"startHash"`
	// 	endHash   string `json:"endHash"`
	// }

	request := FinalizeTransactionHandlerRequest{}
	err := json.Unmarshal(data, &request)
	if err != nil {
		return err
	}

	//convert startHash and endHash to common.Hash
	startHash := common.HexToHash(request.StartHash)
	endHash := common.HexToHash(request.EndHash)

	// get the transaction with the startHash
	filter := bson.D{primitive.E{Key: "hash", Value: startHash}}
	opts := options.FindOne().SetSort(bson.D{primitive.E{Key: "hash", Value: 1}})
	trx, err := service.transactionRepo.FindTransaction(filter, opts)
	if err != nil {
		return err
	}
	// check if endHash is a real transaction on the endChain
	// get the network with the name of trx.endchain
	network, err := service.networkService.GetNetworkByName(trx.EndChain)
	if err != nil {
		return err
	}
	receipt, err := checkReceipt(context.Background(), network.EVMHTTPURL, endHash)
	if err != nil {
		return err
	}
	// if it is, then finalize the transaction
	if receipt.Status == types.ReceiptStatusSuccessful {
		// update the transaction with the endHash and set the confirmed field to true in the database
		trx.EndHash = endHash
		trx.Confirmed = true
		service.UpdateTransaction(trx)
		// notify the other validators that the transaction has been finalized
		// send the transaction to all validators
		validatorList, err := service.validatorService.GetValidators()
		if err != nil {
			return err
		}
		for _, v := range validatorList {
			ip := v.Domain
			if ip != config.Conf.IP { // skip self and send to others
				// call the /transaction/finalize endpoint on the validator with
				// the startHash and endHash
				requestUrl := "http://" + ip + "/transaction/finalize"
				requestBody := map[string]interface{}{
					"startHash": startHash,
					"endHash":   endHash,
				}
				requestBytes, err := json.Marshal(requestBody)
				if err != nil {
					return err
				}
				resp, err := tools.API("POST", requestUrl, requestBytes)
				if err != nil {
					return err
				}
				log.Debug("Status ", resp.StatusCode, " from /transaction/finalize request to ", requestUrl)
				tools.Check(err)
			}
		}
	}

	return nil
}

func checkReceipt(ctx context.Context, evmHTTPURL string, txHash common.Hash) (*types.Receipt, error) {
	// Create a new ethclient.Client instance
	client, err := ethclient.Dial(evmHTTPURL)
	if err != nil {
		return nil, fmt.Errorf("failed to dial EVM HTTP URL: %v", err)
	}
	defer client.Close()

	receipt, err := client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve transaction receipt: %v", err)
	}

	return receipt, nil
}

func (service *TransactionServiceImpl) AddTransaction(transaction Transaction) error {

	filter := bson.D{primitive.E{Key: "hash", Value: transaction.Hash}}
	opts := options.FindOne().SetSort(bson.D{primitive.E{Key: "hash", Value: 1}})
	_, err := service.transactionRepo.FindTransaction(filter, opts)
	if err == nil { // transaction already exists in the database
		return err
	}

	transaction.ID = primitive.NewObjectID()
	_, err = service.transactionRepo.InsertTransaction(transaction)
	if err != nil {
		return err
	}
	log.Debug("Added new transaction with hash: ", transaction.Hash)

	// validate fields and insert transaction
	// send transaction to all validators
	validators, err := service.validatorService.GetValidators() // TODO: this should be move to service
	if err != nil {
		return err
	}
	for _, v := range validators {
		ip := v.Domain
		if ip != config.Conf.IP { // skip self and send to others
			err := SendTransaction(transaction, v)
			tools.Check(err)
		}
	}
	return nil
}

func (service *TransactionServiceImpl) UpdateTransaction(transaction Transaction) error {
	var trx Transaction
	var err error

	fmt.Println("UpdateTransaction: ", transaction.Hash)
	filter := bson.D{primitive.E{Key: "hash", Value: transaction.Hash}}
	opts := options.FindOne().SetSort(bson.D{primitive.E{Key: "hash", Value: 1}})
	trx, err = service.transactionRepo.FindTransaction(filter, opts)
	if err != nil { // transaction doesn't exist in the database
		log.Error("Transaction doesn't exist in the database")
		// return err
	}

	// find transaction
	var result bson.M

	err = service.transactionRepo.FindTransactionWithTransform(result, bson.D{{Key: "hash", Value: transaction.Hash}})
	if err != nil {
		log.Error("Transaction doesn't exist in the database 2", err)
		return err
	}
	fmt.Println("tx: ", result)

	trx.Signers = append(trx.Signers, transaction.Signers...)
	trx.R = append(trx.R, transaction.R...)
	trx.S = append(trx.S, transaction.S...)
	trx.V = append(trx.V, transaction.V...)
	trx.H = append(trx.H, transaction.H...)

	idFilter := bson.M{"_id": trx.ID}
	_, err = service.transactionRepo.UpdateTransaction(idFilter, trx)
	if tools.Check(err) {
		log.Debug("Transaction updated successfully with hash: ", trx.Hash)
	}
	log.Debug("Transaction updated successfully with hash: ", trx.EndChain)
	portNetwork, err := service.networkService.GetNetworkByName(trx.EndChain)
	if err != nil {
		log.Error("Error getting network by name: ", err)
		return err
	}

	log.Debug("Transaction updated successfully with hash: ", portNetwork.ContractAddress)

	requestUrl := "http://localhost:7045/signer/send-transaction"
	// call the key-server to get the signer time

	// the expected request body is a JSON object with the following fields:
	// type sendTransactionHandlerRequest struct {
	// 	Transaction     Transaction `json:"transaction"`
	// 	EVMHTTPURL      string      `json:"EVMHTTPURL"`
	// 	ContractAddress string      `json:"contractAddress"`
	// }
	requestBody := map[string]interface{}{
		"transaction":     trx,
		"EVMHTTPURL":      portNetwork.EVMHTTPURL,
		"ContractAddress": portNetwork.ContractAddress,
	}
	requestBytes, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	// Send the request to the key-server
	resp, err := tools.API("POST", requestUrl, requestBytes)
	if err != nil {
		return err
	}
	log.Debug("Status ", resp.StatusCode, " from /signer/send-transaction request to ", requestUrl)

	return err
}

// TODO: this should be removed
func (service *TransactionServiceImpl) sendTransaction(trx Transaction) error {
	// get netowrk with the name of trx.endchain
	network, err := service.networkService.GetNetworkByName(trx.EndChain)
	if err != nil {
		log.Error("Network not found with name: ", trx.EndChain)
		return err
	}
	fmt.Println("SEND TRANSACTION Network: ", network)
	service.auth.SetEVMClient(network.EVMHTTPURL)
	// authObj, err := service.auth.GetAuth(service.wallet.GetSigner(), network.ContractAddress, service.wallet.GetPublicKey())
	// tools.Check(err, "auth.GetAuth")
	// auth := authObj.Auth

	// fmt.Println("sendTransaction: ", trx.Destination)

	// var data portabi.CrossChainData
	// data.Addresses = trx.Addresses
	// data.Integers = trx.Uint256
	// data.Strings = trx.String
	// data.Bools = trx.Bool
	// tx, err := authObj.Instance.ExecuteInboundMessage(auth, "ETH", trx.Sender, trx.Destination, data, trx.V, trx.R, trx.S, trx.H)
	// tools.Check(err, "instance.ExecuteInboundMessage")
	// log.Debug("tx sent: ", tx.Hash().Hex())
	return nil
}

func NewTransactionService(
	transactionRepo TransactionRepository,
	auth auth.IAuth,
	// wallet wallet.Wallet,
	networkService networkmodule.NetworkService, // Change this line to have the correct parameter name
	validatorService validatormodule.ValidatorService,
) TransactionService {
	return &TransactionServiceImpl{
		transactionRepo: transactionRepo,
		auth:            auth,
		// wallet:           wallet,
		networkService:   networkService, // Now this should work correctly
		validatorService: validatorService,
	}
}
