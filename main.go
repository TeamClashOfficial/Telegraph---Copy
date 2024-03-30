package main

import (
	"fmt"
	"net/http"
	"telegraph/cmd"
	"telegraph/config"
	diContainer "telegraph/container"
	"telegraph/log"
	blockchainmodule "telegraph/modules/blockchain"
	networkmodule "telegraph/modules/network"
	statusmodule "telegraph/modules/status"
	transactionmodule "telegraph/modules/transaction"
	validatormodule "telegraph/modules/validator"
	"telegraph/pkg/chat"
	"telegraph/pkg/listener"
	"telegraph/pkg/register"
	"telegraph/pkg/signer"
	"telegraph/pkg/signer/auth"
	"telegraph/pkg/wallet"

	"github.com/gorilla/mux"
	"go.uber.org/dig"
)

type MainContainerInvoke struct {
	dig.In

	networkmodule.NetworkRouter
	*mux.Router
	networkmodule.NetworkService
	wallet.Wallet
	statusmodule.StatusRouter
	transactionmodule.TransactionRouter
	listener.Listener
	validatormodule.ValidatorRouter
	http.Handler
	blockchainmodule.BlockchainRouter
}

func main() {

	container := bootstrap()
	err := container.Invoke(func(
		containerInvoke MainContainerInvoke,
	) {
		register.CheckForCoreContract(config.Conf, containerInvoke.NetworkService)

		isSigner := register.IsNodeSigner(config.Conf, containerInvoke.Wallet)
		fmt.Println("NODE IS SIGNER: ", isSigner)
		if isSigner {
			go func() {
				cmd.Execute()
				containerInvoke.Listener.StartListener()
			}()
		}

		hub := chat.NewHub()
		go hub.Run()

		containerInvoke.NetworkRouter.RegisterRoutes()
		containerInvoke.StatusRouter.RegisterRoutes()
		containerInvoke.TransactionRouter.RegisterRoutes()
		containerInvoke.ValidatorRouter.RegisterRoutes()
		containerInvoke.BlockchainRouter.RegisterRoutes()

		containerInvoke.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
			chat.ServeWs(hub, w, r)
		})

		log.Info("Serving on: ", config.Conf.IP)
		log.Fatal(http.ListenAndServe(":"+config.Conf.PORT, containerInvoke.Handler))
	})

	if err != nil {
		panic(err)
	}

}

func bootstrap() *dig.Container {
	container := diContainer.GetContainer()
	networkmodule.NetworkModule()
	statusmodule.StatusModule()
	transactionmodule.TransactionModule()
	validatormodule.ValidatorModule()
	blockchainmodule.BlockchainModule()

	container.Provide(auth.NewAuth)
	container.Provide(wallet.NewWallet)
	container.Provide(signer.NewSigner)
	container.Provide(listener.NewListener)
	return container
}
