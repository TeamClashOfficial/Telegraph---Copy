package validatormodule

import (
	"fmt"
	"telegraph/container"
	"telegraph/log"
)

func ValidatorModule() {
	fmt.Println("Starting validator module")

	container := container.GetContainer()

	modules := []interface{}{
		NewValidatorController,
		NewValidatorRouter,
		NewValidatorService,
		NewValidatorRepository,
	}

	// print out the modules
	fmt.Println("Loading validator module")

	for k := range modules {
		if err := container.Provide(modules[k]); err != nil {
			log.Error("error loading container module: ", err)
		}
	}
}
