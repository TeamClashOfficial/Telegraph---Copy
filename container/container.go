package container

import (
	"net/http"
	"sync"
	mongoconnect "telegraph/pkg/db/mongo"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

var (
	container *dig.Container
	once      sync.Once
)

func initContainer() {
	once.Do(func() {
		ethClient, err := ethclient.Dial("https://rinkeby.infura.io/v3/b47d4c6efe59493cbf6ba034f07718ae")
		if err != nil {
			panic(err)
		}
		db := mongoconnect.ConnectDB("telegraph")
		router := mux.NewRouter()
		router.UseEncodedPath()
		handler := cors.Default().Handler(router)

		container = dig.New()
		container.Provide(db)
		container.Provide(router)

		container.Provide(func() *mux.Router { return router })
		container.Provide(func() *mongo.Database { return db })
		container.Provide(func() *ethclient.Client { return ethClient })
		container.Provide(func() http.Handler { return handler })
	})
}

func GetContainer() *dig.Container {
	initContainer()
	return container
}
