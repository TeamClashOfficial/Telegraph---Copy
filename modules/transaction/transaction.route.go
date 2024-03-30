package transactionmodule

import "github.com/gorilla/mux"

type TransactionRouter interface {
	RegisterRoutes()
	GetRouter() *mux.Router
}

type TransactionRouterImpl struct {
	muxRouter             *mux.Router
	TransactionController TransactionController
}

func (router *TransactionRouterImpl) RegisterRoutes() {
	router.muxRouter.HandleFunc("/", router.TransactionController.GetTransactions)
	router.muxRouter.HandleFunc("/create", router.TransactionController.CreateTransaction).Methods("POST")
	router.muxRouter.HandleFunc("/update", router.TransactionController.UpdateTransaction).Methods("PUT")
	router.muxRouter.HandleFunc("/finalize", router.TransactionController.FinalizeTransaction).Methods("POST")
}

func (router *TransactionRouterImpl) GetRouter() *mux.Router {
	return router.muxRouter
}

func NewTransactionRouter(muxRouter *mux.Router, transactionController TransactionController) TransactionRouter {
	return &TransactionRouterImpl{
		muxRouter:             muxRouter,
		TransactionController: transactionController,
	}
}
