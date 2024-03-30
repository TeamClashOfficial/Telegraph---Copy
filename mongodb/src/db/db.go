package db

import (
	"context"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DBObject struct {
	Client    *mongo.Client
	Ctx       context.Context
	CtxCancel context.CancelFunc
	Database  *mongo.Database
}

func Init() *DBObject {
	db := &DBObject{}
	client, ctx, cancel, err := connect(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("error connecting to db", err)
	}
	if err := ping(client, ctx); err != nil {
		log.Fatal("error pinging db", err)
	}
	database := client.Database(os.Getenv("DB_NAME"))
	db = &DBObject{Client: client, Ctx: ctx, CtxCancel: cancel, Database: database}
	return db
}

func connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	return client, ctx, cancel, err
}

// This is a user defined method that accepts
// mongo.Client and context.Context
// This method used to ping the mongoDB, return error if any.
func ping(client *mongo.Client, ctx context.Context) error {

	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occurred, then
	// the error can be handled.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	log.Info("Db connected successfully")
	return nil
}

// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal("client.Disconnect error:", err)
		}
	}()
}
