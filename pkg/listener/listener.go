package listener

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"telegraph/log"
	networkmodule "telegraph/modules/network"
	transactionmodule "telegraph/modules/transaction"
	"time"

	portabi "telegraph/pkg/crypto_utils/port/evm/abi"
	"telegraph/pkg/db/validator"
	"telegraph/pkg/signer"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Listener interface {
	StartListener()
	CreateConnection(network networkmodule.Network)
	restartConnection(network networkmodule.Network)
	testEmitEvent(contractAbi abi.ABI, vLog types.Log)
	swapOutEvent(contractAbi abi.ABI, vLog types.Log)
}

type ListenerImpl struct {
	transactionService transactionmodule.TransactionService
	signer             signer.Signer
	networkService     networkmodule.NetworkService
}

type CrossChainData struct {
	Addresses []common.Address
	Numbers   []*big.Int
	Strings   []string
	Bools     []bool
}

var SwapOut struct {
	Sender         common.Address
	StartChain     string
	EndChain       string
	TransferAmount *big.Int
	Trigger        common.Address
	Data           CrossChainData
	lock           sync.Mutex
}

var TestEvent struct {
	Message string
	Sender  common.Address
	lock    sync.Mutex
}

var NewSignerEvent struct {
	Signer  common.Address
	Domain  string
	Moniker string
	lock    sync.Mutex
}

// Event hashes.
var logTestEmit = []byte("TestEmit(string,address)")
var logTestEmitSigHash = crypto.Keccak256Hash(logTestEmit)
var bridgeSwapOutSig = []byte("BridgeSwapOutData(address,string,string,uint256,address,(address[],uint256[],string[],bool[]))")
var bridgeSwapOutSigHash = crypto.Keccak256Hash(bridgeSwapOutSig)
var newSignerSig = []byte("NewSigner(address,string,string)")
var newSignerSigHash = crypto.Keccak256Hash(newSignerSig)

func (listener *ListenerImpl) StartListener() {
	savedNetworks, err := listener.networkService.GetNetworks()
	if err != nil {
		log.Fatal("networks.GetNetworks error:", err)
	}
	if len(savedNetworks) == 0 {
		log.Error("No networks to listen to")
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(savedNetworks))

	for i, v := range savedNetworks {
		go func(v networkmodule.Network, i int) {
			defer wg.Done()
			log.Debug("Connecting to Chain: ", i, v.Name)
			listener.CreateConnection(v)

		}(v, i)
	}

}

func (listener *ListenerImpl) CreateConnection(network networkmodule.Network) {
	log.Debug("Creating connection to: ", network.Name)
	client, err := ethclient.Dial(network.EVMWSSURL)
	fmt.Println("client: ", client)
	if err != nil {
		fmt.Println("ethclient.Dial error: ", network.Name, err)
		listener.restartConnection(network)
		return
		// log.Error("ethclient.Dial error:", network.Name, err)
	}
	log.Debug("Connected to: ", network.Name)

	contractAddress := common.HexToAddress(network.ContractAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal("client.SubscribeFilterLogs error:", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(portabi.PortabiABI))
	if err != nil {
		log.Fatal("abi.JSON error:", err)
	}

	for {
		select {
		case err := <-sub.Err():
			fmt.Println("sub.Err() error:", err)
			listener.restartConnection(network)
			return
		case vLog := <-logs:
			switch vLog.Topics[0].Hex() {
			case logTestEmitSigHash.Hex():
				listener.testEmitEvent(contractAbi, vLog)
			case bridgeSwapOutSigHash.Hex():
				listener.swapOutEvent(contractAbi, vLog)
			case newSignerSigHash.Hex():
				listener.newSignerEvent(contractAbi, vLog)
			}

		}
	}
}

func (listener *ListenerImpl) restartConnection(network networkmodule.Network) {
	// restart connection after 5 seconds
	time.Sleep(5 * time.Second)
	listener.CreateConnection(network)
}

func (listener *ListenerImpl) testEmitEvent(contractAbi abi.ABI, vLog types.Log) {
	TestEvent.lock.Lock()
	defer TestEvent.lock.Unlock()

	log.Debug("Received TestEmit event")
	err := contractAbi.UnpackIntoInterface(&TestEvent, "TestEmit", vLog.Data)
	if err != nil {
		log.Error("contractAbi.UnpackIntoInterface error:", err)
	}
	log.Debug("Test Emit Event Detected")
}

func (listener *ListenerImpl) swapOutEvent(contractAbi abi.ABI, vLog types.Log) {
	SwapOut.lock.Lock()
	defer SwapOut.lock.Unlock()

	log.Debug("Received BridgeSwapOut event")
	err := contractAbi.UnpackIntoInterface(&SwapOut, "BridgeSwapOutData", vLog.Data)
	if err != nil {
		log.Error("contractAbi.UnpackIntoInterface error:", err)
	}

	// sign transaction
	newTransaction := transactionmodule.Transaction{
		DetectionTime: time.Now(),
		Hash:          vLog.TxHash,
		EndHash:       common.Hash{},
		BlockNumber:   vLog.BlockNumber,
		LogIndex:      vLog.Index,
		Event:         "BridgeSwapOut",
		Sender:        SwapOut.Sender,
		Addresses:     SwapOut.Data.Addresses,
		Uint256:       SwapOut.Data.Numbers,
		String:        SwapOut.Data.Strings,
		Bool:          SwapOut.Data.Bools,
		FeeAmount:     SwapOut.TransferAmount,
		StartChain:    SwapOut.StartChain,
		EndChain:      SwapOut.EndChain,
		SignedCount:   0,
		Signers:       [][]byte{},
		R:             [][32]byte{},
		S:             [][32]byte{},
		V:             []uint8{},
		H:             [][32]byte{},
		Confirmed:     false,
	}
	_ = listener.transactionService.AddTransaction(newTransaction)
	listener.signer.SignTransaction(newTransaction)
}

func (listener *ListenerImpl) newSignerEvent(contractAbi abi.ABI, vLog types.Log) {
	NewSignerEvent.lock.Lock()
	defer NewSignerEvent.lock.Unlock()

	log.Debug("Received New Signer event")
	err := contractAbi.UnpackIntoInterface(&NewSignerEvent, "NewSigner", vLog.Data)
	if err != nil {
		log.Error("contractAbi.UnpackIntoInterface error:", err)
	}

	// Add new signer to the signer list
	validator.AddValidator(NewSignerEvent.Signer, NewSignerEvent.Moniker, NewSignerEvent.Domain, false)
}

func NewListener(transactionService transactionmodule.TransactionService, signer signer.Signer, networkService networkmodule.NetworkService) Listener {
	return &ListenerImpl{
		transactionService: transactionService,
		signer:             signer,
		networkService:     networkService,
	}

}
