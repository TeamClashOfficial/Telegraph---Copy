package dbaccess

import (
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Validator struct {
	NodeOrder    int    `bson:"nodeOrder,omitempty"`
	Moniker      string `bson:"moniker"`
	PubKey       string `bson:"pubKey,omitempty"`
	LastPingTime int    `bson:"lastPingTime,omitempty"`
	IP           string `bson:"IP"`
}

// Util to get Transactions collection
func ValidatorCollection() *mongo.Collection {
	return s.Database.Collection("Validators")
}

func initValidatorIndexes() {
	validatorIndexes := []mongo.IndexModel{
		{
			Keys:    bson.M{"moniker": -1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.M{"IP": -1},
			Options: options.Index().SetUnique(true),
		},
	}
	_, err := ValidatorCollection().Indexes().CreateMany(s.Ctx, validatorIndexes)
	if err != nil {
		log.Fatal("Error occured while creating validator schema indexes", err)
		return
	}
}

func CreateValidator(validator *Validator) (r *mongo.InsertOneResult, err error) {
	r, err = ValidatorCollection().InsertOne(s.Ctx, validator)
	return
}

func FindValidators(filter *Validator) (valids []Validator) {
	cursor, err := ValidatorCollection().Find(s.Ctx, filter)
	if err != nil {
		log.Error("Error occured while creating the cursor", err)
		return
	}
	if err = cursor.All(s.Ctx, &valids); err != nil {
		log.Error("Error occured while un-marshal the cursor", err)
		return
	}
	return
}
