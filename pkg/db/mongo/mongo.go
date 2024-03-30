package mongoconnect

import (
	"context"
	"telegraph/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB(network string) *mongo.Collection {
	// if err := godotenv.Load(); err != nil {
	// 	log.Println("No .env file found")
	// }

	// uri := os.Getenv("MONGODB_URI")
	// if uri == "" {
	// 	log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	// }
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Conf.DbURL))
	if err != nil {
		panic(err)
	}

	return client.Database("telegraph").Collection(network)
}

func ConnectDB(database string) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Conf.DbURL))
	if err != nil {
		panic(err)
	}

	return client.Database(database)
}
