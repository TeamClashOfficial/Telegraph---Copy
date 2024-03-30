package blockchainmodule

import (
	"encoding/json"
	"net/http"
)

type BlockchainController interface {
	CreateNewWallet(w http.ResponseWriter, r *http.Request)
	IsWalletSigner(w http.ResponseWriter, r *http.Request)
	StartRegistration(w http.ResponseWriter, r *http.Request)
	GetFees(w http.ResponseWriter, r *http.Request)
}

type BlockchainControllerImpl struct {
	blockchainService BlockchainService
}

func (controller *BlockchainControllerImpl) CreateNewWallet(w http.ResponseWriter, r *http.Request) {
	newWallet, err := controller.blockchainService.CreateNewWallet()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_ = json.NewEncoder(w).Encode(newWallet)
}

func (controller *BlockchainControllerImpl) IsWalletSigner(w http.ResponseWriter, r *http.Request) {
	isSigner, err := controller.blockchainService.IsWalletSigner()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_ = json.NewEncoder(w).Encode(isSigner)
}

func (controller *BlockchainControllerImpl) GetFees(w http.ResponseWriter, r *http.Request) {
	fees, err := controller.blockchainService.GetFees(w, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_ = json.NewEncoder(w).Encode(fees)
}

func (controller *BlockchainControllerImpl) StartRegistration(w http.ResponseWriter, r *http.Request) {
	controller.blockchainService.StartRegistration(w, r)
}

func NewBlockchainController(blockchainService BlockchainService) BlockchainController {
	return &BlockchainControllerImpl{
		blockchainService: blockchainService,
	}
}
