package statusmodule

import (
	"telegraph/container"
	"telegraph/log"
)

func StatusModule() {
	container := container.GetContainer()

	modules := []interface{}{
		NewStatusService,
		NewStatusController,
		NewStatusRouter,
	}

	for k := range modules {
		if err := container.Provide(modules[k]); err != nil {
			log.Error("error loading container module: ", err)
		}
	}
}
