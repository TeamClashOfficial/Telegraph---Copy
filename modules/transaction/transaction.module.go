package transactionmodule

import (
	"telegraph/container"
	"telegraph/log"
)

func TransactionModule() {
	container := container.GetContainer()

	modules := []interface{}{
		NewTransactionService,
		NewTransactionController,
		NewTransactionRouter,
		NewTransactionRepository,
	}

	for k := range modules {
		if err := container.Provide(modules[k]); err != nil {
			log.Error("error loading container module: ", err)
		}
	}
}
