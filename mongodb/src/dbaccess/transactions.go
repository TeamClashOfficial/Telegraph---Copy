package dbaccess

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Transaction struct {
	OrderID         int       `bson:"orderID,omitempty"`
	OrderIDVotes    []int     `bson:"orderIDVotes,omitempty"`
	DetectionTime   time.Time `bson:"detectionTime,omitempty"`
	Hash            string    `bson:"hash,omitempty"`
	EndHash         string    `bson:"endHash,omitempty"`
	BlockNumber     int       `bson:"blockNumber,omitempty"`
	LogIndex        int       `bson:"logIndex,omitempty"`
	Event           string    `bson:"event,omitempty"`
	Sender          string    `bson:"sender,omitempty"`
	Recipient       string    `bson:"recipient,omitempty"`
	Amount          int       `bson:"amount,omitempty"`
	Moniker         string    `bson:"moniker"`
	StartChain      string    `bson:"startChain,omitempty"`
	StartChainIndex string    `bson:"startChainIndex,omitempty"`
	Destination     string    `bson:"destination,omitempty"`
	EndChain        string    `bson:"endChain,omitempty"`
	SignedCount     string    `bson:"signedCount,omitempty"`
	Signers         []string  `bson:"signers,omitempty"`
	Confirmed       bool      `bson:"confirmed,omitempty"`
}

// Util to get Transactions collection
func TransactionCollection() *mongo.Collection {
	return s.Database.Collection("Transactions")
}

func CreateTransactions(transaction *Order) (r *mongo.InsertOneResult, err error) {
	r, err = TransactionCollection().InsertOne(s.Ctx, transaction)
	return
}
