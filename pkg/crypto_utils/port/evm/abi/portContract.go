// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package portabi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CrossChainData is an auto generated low-level Go binding around an user-defined struct.
type CrossChainData struct {
	Addresses []common.Address
	Integers  []*big.Int
	Strings   []string
	Bools     []bool
}

// PortabiMetaData contains all meta data concerning the Portabi contract.
var PortabiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"startChain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"integers\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"}],\"indexed\":false,\"internalType\":\"structCrossChainData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"BridgeSwapInData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"startChain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"endChain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transferAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trigger\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"integers\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"}],\"indexed\":false,\"internalType\":\"structCrossChainData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"BridgeSwapOutData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"NewSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"TestEmit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_moniker\",\"type\":\"string\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"name\":\"addSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"ownerCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"functionName\",\"type\":\"string\"}],\"name\":\"authorize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"endChain\",\"type\":\"string\"}],\"name\":\"determineFeeInCoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributionContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"entryFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_startChain\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"integers\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"}],\"internalType\":\"structCrossChainData\",\"name\":\"data\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"name\":\"executeInboundMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"getUnclaimedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"halvingHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"interval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"halvingAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inboundHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chain\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inboundIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"initializeAuthLib\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isValidSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastHalvingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastRewardsPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"outboundHistory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"preferredNode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"OPCode\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outboundIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"integers\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"strings\",\"type\":\"string[]\"},{\"internalType\":\"bool[]\",\"name\":\"bools\",\"type\":\"bool[]\"}],\"internalType\":\"structCrossChainData\",\"name\":\"data\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"endChain\",\"type\":\"string\"}],\"name\":\"outboundMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"priceMapping\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_startChain\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"entryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"entryFee\",\"type\":\"uint256\"}],\"name\":\"proxyConstructor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ownerToRemove\",\"type\":\"address\"}],\"name\":\"removeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"setChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"setDistributionContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_address\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fee\",\"type\":\"uint256[]\"}],\"name\":\"setEntryFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"startChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endChain\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setPriceMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signHistory\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signers\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"signerTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signersArr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startChain\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalClaimIntervals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transactionReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCode\",\"type\":\"address\"}],\"name\":\"updateCode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usedHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"voteNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PortabiABI is the input ABI used to generate the binding from.
// Deprecated: Use PortabiMetaData.ABI instead.
var PortabiABI = PortabiMetaData.ABI

// Portabi is an auto generated Go binding around an Ethereum contract.
type Portabi struct {
	PortabiCaller     // Read-only binding to the contract
	PortabiTransactor // Write-only binding to the contract
	PortabiFilterer   // Log filterer for contract events
}

