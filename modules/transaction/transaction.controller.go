package transactionmodule

import (
	"encoding/json"
	"io"
	"net/http"

	"telegraph/pkg/wallet"
	"telegraph/tools"
)

type TransactionController interface {
	GetTransactions(w http.ResponseWriter, r *http.Request)
	CreateTransaction(w http.ResponseWriter, r *http.Request)
	UpdateTransaction(w http.ResponseWriter, r *http.Request)
	FinalizeTransaction(w http.ResponseWriter, r *http.Request)
}

type TransactionControllerImpl struct {
	transactionService TransactionService
	wallet             wallet.Wallet
}

func (controller *TransactionControllerImpl) GetTransactions(w http.ResponseWriter, r *http.Request) {
	result, err := controller.transactionService.GetTransactions()
	if err != nil {
		http.Error(w, "Error getting transactions: "+err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error marshaling JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonResp)
}

func (controller *TransactionControllerImpl) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(io.LimitReader(r.Body, 1048576))
	tools.Check(err)
	authHeader := r.Header.Get("Authorization")
	isValidator := controller.wallet.VerifySignature(authHeader)

	if !isValidator {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte("Unauthorized"))
		return
	}

	err = controller.transactionService.CreateFromBytes(bytes)
	if err != nil {
		http.Error(w, "Error creating transactions: "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write([]byte("Transaction Create started"))
}

func (controller *TransactionControllerImpl) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(io.LimitReader(r.Body, 1048576))
	tools.Check(err)
	authHeader := r.Header.Get("Authorization")
	isValidator := controller.wallet.VerifySignature(authHeader)

	if !isValidator {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte("Unauthorized"))
		return
	}

	err = controller.transactionService.UpdateFromBytes(bytes)
	if err != nil {
		http.Error(w, "Error updating transactions: "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write([]byte("Transaction Update started"))
}

func (controller *TransactionControllerImpl) FinalizeTransaction(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(io.LimitReader(r.Body, 1048576))
	tools.Check(err)
	authHeader := r.Header.Get("Authorization")
	isValidator := controller.wallet.VerifySignature(authHeader)

	if !isValidator {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte("Unauthorized"))
		return
	}

	err = controller.transactionService.FinalizeFromBytes(bytes)
	if err != nil {
		http.Error(w, "Error finalizing transactions: "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write([]byte("Transaction Finalize started"))
}

func NewTransactionController(transactionService TransactionService) TransactionController {
	return &TransactionControllerImpl{
		transactionService: transactionService,
		// wallet:             wallet,
	}
}
