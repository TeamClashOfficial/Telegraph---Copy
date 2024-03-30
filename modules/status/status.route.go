package statusmodule

import "github.com/gorilla/mux"

type StatusRouter interface {
	RegisterRoutes()
	GetRouter() *mux.Router
}

type StatusRouterImpl struct {
	muxRouter        *mux.Router
	StatusController StatusController
}

func (router *StatusRouterImpl) RegisterRoutes() {
	router.muxRouter.HandleFunc("/status", router.StatusController.Status).Methods("GET")
}

func (router *StatusRouterImpl) GetRouter() *mux.Router {
	return router.muxRouter
}

func NewStatusRouter(muxRouter *mux.Router, statusController StatusController) StatusRouter {
	return &StatusRouterImpl{
		muxRouter:        muxRouter,
		StatusController: statusController,
	}
}
