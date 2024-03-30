package validatormodule

import (
	"fmt"

	"github.com/gorilla/mux"
)

type ValidatorRouter interface {
	RegisterRoutes()
	GetRouter() *mux.Router
}

type ValidatorRouterImpl struct {
	muxRouter           *mux.Router
	ValidatorController ValidatorController
}

func (router *ValidatorRouterImpl) RegisterRoutes() {
	fmt.Println("Registering validator routes")
	router.muxRouter.HandleFunc("/validator/", router.ValidatorController.GetMyValidator).Methods("GET")
	router.muxRouter.HandleFunc("/validator/params", router.ValidatorController.AddValidatorParams).Methods("POST")
	router.muxRouter.HandleFunc("/validator/params", router.ValidatorController.UpdateValidatorParams).Methods("PUT")
	router.muxRouter.HandleFunc("/validator/start", router.ValidatorController.StartValidator).Methods("POST")
	router.muxRouter.HandleFunc("/validator/sign", router.ValidatorController.SignValidator).Methods("POST")
	router.muxRouter.HandleFunc("/validator/update", router.ValidatorController.UpdateValidator).Methods("POST")
	router.muxRouter.HandleFunc("/validator/reshare", router.ValidatorController.ReshareValidator).Methods("POST")
}

func (router *ValidatorRouterImpl) GetRouter() *mux.Router {
	return router.muxRouter
}

func NewValidatorRouter(muxRouter *mux.Router, statusController ValidatorController) ValidatorRouter {
	return &ValidatorRouterImpl{
		muxRouter:           muxRouter,
		ValidatorController: statusController,
	}
}
