package blockchainmodule

import (
	"telegraph/container"
	"telegraph/log"
)

func BlockchainModule() {
	container := container.GetContainer()

	modules := []interface{}{
		NewBlockchainController,
		NewBlockchainService,
		NewBlockchainRouter,
	}

	for k := range modules {
		if err := container.Provide(modules[k]); err != nil {
			log.Error("error loading container module: ", err)
		}
	}
}
