package validator

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"

	mongoconnect "telegraph/pkg/db/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Validator struct {
	PublicETHAddress common.Address `bson:"publicethaddress,omitempty"`
	Moniker          string         `bson:"moniker,omitempty"`
	LastPingTime     time.Time      `bson:"lastpingtime,omitempty"`
	Domain           string         `bson:"domain,omitempty"`
	IsMine           bool           `bson:"ismine,omitempty"`
}

// Import DB connection.
var coll *mongo.Collection = mongoconnect.ConnectToDB("validator")

func GetValidators() []Validator {
	results, err := coll.Find(context.TODO(), bson.D{{}})
	if errors.Is(err, mongo.ErrNoDocuments) {
		fmt.Printf("No Validators found")
		return nil
	}
	if err != nil {
		panic(err)
	}
	decoded := []Validator{}
	for results.Next(context.TODO()) {
		var elem Validator
		err := results.Decode(&elem)
		if err != nil {
			log.Fatal("results.Decode error:", err)
		}
		if elem.Domain != "" && elem.Moniker != "" { // to prevent adding empty validator
			decoded = append(decoded, elem)
		}
	}
	return decoded
}

func FindValidatorByAddress(address common.Address) (Validator, error) {
	var validator Validator
	filter := bson.D{primitive.E{Key: "publicethaddress", Value: address}}
	opts := options.FindOne().SetSort(bson.D{primitive.E{Key: "publicethaddress", Value: 1}})
	err := coll.FindOne(context.TODO(), filter, opts).Decode(&validator)
	if err != nil {
		return validator, err
	}
	return validator, err
}

func FindMyValidator() (Validator, error) {
	var validator Validator
	filter := bson.D{primitive.E{Key: "ismine", Value: true}}
	opts := options.FindOne().SetSort(bson.D{primitive.E{Key: "ismine", Value: 1}})
	err := coll.FindOne(context.TODO(), filter, opts).Decode(&validator)
	if err != nil {
		return validator, err
	}
	return validator, err
}

func AddValidator(publicAddress common.Address, moniker, domain string, isMine bool) error {
	ctx := context.Background()
	validator := Validator{
		PublicETHAddress: publicAddress,
		Moniker:          moniker,
		LastPingTime:     time.Now(),
		Domain:           domain,
		IsMine:           isMine,
	}

	filter := bson.D{primitive.E{Key: "domain", Value: domain}}
	opts := options.FindOne().SetSort(bson.D{primitive.E{Key: "domain", Value: 1}})
	err := coll.FindOne(context.TODO(), filter, opts).Decode(&validator)
	if err == nil { // validator already exists in the database
		return err
	}

	_, err = coll.InsertOne(ctx, validator)
	return err
}

func RemoveValidator() {

}

func NotifyValidator() {

}
