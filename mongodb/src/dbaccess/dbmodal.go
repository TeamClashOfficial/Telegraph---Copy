package dbaccess

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var s *Service

// Service ...
type Service struct {
	Client    *mongo.Client
	Ctx       context.Context
	CtxCancel context.CancelFunc
	Database  *mongo.Database
}

// SetupService ...
func SetupService(client *mongo.Client, ctx context.Context, cancel context.CancelFunc, db *mongo.Database) {
	s = &Service{
		Client:    client,
		Ctx:       ctx,
		CtxCancel: cancel,
		Database:  db,
	}
	// setup collections indexes
	setupIndexes()
}

func setupIndexes() {
	// add all schemas indexes here
	initValidatorIndexes()
}
