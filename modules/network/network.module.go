package networkmodule

import (
	"telegraph/container"
	"telegraph/log"
)

func NetworkModule() {
	container := container.GetContainer()

	modules := []interface{}{
		NewNetworkRepository,
		NewNetworkService,
		NewNetworkController,
		newNetworkRouter,
	}

	for k := range modules {
		if err := container.Provide(modules[k]); err != nil {
			log.Error("error loading container module: ", err)
		}
	}
}
