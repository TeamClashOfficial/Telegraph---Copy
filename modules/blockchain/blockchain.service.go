package blockchainmodule

import (
	"log"
	"math/big"
	"net/http"
	"telegraph/config"
	networkmodule "telegraph/modules/network"
	validatormodule "telegraph/modules/validator"
	register "telegraph/pkg/register"
	wallet "telegraph/pkg/wallet"
)

type BlockchainService interface {
	CreateNewWallet() (wallet.Wallet, error)
	StartRegistration(w http.ResponseWriter, r *http.Request)
	IsWalletSigner() (bool, error)
	GetFees(w http.ResponseWriter, r *http.Request) (*big.Int, error)
}

type BlockchainServiceImpl struct {
	validatorService validatormodule.ValidatorService
}

func (blockchainService *BlockchainServiceImpl) CreateNewWallet() (wallet.Wallet, error) {
	// Call the NewWallet function from the wallet package
	newWallet := wallet.NewWallet(blockchainService.validatorService)

	return newWallet, nil // Return the new wallet and nil error
}

func (blockchainService *BlockchainServiceImpl) IsWalletSigner() (bool, error) {
	// Get the existing wallet or create a new one if it doesn't exist
	log.Println("Getting wallet")
	myWallet := wallet.NewWallet(blockchainService.validatorService)
	log.Println(myWallet)
	// Call the IsSigner function from the wallet package
	isSigner := register.IsNodeSigner(config.Conf, myWallet)

	return isSigner, nil // Return the isSigner boolean and nil error
}

func (blockchainService *BlockchainServiceImpl) GetFees(w http.ResponseWriter, r *http.Request) (*big.Int, error) {
	// Call the GetFees function from the wallet package
	var req register.RegisterRequest
	req.EntryFeeAddress = config.Conf.HokkETH
	fees := register.GetEntryFee(config.Conf, req)

	return fees, nil
}

func (blockchainService *BlockchainServiceImpl) StartRegistration(w http.ResponseWriter, r *http.Request) {
	// Parse the request data
	var req register.RegisterRequest
	req.EntryFeeAddress = config.Conf.HokkETH

	// Create a wallet and call the StartRegistration function
	myWallet := wallet.NewWallet(blockchainService.validatorService)
	register.StartRegistration(config.Conf, myWallet, req, &networkmodule.NetworkServiceImpl{})

	// Send a success response
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Registration started"))
}

func NewBlockchainService() BlockchainService {
	return &BlockchainServiceImpl{}
}
