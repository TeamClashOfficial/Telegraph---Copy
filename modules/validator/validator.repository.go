package validatormodule

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Validator struct {
	PublicETHAddress common.Address `bson:"publicethaddress,omitempty" json:"publicethaddress"`
	Moniker          string         `bson:"moniker,omitempty" json:"moniker"`
	LastPingTime     time.Time      `bson:"lastpingtime,omitempty" json:"lastpingtime"`
	Domain           string         `bson:"domain,omitempty" json:"domain"`
	IsMine           bool           `bson:"ismine,omitempty" json:"ismine"`
}

type ValidatorRepository interface {
	FindValidators() ([]Validator, error)
	FindValidator(filter interface{}, opts ...*options.FindOneOptions) (Validator, error)
	InsertValidator(validator Validator) (*mongo.InsertOneResult, error)
}

type ValidatorRepositoryImpl struct {
	db *mongo.Database
}

const COLLECTION_NAME = "validator"

func (validatorRepository *ValidatorRepositoryImpl) FindValidators() ([]Validator, error) {
	results, err := validatorRepository.db.Collection(COLLECTION_NAME).Find(context.TODO(), bson.D{{}})
	if errors.Is(err, mongo.ErrNoDocuments) {
		fmt.Printf("No Validators found")
		return nil, nil
	}
	if err != nil {
		return nil, err
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
	return decoded, nil
}

func (validatorRepository *ValidatorRepositoryImpl) FindValidator(filter interface{}, opts ...*options.FindOneOptions) (Validator, error) {
	var validator Validator
	err := validatorRepository.db.Collection(COLLECTION_NAME).FindOne(context.TODO(), filter, opts...).Decode(&validator)
	if err != nil {
		return validator, err
	}
	return validator, err
}

func (validatorRepository *ValidatorRepositoryImpl) InsertValidator(validator Validator) (*mongo.InsertOneResult, error) {
	return validatorRepository.db.Collection(COLLECTION_NAME).InsertOne(context.TODO(), validator)
}

func NewValidatorRepository(db *mongo.Database) ValidatorRepository {
	rep := new(ValidatorRepositoryImpl)
	rep.db = db
	return rep
}
