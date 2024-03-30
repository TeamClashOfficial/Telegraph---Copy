package dbaccess

import (
	"time"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
)

// Order Schema
type Order struct {
	CurrentNode     int       `bson:"currentNode,omitempty"`
	OrderHeight     int       `bson:"orderHeight,omitempty"`
	TimeSinceSwitch time.Time `bson:"timeSinceSwitch,omitempty"`
}

// Util to get order collection
func OrderCollection() *mongo.Collection {
	return s.Database.Collection("Orders")
}

func CreateOrder(order *Order) (r *mongo.InsertOneResult, err error) {
	r, err = OrderCollection().InsertOne(s.Ctx, order)
	return
}

func FindOrders(order *Order) (orders []Order) {
	cursor, err := OrderCollection().Find(s.Ctx, order)
	if err != nil {
		log.Error("Error occured while creating the cursor", err)
		return
	}
	if err = cursor.All(s.Ctx, &orders); err != nil {
		log.Error("Error occured while un-marshal the cursor", err)
		return
	}
	return
}
