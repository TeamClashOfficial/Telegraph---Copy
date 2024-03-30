package blockchainmodule

import (
	"github.com/gorilla/mux"
	"net/http"
	"telegraph/limiter"
)

type BlockchainRouter interface {
	RegisterRoutes()
	GetRouter() *mux.Router
}

type BlockchainRouterImpl struct {
	muxRouter            *mux.Router
	BlockchainController BlockchainController
}

func (router *BlockchainRouterImpl) RegisterRoutes() {
	rl := limiter.GetRateLimiter()

	router.muxRouter.Handle("/", rl.RateLimit(http.HandlerFunc(router.BlockchainController.CreateNewWallet))).Methods("POST")
	router.muxRouter.HandleFunc("/signer", router.BlockchainController.IsWalletSigner).Methods("GET")
	router.muxRouter.HandleFunc("/signer", router.BlockchainController.StartRegistration).Methods("POST")
	router.muxRouter.HandleFunc("/fees", router.BlockchainController.GetFees).Methods("GET")
}

func (router *BlockchainRouterImpl) GetRouter() *mux.Router {
	return router.muxRouter
}

func NewBlockchainRouter(muxRouter *mux.Router, blockchainController BlockchainController) BlockchainRouter {
	return &BlockchainRouterImpl{
		muxRouter:            muxRouter,
		BlockchainController: blockchainController,
	}
}