// PortabiCaller is an auto generated read-only Go binding around an Ethereum contract.
type PortabiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PortabiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PortabiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PortabiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PortabiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PortabiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PortabiSession struct {
	Contract     *Portabi          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PortabiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PortabiCallerSession struct {
	Contract *PortabiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PortabiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PortabiTransactorSession struct {
	Contract     *PortabiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PortabiRaw is an auto generated low-level Go binding around an Ethereum contract.
type PortabiRaw struct {
	Contract *Portabi // Generic contract binding to access the raw methods on
}

// PortabiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PortabiCallerRaw struct {
	Contract *PortabiCaller // Generic read-only contract binding to access the raw methods on
}

// PortabiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PortabiTransactorRaw struct {
	Contract *PortabiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPortabi creates a new instance of Portabi, bound to a specific deployed contract.
func NewPortabi(address common.Address, backend bind.ContractBackend) (*Portabi, error) {
	contract, err := bindPortabi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Portabi{PortabiCaller: PortabiCaller{contract: contract}, PortabiTransactor: PortabiTransactor{contract: contract}, PortabiFilterer: PortabiFilterer{contract: contract}}, nil
}

// NewPortabiCaller creates a new read-only instance of Portabi, bound to a specific deployed contract.
func NewPortabiCaller(address common.Address, caller bind.ContractCaller) (*PortabiCaller, error) {
	contract, err := bindPortabi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PortabiCaller{contract: contract}, nil
}

// NewPortabiTransactor creates a new write-only instance of Portabi, bound to a specific deployed contract.
func NewPortabiTransactor(address common.Address, transactor bind.ContractTransactor) (*PortabiTransactor, error) {
	contract, err := bindPortabi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PortabiTransactor{contract: contract}, nil
}

// NewPortabiFilterer creates a new log filterer instance of Portabi, bound to a specific deployed contract.
func NewPortabiFilterer(address common.Address, filterer bind.ContractFilterer) (*PortabiFilterer, error) {
	contract, err := bindPortabi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PortabiFilterer{contract: contract}, nil
}

// bindPortabi binds a generic wrapper to an already deployed contract.
func bindPortabi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PortabiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Portabi *PortabiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Portabi.Contract.PortabiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Portabi *PortabiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Portabi.Contract.PortabiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Portabi *PortabiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Portabi.Contract.PortabiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Portabi *PortabiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Portabi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Portabi *PortabiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Portabi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Portabi *PortabiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Portabi.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Portabi *PortabiCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Portabi *PortabiSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Portabi.Contract.Allowance(&_Portabi.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Portabi *PortabiCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Portabi.Contract.Allowance(&_Portabi.CallOpts, owner, spender)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(bool initialized, uint256 ownerCount)
func (_Portabi *PortabiCaller) Auth(opts *bind.CallOpts) (struct {
	Initialized bool
	OwnerCount  *big.Int
}, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "auth")

	outstruct := new(struct {
		Initialized bool
		OwnerCount  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Initialized = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.OwnerCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(bool initialized, uint256 ownerCount)
func (_Portabi *PortabiSession) Auth() (struct {
	Initialized bool
	OwnerCount  *big.Int
}, error) {
	return _Portabi.Contract.Auth(&_Portabi.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(bool initialized, uint256 ownerCount)
func (_Portabi *PortabiCallerSession) Auth() (struct {
	Initialized bool
	OwnerCount  *big.Int
}, error) {
	return _Portabi.Contract.Auth(&_Portabi.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Portabi *PortabiCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Portabi *PortabiSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Portabi.Contract.BalanceOf(&_Portabi.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Portabi *PortabiCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Portabi.Contract.BalanceOf(&_Portabi.CallOpts, account)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Portabi *PortabiCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Portabi *PortabiSession) ChainId() (*big.Int, error) {
	return _Portabi.Contract.ChainId(&_Portabi.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Portabi *PortabiCallerSession) ChainId() (*big.Int, error) {
	return _Portabi.Contract.ChainId(&_Portabi.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Portabi *PortabiCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Portabi *PortabiSession) Decimals() (uint8, error) {
	return _Portabi.Contract.Decimals(&_Portabi.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Portabi *PortabiCallerSession) Decimals() (uint8, error) {
	return _Portabi.Contract.Decimals(&_Portabi.CallOpts)
}

// DetermineFeeInCoin is a free data retrieval call binding the contract method 0x026cd651.
//
// Solidity: function determineFeeInCoin(string endChain) view returns(uint256)
func (_Portabi *PortabiCaller) DetermineFeeInCoin(opts *bind.CallOpts, endChain string) (*big.Int, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "determineFeeInCoin", endChain)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DetermineFeeInCoin is a free data retrieval call binding the contract method 0x026cd651.
//
// Solidity: function determineFeeInCoin(string endChain) view returns(uint256)
func (_Portabi *PortabiSession) DetermineFeeInCoin(endChain string) (*big.Int, error) {
	return _Portabi.Contract.DetermineFeeInCoin(&_Portabi.CallOpts, endChain)
}

// DetermineFeeInCoin is a free data retrieval call binding the contract method 0x026cd651.
//
// Solidity: function determineFeeInCoin(string endChain) view returns(uint256)
func (_Portabi *PortabiCallerSession) DetermineFeeInCoin(endChain string) (*big.Int, error) {
	return _Portabi.Contract.DetermineFeeInCoin(&_Portabi.CallOpts, endChain)
}

// DistributionContract is a free data retrieval call binding the contract method 0x5a4528c2.
//
// Solidity: function distributionContract() view returns(address)
func (_Portabi *PortabiCaller) DistributionContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "distributionContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DistributionContract is a free data retrieval call binding the contract method 0x5a4528c2.
//
// Solidity: function distributionContract() view returns(address)
func (_Portabi *PortabiSession) DistributionContract() (common.Address, error) {
	return _Portabi.Contract.DistributionContract(&_Portabi.CallOpts)
}

// DistributionContract is a free data retrieval call binding the contract method 0x5a4528c2.
//
// Solidity: function distributionContract() view returns(address)
func (_Portabi *PortabiCallerSession) DistributionContract() (common.Address, error) {
	return _Portabi.Contract.DistributionContract(&_Portabi.CallOpts)
}

// EntryFees is a free data retrieval call binding the contract method 0xdb5c448e.
//
// Solidity: function entryFees(address ) view returns(uint256)
func (_Portabi *PortabiCaller) EntryFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "entryFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EntryFees is a free data retrieval call binding the contract method 0xdb5c448e.
//
// Solidity: function entryFees(address ) view returns(uint256)
func (_Portabi *PortabiSession) EntryFees(arg0 common.Address) (*big.Int, error) {
	return _Portabi.Contract.EntryFees(&_Portabi.CallOpts, arg0)
}

// EntryFees is a free data retrieval call binding the contract method 0xdb5c448e.
//
// Solidity: function entryFees(address ) view returns(uint256)
func (_Portabi *PortabiCallerSession) EntryFees(arg0 common.Address) (*big.Int, error) {
	return _Portabi.Contract.EntryFees(&_Portabi.CallOpts, arg0)
}

// GetUnclaimedRewards is a free data retrieval call binding the contract method 0x69a69e29.
//
// Solidity: function getUnclaimedRewards(address _signer) view returns(uint256)
func (_Portabi *PortabiCaller) GetUnclaimedRewards(opts *bind.CallOpts, _signer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "getUnclaimedRewards", _signer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnclaimedRewards is a free data retrieval call binding the contract method 0x69a69e29.
//
// Solidity: function getUnclaimedRewards(address _signer) view returns(uint256)
func (_Portabi *PortabiSession) GetUnclaimedRewards(_signer common.Address) (*big.Int, error) {
	return _Portabi.Contract.GetUnclaimedRewards(&_Portabi.CallOpts, _signer)
}

// GetUnclaimedRewards is a free data retrieval call binding the contract method 0x69a69e29.
//
// Solidity: function getUnclaimedRewards(address _signer) view returns(uint256)
func (_Portabi *PortabiCallerSession) GetUnclaimedRewards(_signer common.Address) (*big.Int, error) {
	return _Portabi.Contract.GetUnclaimedRewards(&_Portabi.CallOpts, _signer)
}

// HalvingHistory is a free data retrieval call binding the contract method 0x7d5d5aba.
//
// Solidity: function halvingHistory(uint256 ) view returns(uint256 interval, uint256 halvingAmount)
func (_Portabi *PortabiCaller) HalvingHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Interval      *big.Int
	HalvingAmount *big.Int
}, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "halvingHistory", arg0)

	outstruct := new(struct {
		Interval      *big.Int
		HalvingAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Interval = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.HalvingAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// HalvingHistory is a free data retrieval call binding the contract method 0x7d5d5aba.
//
// Solidity: function halvingHistory(uint256 ) view returns(uint256 interval, uint256 halvingAmount)
func (_Portabi *PortabiSession) HalvingHistory(arg0 *big.Int) (struct {
	Interval      *big.Int
	HalvingAmount *big.Int
}, error) {
	return _Portabi.Contract.HalvingHistory(&_Portabi.CallOpts, arg0)
}

// HalvingHistory is a free data retrieval call binding the contract method 0x7d5d5aba.
//
// Solidity: function halvingHistory(uint256 ) view returns(uint256 interval, uint256 halvingAmount)
func (_Portabi *PortabiCallerSession) HalvingHistory(arg0 *big.Int) (struct {
	Interval      *big.Int
	HalvingAmount *big.Int
}, error) {
	return _Portabi.Contract.HalvingHistory(&_Portabi.CallOpts, arg0)
}

// InboundHistory is a free data retrieval call binding the contract method 0xb0684bc2.
//
// Solidity: function inboundHistory(uint256 ) view returns(uint256 amount, address sender, address destination, string chain)
func (_Portabi *PortabiCaller) InboundHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount      *big.Int
	Sender      common.Address
	Destination common.Address
	Chain       string
}, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "inboundHistory", arg0)

	outstruct := new(struct {
		Amount      *big.Int
		Sender      common.Address
		Destination common.Address
		Chain       string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Sender = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Destination = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Chain = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// InboundHistory is a free data retrieval call binding the contract method 0xb0684bc2.
//
// Solidity: function inboundHistory(uint256 ) view returns(uint256 amount, address sender, address destination, string chain)
func (_Portabi *PortabiSession) InboundHistory(arg0 *big.Int) (struct {
	Amount      *big.Int
	Sender      common.Address
	Destination common.Address
	Chain       string
}, error) {
	return _Portabi.Contract.InboundHistory(&_Portabi.CallOpts, arg0)
}

// InboundHistory is a free data retrieval call binding the contract method 0xb0684bc2.
//
// Solidity: function inboundHistory(uint256 ) view returns(uint256 amount, address sender, address destination, string chain)
func (_Portabi *PortabiCallerSession) InboundHistory(arg0 *big.Int) (struct {
	Amount      *big.Int
	Sender      common.Address
	Destination common.Address
	Chain       string
}, error) {
	return _Portabi.Contract.InboundHistory(&_Portabi.CallOpts, arg0)
}

// InboundIndex is a free data retrieval call binding the contract method 0x4fba35ef.
//
// Solidity: function inboundIndex() view returns(uint256)
func (_Portabi *PortabiCaller) InboundIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "inboundIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InboundIndex is a free data retrieval call binding the contract method 0x4fba35ef.
//
// Solidity: function inboundIndex() view returns(uint256)
func (_Portabi *PortabiSession) InboundIndex() (*big.Int, error) {
	return _Portabi.Contract.InboundIndex(&_Portabi.CallOpts)
}

// InboundIndex is a free data retrieval call binding the contract method 0x4fba35ef.
//
// Solidity: function inboundIndex() view returns(uint256)
func (_Portabi *PortabiCallerSession) InboundIndex() (*big.Int, error) {
	return _Portabi.Contract.InboundIndex(&_Portabi.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Portabi *PortabiCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Portabi *PortabiSession) Initialized() (bool, error) {
	return _Portabi.Contract.Initialized(&_Portabi.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Portabi *PortabiCallerSession) Initialized() (bool, error) {
	return _Portabi.Contract.Initialized(&_Portabi.CallOpts)
}

// IsValidSigner is a free data retrieval call binding the contract method 0xd5f50582.
//
// Solidity: function isValidSigner(address ) view returns(bool)
func (_Portabi *PortabiCaller) IsValidSigner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "isValidSigner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidSigner is a free data retrieval call binding the contract method 0xd5f50582.
//
// Solidity: function isValidSigner(address ) view returns(bool)
func (_Portabi *PortabiSession) IsValidSigner(arg0 common.Address) (bool, error) {
	return _Portabi.Contract.IsValidSigner(&_Portabi.CallOpts, arg0)
}

// IsValidSigner is a free data retrieval call binding the contract method 0xd5f50582.
//
// Solidity: function isValidSigner(address ) view returns(bool)
func (_Portabi *PortabiCallerSession) IsValidSigner(arg0 common.Address) (bool, error) {
	return _Portabi.Contract.IsValidSigner(&_Portabi.CallOpts, arg0)
}

// LastHalvingTime is a free data retrieval call binding the contract method 0x8cbb6df7.
//
// Solidity: function lastHalvingTime() view returns(uint256)
func (_Portabi *PortabiCaller) LastHalvingTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "lastHalvingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastHalvingTime is a free data retrieval call binding the contract method 0x8cbb6df7.
//
// Solidity: function lastHalvingTime() view returns(uint256)
func (_Portabi *PortabiSession) LastHalvingTime() (*big.Int, error) {
	return _Portabi.Contract.LastHalvingTime(&_Portabi.CallOpts)
}

// LastHalvingTime is a free data retrieval call binding the contract method 0x8cbb6df7.
//
// Solidity: function lastHalvingTime() view returns(uint256)
func (_Portabi *PortabiCallerSession) LastHalvingTime() (*big.Int, error) {
	return _Portabi.Contract.LastHalvingTime(&_Portabi.CallOpts)
}

// LastRewardsPerShare is a free data retrieval call binding the contract method 0x6f19e0c8.
//
// Solidity: function lastRewardsPerShare(address ) view returns(uint256)
func (_Portabi *PortabiCaller) LastRewardsPerShare(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "lastRewardsPerShare", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardsPerShare is a free data retrieval call binding the contract method 0x6f19e0c8.
//
// Solidity: function lastRewardsPerShare(address ) view returns(uint256)
func (_Portabi *PortabiSession) LastRewardsPerShare(arg0 common.Address) (*big.Int, error) {
	return _Portabi.Contract.LastRewardsPerShare(&_Portabi.CallOpts, arg0)
}

// LastRewardsPerShare is a free data retrieval call binding the contract method 0x6f19e0c8.
//
// Solidity: function lastRewardsPerShare(address ) view returns(uint256)
func (_Portabi *PortabiCallerSession) LastRewardsPerShare(arg0 common.Address) (*big.Int, error) {
	return _Portabi.Contract.LastRewardsPerShare(&_Portabi.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Portabi *PortabiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Portabi *PortabiSession) Name() (string, error) {
	return _Portabi.Contract.Name(&_Portabi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Portabi *PortabiCallerSession) Name() (string, error) {
	return _Portabi.Contract.Name(&_Portabi.CallOpts)
}

// NodeStartTime is a free data retrieval call binding the contract method 0xd50a1857.
//
// Solidity: function nodeStartTime() view returns(uint256)
func (_Portabi *PortabiCaller) NodeStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "nodeStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeStartTime is a free data retrieval call binding the contract method 0xd50a1857.
//
// Solidity: function nodeStartTime() view returns(uint256)
func (_Portabi *PortabiSession) NodeStartTime() (*big.Int, error) {
	return _Portabi.Contract.NodeStartTime(&_Portabi.CallOpts)
}

// NodeStartTime is a free data retrieval call binding the contract method 0xd50a1857.
//
// Solidity: function nodeStartTime() view returns(uint256)
func (_Portabi *PortabiCallerSession) NodeStartTime() (*big.Int, error) {
	return _Portabi.Contract.NodeStartTime(&_Portabi.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Portabi *PortabiCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Portabi *PortabiSession) Nonce() (*big.Int, error) {
	return _Portabi.Contract.Nonce(&_Portabi.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Portabi *PortabiCallerSession) Nonce() (*big.Int, error) {
	return _Portabi.Contract.Nonce(&_Portabi.CallOpts)
}

// OutboundHistory is a free data retrieval call binding the contract method 0xc75fe374.
//
// Solidity: function outboundHistory(uint256 ) view returns(address sender, uint256 feeAmount, address destination, string chain, string preferredNode, string OPCode)
func (_Portabi *PortabiCaller) OutboundHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Sender        common.Address
	FeeAmount     *big.Int
	Destination   common.Address
	Chain         string
	PreferredNode string
	OPCode        string
}, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "outboundHistory", arg0)

	outstruct := new(struct {
		Sender        common.Address
		FeeAmount     *big.Int
		Destination   common.Address
		Chain         string
		PreferredNode string
		OPCode        string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sender = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.FeeAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Destination = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Chain = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.PreferredNode = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.OPCode = *abi.ConvertType(out[5], new(string)).(*string)

	return *outstruct, err

}

// OutboundHistory is a free data retrieval call binding the contract method 0xc75fe374.
//
// Solidity: function outboundHistory(uint256 ) view returns(address sender, uint256 feeAmount, address destination, string chain, string preferredNode, string OPCode)
func (_Portabi *PortabiSession) OutboundHistory(arg0 *big.Int) (struct {
	Sender        common.Address
	FeeAmount     *big.Int
	Destination   common.Address
	Chain         string
	PreferredNode string
	OPCode        string
}, error) {
	return _Portabi.Contract.OutboundHistory(&_Portabi.CallOpts, arg0)
}

// OutboundHistory is a free data retrieval call binding the contract method 0xc75fe374.
//
// Solidity: function outboundHistory(uint256 ) view returns(address sender, uint256 feeAmount, address destination, string chain, string preferredNode, string OPCode)
func (_Portabi *PortabiCallerSession) OutboundHistory(arg0 *big.Int) (struct {
	Sender        common.Address
	FeeAmount     *big.Int
	Destination   common.Address
	Chain         string
	PreferredNode string
	OPCode        string
}, error) {
	return _Portabi.Contract.OutboundHistory(&_Portabi.CallOpts, arg0)
}

// OutboundIndex is a free data retrieval call binding the contract method 0x488ec3e8.
//
// Solidity: function outboundIndex() view returns(uint256)
func (_Portabi *PortabiCaller) OutboundIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "outboundIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutboundIndex is a free data retrieval call binding the contract method 0x488ec3e8.
//
// Solidity: function outboundIndex() view returns(uint256)
func (_Portabi *PortabiSession) OutboundIndex() (*big.Int, error) {
	return _Portabi.Contract.OutboundIndex(&_Portabi.CallOpts)
}

// OutboundIndex is a free data retrieval call binding the contract method 0x488ec3e8.
//
// Solidity: function outboundIndex() view returns(uint256)
func (_Portabi *PortabiCallerSession) OutboundIndex() (*big.Int, error) {
	return _Portabi.Contract.OutboundIndex(&_Portabi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Portabi *PortabiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Portabi *PortabiSession) Owner() (common.Address, error) {
	return _Portabi.Contract.Owner(&_Portabi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Portabi *PortabiCallerSession) Owner() (common.Address, error) {
	return _Portabi.Contract.Owner(&_Portabi.CallOpts)
}

// PriceMapping is a free data retrieval call binding the contract method 0x2616dec5.
//
// Solidity: function priceMapping(string , string ) view returns(uint256)
func (_Portabi *PortabiCaller) PriceMapping(opts *bind.CallOpts, arg0 string, arg1 string) (*big.Int, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "priceMapping", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceMapping is a free data retrieval call binding the contract method 0x2616dec5.
//
// Solidity: function priceMapping(string , string ) view returns(uint256)
func (_Portabi *PortabiSession) PriceMapping(arg0 string, arg1 string) (*big.Int, error) {
	return _Portabi.Contract.PriceMapping(&_Portabi.CallOpts, arg0, arg1)
}

// PriceMapping is a free data retrieval call binding the contract method 0x2616dec5.
//
// Solidity: function priceMapping(string , string ) view returns(uint256)
func (_Portabi *PortabiCallerSession) PriceMapping(arg0 string, arg1 string) (*big.Int, error) {
	return _Portabi.Contract.PriceMapping(&_Portabi.CallOpts, arg0, arg1)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() pure returns(bytes32)
func (_Portabi *PortabiCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() pure returns(bytes32)
func (_Portabi *PortabiSession) ProxiableUUID() ([32]byte, error) {
	return _Portabi.Contract.ProxiableUUID(&_Portabi.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() pure returns(bytes32)
func (_Portabi *PortabiCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Portabi.Contract.ProxiableUUID(&_Portabi.CallOpts)
}

// RewardsPerShare is a free data retrieval call binding the contract method 0xc7e1d0b1.
//
// Solidity: function rewardsPerShare() view returns(uint256)
func (_Portabi *PortabiCaller) RewardsPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "rewardsPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsPerShare is a free data retrieval call binding the contract method 0xc7e1d0b1.
//
// Solidity: function rewardsPerShare() view returns(uint256)
func (_Portabi *PortabiSession) RewardsPerShare() (*big.Int, error) {
	return _Portabi.Contract.RewardsPerShare(&_Portabi.CallOpts)
}

// RewardsPerShare is a free data retrieval call binding the contract method 0xc7e1d0b1.
//
// Solidity: function rewardsPerShare() view returns(uint256)
func (_Portabi *PortabiCallerSession) RewardsPerShare() (*big.Int, error) {
	return _Portabi.Contract.RewardsPerShare(&_Portabi.CallOpts)
}

// SignHistory is a free data retrieval call binding the contract method 0xbe5755d7.
//
// Solidity: function signHistory(uint256 , address ) view returns(bool)
func (_Portabi *PortabiCaller) SignHistory(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "signHistory", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SignHistory is a free data retrieval call binding the contract method 0xbe5755d7.
//
// Solidity: function signHistory(uint256 , address ) view returns(bool)
func (_Portabi *PortabiSession) SignHistory(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Portabi.Contract.SignHistory(&_Portabi.CallOpts, arg0, arg1)
}

// SignHistory is a free data retrieval call binding the contract method 0xbe5755d7.
//
// Solidity: function signHistory(uint256 , address ) view returns(bool)
func (_Portabi *PortabiCallerSession) SignHistory(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Portabi.Contract.SignHistory(&_Portabi.CallOpts, arg0, arg1)
}

// Signers is a free data retrieval call binding the contract method 0x736c0d5b.
//
// Solidity: function signers(address ) view returns(string domain, string moniker, uint256 signerTime)
func (_Portabi *PortabiCaller) Signers(opts *bind.CallOpts, arg0 common.Address) (struct {
	Domain     string
	Moniker    string
	SignerTime *big.Int
}, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "signers", arg0)

	outstruct := new(struct {
		Domain     string
		Moniker    string
		SignerTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Domain = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Moniker = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.SignerTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Signers is a free data retrieval call binding the contract method 0x736c0d5b.
//
// Solidity: function signers(address ) view returns(string domain, string moniker, uint256 signerTime)
func (_Portabi *PortabiSession) Signers(arg0 common.Address) (struct {
	Domain     string
	Moniker    string
	SignerTime *big.Int
}, error) {
	return _Portabi.Contract.Signers(&_Portabi.CallOpts, arg0)
}

// Signers is a free data retrieval call binding the contract method 0x736c0d5b.
//
// Solidity: function signers(address ) view returns(string domain, string moniker, uint256 signerTime)
func (_Portabi *PortabiCallerSession) Signers(arg0 common.Address) (struct {
	Domain     string
	Moniker    string
	SignerTime *big.Int
}, error) {
	return _Portabi.Contract.Signers(&_Portabi.CallOpts, arg0)
}

// SignersArr is a free data retrieval call binding the contract method 0xdc5fbe14.
//
// Solidity: function signersArr(uint256 ) view returns(address)
func (_Portabi *PortabiCaller) SignersArr(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "signersArr", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignersArr is a free data retrieval call binding the contract method 0xdc5fbe14.
//
// Solidity: function signersArr(uint256 ) view returns(address)
func (_Portabi *PortabiSession) SignersArr(arg0 *big.Int) (common.Address, error) {
	return _Portabi.Contract.SignersArr(&_Portabi.CallOpts, arg0)
}

// SignersArr is a free data retrieval call binding the contract method 0xdc5fbe14.
//
// Solidity: function signersArr(uint256 ) view returns(address)
func (_Portabi *PortabiCallerSession) SignersArr(arg0 *big.Int) (common.Address, error) {
	return _Portabi.Contract.SignersArr(&_Portabi.CallOpts, arg0)
}

// StartChain is a free data retrieval call binding the contract method 0x80185794.
//
// Solidity: function startChain() view returns(string)
func (_Portabi *PortabiCaller) StartChain(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "startChain")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StartChain is a free data retrieval call binding the contract method 0x80185794.
//
// Solidity: function startChain() view returns(string)
func (_Portabi *PortabiSession) StartChain() (string, error) {
	return _Portabi.Contract.StartChain(&_Portabi.CallOpts)
}

// StartChain is a free data retrieval call binding the contract method 0x80185794.
//
// Solidity: function startChain() view returns(string)
func (_Portabi *PortabiCallerSession) StartChain() (string, error) {
	return _Portabi.Contract.StartChain(&_Portabi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Portabi *PortabiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Portabi *PortabiSession) Symbol() (string, error) {
	return _Portabi.Contract.Symbol(&_Portabi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Portabi *PortabiCallerSession) Symbol() (string, error) {
	return _Portabi.Contract.Symbol(&_Portabi.CallOpts)
}

// TotalClaimIntervals is a free data retrieval call binding the contract method 0x6c84b714.
//
// Solidity: function totalClaimIntervals() view returns(uint256)
func (_Portabi *PortabiCaller) TotalClaimIntervals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "totalClaimIntervals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalClaimIntervals is a free data retrieval call binding the contract method 0x6c84b714.
//
// Solidity: function totalClaimIntervals() view returns(uint256)
func (_Portabi *PortabiSession) TotalClaimIntervals() (*big.Int, error) {
	return _Portabi.Contract.TotalClaimIntervals(&_Portabi.CallOpts)
}

// TotalClaimIntervals is a free data retrieval call binding the contract method 0x6c84b714.
//
// Solidity: function totalClaimIntervals() view returns(uint256)
func (_Portabi *PortabiCallerSession) TotalClaimIntervals() (*big.Int, error) {
	return _Portabi.Contract.TotalClaimIntervals(&_Portabi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Portabi *PortabiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Portabi *PortabiSession) TotalSupply() (*big.Int, error) {
	return _Portabi.Contract.TotalSupply(&_Portabi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Portabi *PortabiCallerSession) TotalSupply() (*big.Int, error) {
	return _Portabi.Contract.TotalSupply(&_Portabi.CallOpts)
}

// TransactionReward is a free data retrieval call binding the contract method 0x9437a448.
//
// Solidity: function transactionReward() view returns(uint256)
func (_Portabi *PortabiCaller) TransactionReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "transactionReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransactionReward is a free data retrieval call binding the contract method 0x9437a448.
//
// Solidity: function transactionReward() view returns(uint256)
func (_Portabi *PortabiSession) TransactionReward() (*big.Int, error) {
	return _Portabi.Contract.TransactionReward(&_Portabi.CallOpts)
}

// TransactionReward is a free data retrieval call binding the contract method 0x9437a448.
//
// Solidity: function transactionReward() view returns(uint256)
func (_Portabi *PortabiCallerSession) TransactionReward() (*big.Int, error) {
	return _Portabi.Contract.TransactionReward(&_Portabi.CallOpts)
}

// UsedHashes is a free data retrieval call binding the contract method 0xaef18bf7.
//
// Solidity: function usedHashes(bytes32 ) view returns(bool)
func (_Portabi *PortabiCaller) UsedHashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Portabi.contract.Call(opts, &out, "usedHashes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedHashes is a free data retrieval call binding the contract method 0xaef18bf7.
//
// Solidity: function usedHashes(bytes32 ) view returns(bool)
func (_Portabi *PortabiSession) UsedHashes(arg0 [32]byte) (bool, error) {
	return _Portabi.Contract.UsedHashes(&_Portabi.CallOpts, arg0)
}

// UsedHashes is a free data retrieval call binding the contract method 0xaef18bf7.
//
// Solidity: function usedHashes(bytes32 ) view returns(bool)
func (_Portabi *PortabiCallerSession) UsedHashes(arg0 [32]byte) (bool, error) {
	return _Portabi.Contract.UsedHashes(&_Portabi.CallOpts, arg0)
}

// AddSigner is a paid mutator transaction binding the contract method 0x0e7258b3.
//
// Solidity: function addSigner(address _signer, address _feeAddress, string _domain, string _moniker, uint8[] sigV, bytes32[] sigR, bytes32[] sigS, bytes32[] hashes) returns()
func (_Portabi *PortabiTransactor) AddSigner(opts *bind.TransactOpts, _signer common.Address, _feeAddress common.Address, _domain string, _moniker string, sigV []uint8, sigR [][32]byte, sigS [][32]byte, hashes [][32]byte) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "addSigner", _signer, _feeAddress, _domain, _moniker, sigV, sigR, sigS, hashes)
}

// AddSigner is a paid mutator transaction binding the contract method 0x0e7258b3.
//
// Solidity: function addSigner(address _signer, address _feeAddress, string _domain, string _moniker, uint8[] sigV, bytes32[] sigR, bytes32[] sigS, bytes32[] hashes) returns()
func (_Portabi *PortabiSession) AddSigner(_signer common.Address, _feeAddress common.Address, _domain string, _moniker string, sigV []uint8, sigR [][32]byte, sigS [][32]byte, hashes [][32]byte) (*types.Transaction, error) {
	return _Portabi.Contract.AddSigner(&_Portabi.TransactOpts, _signer, _feeAddress, _domain, _moniker, sigV, sigR, sigS, hashes)
}

// AddSigner is a paid mutator transaction binding the contract method 0x0e7258b3.
//
// Solidity: function addSigner(address _signer, address _feeAddress, string _domain, string _moniker, uint8[] sigV, bytes32[] sigR, bytes32[] sigS, bytes32[] hashes) returns()
func (_Portabi *PortabiTransactorSession) AddSigner(_signer common.Address, _feeAddress common.Address, _domain string, _moniker string, sigV []uint8, sigR [][32]byte, sigS [][32]byte, hashes [][32]byte) (*types.Transaction, error) {
	return _Portabi.Contract.AddSigner(&_Portabi.TransactOpts, _signer, _feeAddress, _domain, _moniker, sigV, sigR, sigS, hashes)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Portabi *PortabiTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Portabi *PortabiSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Portabi.Contract.Approve(&_Portabi.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Portabi *PortabiTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Portabi.Contract.Approve(&_Portabi.TransactOpts, spender, amount)
}

// Authorize is a paid mutator transaction binding the contract method 0x4f5fcb34.
//
// Solidity: function authorize(string functionName) returns()
func (_Portabi *PortabiTransactor) Authorize(opts *bind.TransactOpts, functionName string) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "authorize", functionName)
}

// Authorize is a paid mutator transaction binding the contract method 0x4f5fcb34.
//
// Solidity: function authorize(string functionName) returns()
func (_Portabi *PortabiSession) Authorize(functionName string) (*types.Transaction, error) {
	return _Portabi.Contract.Authorize(&_Portabi.TransactOpts, functionName)
}

// Authorize is a paid mutator transaction binding the contract method 0x4f5fcb34.
//
// Solidity: function authorize(string functionName) returns()
func (_Portabi *PortabiTransactorSession) Authorize(functionName string) (*types.Transaction, error) {
	return _Portabi.Contract.Authorize(&_Portabi.TransactOpts, functionName)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Portabi *PortabiTransactor) ClaimReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "claimReward")
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Portabi *PortabiSession) ClaimReward() (*types.Transaction, error) {
	return _Portabi.Contract.ClaimReward(&_Portabi.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Portabi *PortabiTransactorSession) ClaimReward() (*types.Transaction, error) {
	return _Portabi.Contract.ClaimReward(&_Portabi.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Portabi *PortabiTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Portabi *PortabiSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Portabi.Contract.DecreaseAllowance(&_Portabi.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Portabi *PortabiTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Portabi.Contract.DecreaseAllowance(&_Portabi.TransactOpts, spender, subtractedValue)
}

// ExecuteInboundMessage is a paid mutator transaction binding the contract method 0xe462c079.
//
// Solidity: function executeInboundMessage(string _startChain, address sender, address destination, (address[],uint256[],string[],bool[]) data, uint8[] sigV, bytes32[] sigR, bytes32[] sigS, bytes32[] hashes) returns()
func (_Portabi *PortabiTransactor) ExecuteInboundMessage(opts *bind.TransactOpts, _startChain string, sender common.Address, destination common.Address, data CrossChainData, sigV []uint8, sigR [][32]byte, sigS [][32]byte, hashes [][32]byte) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "executeInboundMessage", _startChain, sender, destination, data, sigV, sigR, sigS, hashes)
}

// ExecuteInboundMessage is a paid mutator transaction binding the contract method 0xe462c079.
//
// Solidity: function executeInboundMessage(string _startChain, address sender, address destination, (address[],uint256[],string[],bool[]) data, uint8[] sigV, bytes32[] sigR, bytes32[] sigS, bytes32[] hashes) returns()
func (_Portabi *PortabiSession) ExecuteInboundMessage(_startChain string, sender common.Address, destination common.Address, data CrossChainData, sigV []uint8, sigR [][32]byte, sigS [][32]byte, hashes [][32]byte) (*types.Transaction, error) {
	return _Portabi.Contract.ExecuteInboundMessage(&_Portabi.TransactOpts, _startChain, sender, destination, data, sigV, sigR, sigS, hashes)
}

// ExecuteInboundMessage is a paid mutator transaction binding the contract method 0xe462c079.
//
// Solidity: function executeInboundMessage(string _startChain, address sender, address destination, (address[],uint256[],string[],bool[]) data, uint8[] sigV, bytes32[] sigR, bytes32[] sigS, bytes32[] hashes) returns()
func (_Portabi *PortabiTransactorSession) ExecuteInboundMessage(_startChain string, sender common.Address, destination common.Address, data CrossChainData, sigV []uint8, sigR [][32]byte, sigS [][32]byte, hashes [][32]byte) (*types.Transaction, error) {
	return _Portabi.Contract.ExecuteInboundMessage(&_Portabi.TransactOpts, _startChain, sender, destination, data, sigV, sigR, sigS, hashes)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Portabi *PortabiTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Portabi *PortabiSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Portabi.Contract.IncreaseAllowance(&_Portabi.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Portabi *PortabiTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Portabi.Contract.IncreaseAllowance(&_Portabi.TransactOpts, spender, addedValue)
}

// InitializeAuthLib is a paid mutator transaction binding the contract method 0x14395e2e.
//
// Solidity: function initializeAuthLib(address[] owners) returns()
func (_Portabi *PortabiTransactor) InitializeAuthLib(opts *bind.TransactOpts, owners []common.Address) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "initializeAuthLib", owners)
}

// InitializeAuthLib is a paid mutator transaction binding the contract method 0x14395e2e.
//
// Solidity: function initializeAuthLib(address[] owners) returns()
func (_Portabi *PortabiSession) InitializeAuthLib(owners []common.Address) (*types.Transaction, error) {
	return _Portabi.Contract.InitializeAuthLib(&_Portabi.TransactOpts, owners)
}

// InitializeAuthLib is a paid mutator transaction binding the contract method 0x14395e2e.
//
// Solidity: function initializeAuthLib(address[] owners) returns()
func (_Portabi *PortabiTransactorSession) InitializeAuthLib(owners []common.Address) (*types.Transaction, error) {
	return _Portabi.Contract.InitializeAuthLib(&_Portabi.TransactOpts, owners)
}

// OutboundMessage is a paid mutator transaction binding the contract method 0x73a6eb14.
//
// Solidity: function outboundMessage(address sender, address destination, (address[],uint256[],string[],bool[]) data, string endChain) payable returns()
func (_Portabi *PortabiTransactor) OutboundMessage(opts *bind.TransactOpts, sender common.Address, destination common.Address, data CrossChainData, endChain string) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "outboundMessage", sender, destination, data, endChain)
}

// OutboundMessage is a paid mutator transaction binding the contract method 0x73a6eb14.
//
// Solidity: function outboundMessage(address sender, address destination, (address[],uint256[],string[],bool[]) data, string endChain) payable returns()
func (_Portabi *PortabiSession) OutboundMessage(sender common.Address, destination common.Address, data CrossChainData, endChain string) (*types.Transaction, error) {
	return _Portabi.Contract.OutboundMessage(&_Portabi.TransactOpts, sender, destination, data, endChain)
}

// OutboundMessage is a paid mutator transaction binding the contract method 0x73a6eb14.
//
// Solidity: function outboundMessage(address sender, address destination, (address[],uint256[],string[],bool[]) data, string endChain) payable returns()
func (_Portabi *PortabiTransactorSession) OutboundMessage(sender common.Address, destination common.Address, data CrossChainData, endChain string) (*types.Transaction, error) {
	return _Portabi.Contract.OutboundMessage(&_Portabi.TransactOpts, sender, destination, data, endChain)
}

// ProxyConstructor is a paid mutator transaction binding the contract method 0x651f78f9.
//
// Solidity: function proxyConstructor(string _startChain, uint256 _chainId, string _name, string _symbol, address entryAddress, uint256 entryFee) returns()
func (_Portabi *PortabiTransactor) ProxyConstructor(opts *bind.TransactOpts, _startChain string, _chainId *big.Int, _name string, _symbol string, entryAddress common.Address, entryFee *big.Int) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "proxyConstructor", _startChain, _chainId, _name, _symbol, entryAddress, entryFee)
}

// ProxyConstructor is a paid mutator transaction binding the contract method 0x651f78f9.
//
// Solidity: function proxyConstructor(string _startChain, uint256 _chainId, string _name, string _symbol, address entryAddress, uint256 entryFee) returns()
func (_Portabi *PortabiSession) ProxyConstructor(_startChain string, _chainId *big.Int, _name string, _symbol string, entryAddress common.Address, entryFee *big.Int) (*types.Transaction, error) {
	return _Portabi.Contract.ProxyConstructor(&_Portabi.TransactOpts, _startChain, _chainId, _name, _symbol, entryAddress, entryFee)
}

// ProxyConstructor is a paid mutator transaction binding the contract method 0x651f78f9.
//
// Solidity: function proxyConstructor(string _startChain, uint256 _chainId, string _name, string _symbol, address entryAddress, uint256 entryFee) returns()
func (_Portabi *PortabiTransactorSession) ProxyConstructor(_startChain string, _chainId *big.Int, _name string, _symbol string, entryAddress common.Address, entryFee *big.Int) (*types.Transaction, error) {
	return _Portabi.Contract.ProxyConstructor(&_Portabi.TransactOpts, _startChain, _chainId, _name, _symbol, entryAddress, entryFee)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address ownerToRemove) returns()
func (_Portabi *PortabiTransactor) RemoveOwner(opts *bind.TransactOpts, ownerToRemove common.Address) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "removeOwner", ownerToRemove)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address ownerToRemove) returns()
func (_Portabi *PortabiSession) RemoveOwner(ownerToRemove common.Address) (*types.Transaction, error) {
	return _Portabi.Contract.RemoveOwner(&_Portabi.TransactOpts, ownerToRemove)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address ownerToRemove) returns()
func (_Portabi *PortabiTransactorSession) RemoveOwner(ownerToRemove common.Address) (*types.Transaction, error) {
	return _Portabi.Contract.RemoveOwner(&_Portabi.TransactOpts, ownerToRemove)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Portabi *PortabiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Portabi *PortabiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Portabi.Contract.RenounceOwnership(&_Portabi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Portabi *PortabiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Portabi.Contract.RenounceOwnership(&_Portabi.TransactOpts)
}

// SetChainId is a paid mutator transaction binding the contract method 0xef0e2ff4.
//
// Solidity: function setChainId(uint256 _chainId) returns()
func (_Portabi *PortabiTransactor) SetChainId(opts *bind.TransactOpts, _chainId *big.Int) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "setChainId", _chainId)
}

// SetChainId is a paid mutator transaction binding the contract method 0xef0e2ff4.
//
// Solidity: function setChainId(uint256 _chainId) returns()
func (_Portabi *PortabiSession) SetChainId(_chainId *big.Int) (*types.Transaction, error) {
	return _Portabi.Contract.SetChainId(&_Portabi.TransactOpts, _chainId)
}

// SetChainId is a paid mutator transaction binding the contract method 0xef0e2ff4.
//
// Solidity: function setChainId(uint256 _chainId) returns()
func (_Portabi *PortabiTransactorSession) SetChainId(_chainId *big.Int) (*types.Transaction, error) {
	return _Portabi.Contract.SetChainId(&_Portabi.TransactOpts, _chainId)
}

// SetDistributionContract is a paid mutator transaction binding the contract method 0xe7a764ef.
//
// Solidity: function setDistributionContract(address _contract) returns()
func (_Portabi *PortabiTransactor) SetDistributionContract(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "setDistributionContract", _contract)
}

// SetDistributionContract is a paid mutator transaction binding the contract method 0xe7a764ef.
//
// Solidity: function setDistributionContract(address _contract) returns()
func (_Portabi *PortabiSession) SetDistributionContract(_contract common.Address) (*types.Transaction, error) {
	return _Portabi.Contract.SetDistributionContract(&_Portabi.TransactOpts, _contract)
}

// SetDistributionContract is a paid mutator transaction binding the contract method 0xe7a764ef.
//
// Solidity: function setDistributionContract(address _contract) returns()
func (_Portabi *PortabiTransactorSession) SetDistributionContract(_contract common.Address) (*types.Transaction, error) {
	return _Portabi.Contract.SetDistributionContract(&_Portabi.TransactOpts, _contract)
}

// SetEntryFees is a paid mutator transaction binding the contract method 0x1d5ec885.
//
// Solidity: function setEntryFees(address[] _address, uint256[] _fee) returns()
func (_Portabi *PortabiTransactor) SetEntryFees(opts *bind.TransactOpts, _address []common.Address, _fee []*big.Int) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "setEntryFees", _address, _fee)
}

// SetEntryFees is a paid mutator transaction binding the contract method 0x1d5ec885.
//
// Solidity: function setEntryFees(address[] _address, uint256[] _fee) returns()
func (_Portabi *PortabiSession) SetEntryFees(_address []common.Address, _fee []*big.Int) (*types.Transaction, error) {
	return _Portabi.Contract.SetEntryFees(&_Portabi.TransactOpts, _address, _fee)
}

// SetEntryFees is a paid mutator transaction binding the contract method 0x1d5ec885.
//
// Solidity: function setEntryFees(address[] _address, uint256[] _fee) returns()
func (_Portabi *PortabiTransactorSession) SetEntryFees(_address []common.Address, _fee []*big.Int) (*types.Transaction, error) {
	return _Portabi.Contract.SetEntryFees(&_Portabi.TransactOpts, _address, _fee)
}

// SetPriceMapping is a paid mutator transaction binding the contract method 0x4f781c59.
//
// Solidity: function setPriceMapping(string startChain, string endChain, uint256 price) returns()
func (_Portabi *PortabiTransactor) SetPriceMapping(opts *bind.TransactOpts, startChain string, endChain string, price *big.Int) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "setPriceMapping", startChain, endChain, price)
}

// SetPriceMapping is a paid mutator transaction binding the contract method 0x4f781c59.
//
// Solidity: function setPriceMapping(string startChain, string endChain, uint256 price) returns()
func (_Portabi *PortabiSession) SetPriceMapping(startChain string, endChain string, price *big.Int) (*types.Transaction, error) {
	return _Portabi.Contract.SetPriceMapping(&_Portabi.TransactOpts, startChain, endChain, price)
}

// SetPriceMapping is a paid mutator transaction binding the contract method 0x4f781c59.
//
// Solidity: function setPriceMapping(string startChain, string endChain, uint256 price) returns()
func (_Portabi *PortabiTransactorSession) SetPriceMapping(startChain string, endChain string, price *big.Int) (*types.Transaction, error) {
	return _Portabi.Contract.SetPriceMapping(&_Portabi.TransactOpts, startChain, endChain, price)
}

// TestEvent is a paid mutator transaction binding the contract method 0x4f9d719e.
//
// Solidity: function testEvent() returns()
func (_Portabi *PortabiTransactor) TestEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "testEvent")
}

// TestEvent is a paid mutator transaction binding the contract method 0x4f9d719e.
//
// Solidity: function testEvent() returns()
func (_Portabi *PortabiSession) TestEvent() (*types.Transaction, error) {
	return _Portabi.Contract.TestEvent(&_Portabi.TransactOpts)
}

// TestEvent is a paid mutator transaction binding the contract method 0x4f9d719e.
//
// Solidity: function testEvent() returns()
func (_Portabi *PortabiTransactorSession) TestEvent() (*types.Transaction, error) {
	return _Portabi.Contract.TestEvent(&_Portabi.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Portabi *PortabiTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Portabi *PortabiSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Portabi.Contract.Transfer(&_Portabi.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Portabi *PortabiTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Portabi.Contract.Transfer(&_Portabi.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Portabi *PortabiTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Portabi *PortabiSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Portabi.Contract.TransferFrom(&_Portabi.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Portabi *PortabiTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Portabi.Contract.TransferFrom(&_Portabi.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Portabi *PortabiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Portabi *PortabiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Portabi.Contract.TransferOwnership(&_Portabi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Portabi *PortabiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Portabi.Contract.TransferOwnership(&_Portabi.TransactOpts, newOwner)
}

// UpdateCode is a paid mutator transaction binding the contract method 0x46951954.
//
// Solidity: function updateCode(address newCode) returns()
func (_Portabi *PortabiTransactor) UpdateCode(opts *bind.TransactOpts, newCode common.Address) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "updateCode", newCode)
}

// UpdateCode is a paid mutator transaction binding the contract method 0x46951954.
//
// Solidity: function updateCode(address newCode) returns()
func (_Portabi *PortabiSession) UpdateCode(newCode common.Address) (*types.Transaction, error) {
	return _Portabi.Contract.UpdateCode(&_Portabi.TransactOpts, newCode)
}

// UpdateCode is a paid mutator transaction binding the contract method 0x46951954.
//
// Solidity: function updateCode(address newCode) returns()
func (_Portabi *PortabiTransactorSession) UpdateCode(newCode common.Address) (*types.Transaction, error) {
	return _Portabi.Contract.UpdateCode(&_Portabi.TransactOpts, newCode)
}

// VoteNewOwner is a paid mutator transaction binding the contract method 0x7aa6a7f6.
//
// Solidity: function voteNewOwner(address owner) returns()
func (_Portabi *PortabiTransactor) VoteNewOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Portabi.contract.Transact(opts, "voteNewOwner", owner)
}

// VoteNewOwner is a paid mutator transaction binding the contract method 0x7aa6a7f6.
//
// Solidity: function voteNewOwner(address owner) returns()
func (_Portabi *PortabiSession) VoteNewOwner(owner common.Address) (*types.Transaction, error) {
	return _Portabi.Contract.VoteNewOwner(&_Portabi.TransactOpts, owner)
}

// VoteNewOwner is a paid mutator transaction binding the contract method 0x7aa6a7f6.
//
// Solidity: function voteNewOwner(address owner) returns()
func (_Portabi *PortabiTransactorSession) VoteNewOwner(owner common.Address) (*types.Transaction, error) {
	return _Portabi.Contract.VoteNewOwner(&_Portabi.TransactOpts, owner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Portabi *PortabiTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Portabi.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Portabi *PortabiSession) Receive() (*types.Transaction, error) {
	return _Portabi.Contract.Receive(&_Portabi.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Portabi *PortabiTransactorSession) Receive() (*types.Transaction, error) {
	return _Portabi.Contract.Receive(&_Portabi.TransactOpts)
}

// PortabiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Portabi contract.
type PortabiApprovalIterator struct {
	Event *PortabiApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PortabiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortabiApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PortabiApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PortabiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortabiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortabiApproval represents a Approval event raised by the Portabi contract.
type PortabiApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Portabi *PortabiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PortabiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Portabi.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PortabiApprovalIterator{contract: _Portabi.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Portabi *PortabiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PortabiApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Portabi.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortabiApproval)
				if err := _Portabi.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Portabi *PortabiFilterer) ParseApproval(log types.Log) (*PortabiApproval, error) {
	event := new(PortabiApproval)
	if err := _Portabi.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortabiBridgeSwapInDataIterator is returned from FilterBridgeSwapInData and is used to iterate over the raw logs and unpacked data for BridgeSwapInData events raised by the Portabi contract.
type PortabiBridgeSwapInDataIterator struct {
	Event *PortabiBridgeSwapInData // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PortabiBridgeSwapInDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortabiBridgeSwapInData)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PortabiBridgeSwapInData)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PortabiBridgeSwapInDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortabiBridgeSwapInDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortabiBridgeSwapInData represents a BridgeSwapInData event raised by the Portabi contract.
type PortabiBridgeSwapInData struct {
	StartChain  string
	Sender      common.Address
	Destination common.Address
	Data        CrossChainData
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBridgeSwapInData is a free log retrieval operation binding the contract event 0xd0ad1f9d28a2932b721e7ccc895db127b1e935262cf71197034cbbcfbc714872.
//
// Solidity: event BridgeSwapInData(string startChain, address sender, address destination, (address[],uint256[],string[],bool[]) data)
func (_Portabi *PortabiFilterer) FilterBridgeSwapInData(opts *bind.FilterOpts) (*PortabiBridgeSwapInDataIterator, error) {

	logs, sub, err := _Portabi.contract.FilterLogs(opts, "BridgeSwapInData")
	if err != nil {
		return nil, err
	}
	return &PortabiBridgeSwapInDataIterator{contract: _Portabi.contract, event: "BridgeSwapInData", logs: logs, sub: sub}, nil
}

// WatchBridgeSwapInData is a free log subscription operation binding the contract event 0xd0ad1f9d28a2932b721e7ccc895db127b1e935262cf71197034cbbcfbc714872.
//
// Solidity: event BridgeSwapInData(string startChain, address sender, address destination, (address[],uint256[],string[],bool[]) data)
func (_Portabi *PortabiFilterer) WatchBridgeSwapInData(opts *bind.WatchOpts, sink chan<- *PortabiBridgeSwapInData) (event.Subscription, error) {

	logs, sub, err := _Portabi.contract.WatchLogs(opts, "BridgeSwapInData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortabiBridgeSwapInData)
				if err := _Portabi.contract.UnpackLog(event, "BridgeSwapInData", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgeSwapInData is a log parse operation binding the contract event 0xd0ad1f9d28a2932b721e7ccc895db127b1e935262cf71197034cbbcfbc714872.
//
// Solidity: event BridgeSwapInData(string startChain, address sender, address destination, (address[],uint256[],string[],bool[]) data)
func (_Portabi *PortabiFilterer) ParseBridgeSwapInData(log types.Log) (*PortabiBridgeSwapInData, error) {
	event := new(PortabiBridgeSwapInData)
	if err := _Portabi.contract.UnpackLog(event, "BridgeSwapInData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortabiBridgeSwapOutDataIterator is returned from FilterBridgeSwapOutData and is used to iterate over the raw logs and unpacked data for BridgeSwapOutData events raised by the Portabi contract.
type PortabiBridgeSwapOutDataIterator struct {
	Event *PortabiBridgeSwapOutData // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PortabiBridgeSwapOutDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortabiBridgeSwapOutData)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PortabiBridgeSwapOutData)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PortabiBridgeSwapOutDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortabiBridgeSwapOutDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortabiBridgeSwapOutData represents a BridgeSwapOutData event raised by the Portabi contract.
type PortabiBridgeSwapOutData struct {
	Sender         common.Address
	StartChain     string
	EndChain       string
	TransferAmount *big.Int
	Trigger        common.Address
	Data           CrossChainData
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeSwapOutData is a free log retrieval operation binding the contract event 0x51e93e11b0ffb8b5be5f72e09818d40206eb8eccfcf2e6e39bab2e28d386a09c.
//
// Solidity: event BridgeSwapOutData(address sender, string startChain, string endChain, uint256 transferAmount, address trigger, (address[],uint256[],string[],bool[]) data)
func (_Portabi *PortabiFilterer) FilterBridgeSwapOutData(opts *bind.FilterOpts) (*PortabiBridgeSwapOutDataIterator, error) {

	logs, sub, err := _Portabi.contract.FilterLogs(opts, "BridgeSwapOutData")
	if err != nil {
		return nil, err
	}
	return &PortabiBridgeSwapOutDataIterator{contract: _Portabi.contract, event: "BridgeSwapOutData", logs: logs, sub: sub}, nil
}

// WatchBridgeSwapOutData is a free log subscription operation binding the contract event 0x51e93e11b0ffb8b5be5f72e09818d40206eb8eccfcf2e6e39bab2e28d386a09c.
//
// Solidity: event BridgeSwapOutData(address sender, string startChain, string endChain, uint256 transferAmount, address trigger, (address[],uint256[],string[],bool[]) data)
func (_Portabi *PortabiFilterer) WatchBridgeSwapOutData(opts *bind.WatchOpts, sink chan<- *PortabiBridgeSwapOutData) (event.Subscription, error) {

	logs, sub, err := _Portabi.contract.WatchLogs(opts, "BridgeSwapOutData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortabiBridgeSwapOutData)
				if err := _Portabi.contract.UnpackLog(event, "BridgeSwapOutData", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgeSwapOutData is a log parse operation binding the contract event 0x51e93e11b0ffb8b5be5f72e09818d40206eb8eccfcf2e6e39bab2e28d386a09c.
//
// Solidity: event BridgeSwapOutData(address sender, string startChain, string endChain, uint256 transferAmount, address trigger, (address[],uint256[],string[],bool[]) data)
func (_Portabi *PortabiFilterer) ParseBridgeSwapOutData(log types.Log) (*PortabiBridgeSwapOutData, error) {
	event := new(PortabiBridgeSwapOutData)
	if err := _Portabi.contract.UnpackLog(event, "BridgeSwapOutData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortabiNewSignerIterator is returned from FilterNewSigner and is used to iterate over the raw logs and unpacked data for NewSigner events raised by the Portabi contract.
type PortabiNewSignerIterator struct {
	Event *PortabiNewSigner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PortabiNewSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortabiNewSigner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PortabiNewSigner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PortabiNewSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortabiNewSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortabiNewSigner represents a NewSigner event raised by the Portabi contract.
type PortabiNewSigner struct {
	Signer common.Address
	Domain string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewSigner is a free log retrieval operation binding the contract event 0x065514f3342dbf18a95c284b60121ac59a6b20a65247a67c3622db5ee95cf65e.
//
// Solidity: event NewSigner(address signer, string domain)
func (_Portabi *PortabiFilterer) FilterNewSigner(opts *bind.FilterOpts) (*PortabiNewSignerIterator, error) {

	logs, sub, err := _Portabi.contract.FilterLogs(opts, "NewSigner")
	if err != nil {
		return nil, err
	}
	return &PortabiNewSignerIterator{contract: _Portabi.contract, event: "NewSigner", logs: logs, sub: sub}, nil
}

// WatchNewSigner is a free log subscription operation binding the contract event 0x065514f3342dbf18a95c284b60121ac59a6b20a65247a67c3622db5ee95cf65e.
//
// Solidity: event NewSigner(address signer, string domain)
func (_Portabi *PortabiFilterer) WatchNewSigner(opts *bind.WatchOpts, sink chan<- *PortabiNewSigner) (event.Subscription, error) {

	logs, sub, err := _Portabi.contract.WatchLogs(opts, "NewSigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortabiNewSigner)
				if err := _Portabi.contract.UnpackLog(event, "NewSigner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewSigner is a log parse operation binding the contract event 0x065514f3342dbf18a95c284b60121ac59a6b20a65247a67c3622db5ee95cf65e.
//
// Solidity: event NewSigner(address signer, string domain)
func (_Portabi *PortabiFilterer) ParseNewSigner(log types.Log) (*PortabiNewSigner, error) {
	event := new(PortabiNewSigner)
	if err := _Portabi.contract.UnpackLog(event, "NewSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortabiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Portabi contract.
type PortabiOwnershipTransferredIterator struct {
	Event *PortabiOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PortabiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortabiOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PortabiOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PortabiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortabiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortabiOwnershipTransferred represents a OwnershipTransferred event raised by the Portabi contract.
type PortabiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Portabi *PortabiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PortabiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Portabi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PortabiOwnershipTransferredIterator{contract: _Portabi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Portabi *PortabiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PortabiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Portabi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortabiOwnershipTransferred)
				if err := _Portabi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Portabi *PortabiFilterer) ParseOwnershipTransferred(log types.Log) (*PortabiOwnershipTransferred, error) {
	event := new(PortabiOwnershipTransferred)
	if err := _Portabi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortabiTestEmitIterator is returned from FilterTestEmit and is used to iterate over the raw logs and unpacked data for TestEmit events raised by the Portabi contract.
type PortabiTestEmitIterator struct {
	Event *PortabiTestEmit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PortabiTestEmitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortabiTestEmit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PortabiTestEmit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PortabiTestEmitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortabiTestEmitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortabiTestEmit represents a TestEmit event raised by the Portabi contract.
type PortabiTestEmit struct {
	Message string
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTestEmit is a free log retrieval operation binding the contract event 0x06f2e9f9aaec2446f56073dfe76a380edde4d9e629dfd95ca95c2ebaca7ddabb.
//
// Solidity: event TestEmit(string message, address sender)
func (_Portabi *PortabiFilterer) FilterTestEmit(opts *bind.FilterOpts) (*PortabiTestEmitIterator, error) {

	logs, sub, err := _Portabi.contract.FilterLogs(opts, "TestEmit")
	if err != nil {
		return nil, err
	}
	return &PortabiTestEmitIterator{contract: _Portabi.contract, event: "TestEmit", logs: logs, sub: sub}, nil
}

// WatchTestEmit is a free log subscription operation binding the contract event 0x06f2e9f9aaec2446f56073dfe76a380edde4d9e629dfd95ca95c2ebaca7ddabb.
//
// Solidity: event TestEmit(string message, address sender)
func (_Portabi *PortabiFilterer) WatchTestEmit(opts *bind.WatchOpts, sink chan<- *PortabiTestEmit) (event.Subscription, error) {

	logs, sub, err := _Portabi.contract.WatchLogs(opts, "TestEmit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortabiTestEmit)
				if err := _Portabi.contract.UnpackLog(event, "TestEmit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTestEmit is a log parse operation binding the contract event 0x06f2e9f9aaec2446f56073dfe76a380edde4d9e629dfd95ca95c2ebaca7ddabb.
//
// Solidity: event TestEmit(string message, address sender)
func (_Portabi *PortabiFilterer) ParseTestEmit(log types.Log) (*PortabiTestEmit, error) {
	event := new(PortabiTestEmit)
	if err := _Portabi.contract.UnpackLog(event, "TestEmit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortabiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Portabi contract.
type PortabiTransferIterator struct {
	Event *PortabiTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PortabiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortabiTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PortabiTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PortabiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortabiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortabiTransfer represents a Transfer event raised by the Portabi contract.
type PortabiTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Portabi *PortabiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PortabiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Portabi.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PortabiTransferIterator{contract: _Portabi.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Portabi *PortabiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PortabiTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Portabi.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortabiTransfer)
				if err := _Portabi.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Portabi *PortabiFilterer) ParseTransfer(log types.Log) (*PortabiTransfer, error) {
	event := new(PortabiTransfer)
	if err := _Portabi.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
