package networkmodule

import (
	"github.com/gorilla/mux"
)

type NetworkRouter interface {
	RegisterRoutes()
	GetRouter() *mux.Router
}

type NetworkRouterImpl struct {
	muxRouter         *mux.Router
	NetworkController NetworkController
}

func (router *NetworkRouterImpl) RegisterRoutes() {
	router.muxRouter.HandleFunc("/network", router.NetworkController.GetNetworks).Methods("GET")
	router.muxRouter.HandleFunc("/network", router.NetworkController.AddNetwork).Methods("POST")
	router.muxRouter.HandleFunc("/network", router.NetworkController.UpdateNetwork).Methods("PUT")
	router.muxRouter.HandleFunc("/network", router.NetworkController.RemoveNetwork).Methods("DELETE")
}

func (router *NetworkRouterImpl) GetRouter() *mux.Router {
	return router.muxRouter
}

func newNetworkRouter(muxRouter *mux.Router, newNetworkController NetworkController) NetworkRouter {
	return &NetworkRouterImpl{
		muxRouter:         muxRouter,
		NetworkController: newNetworkController,
	}
}
