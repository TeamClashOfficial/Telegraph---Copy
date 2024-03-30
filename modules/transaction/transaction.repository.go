package transactionmodule

import (
	"context"
	"math/big"
	"telegraph/pkg/signer/auth"

	// "telegraph/pkg/wallet"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Transaction struct {
	ID            primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	DetectionTime time.Time          `bson:"detectiontime,omitempty" json:"detectiontime,omitempty"`
	Hash          common.Hash        `bson:"hash,omitempty" json:"hash,omitempty"`
	EndHash       common.Hash        `bson:"endhash,omitempty" json:"endhash,omitempty"`
	BlockNumber   uint64             `bson:"blocknumber,omitempty" json:"blocknumber,omitempty"`
	LogIndex      uint               `bson:"logindex,omitempty" json:"logindex,omitempty"`
	Event         string             `bson:"event,omitempty" json:"event,omitempty"`
	Sender        common.Address     `bson:"sender,omitempty" json:"sender,omitempty"`
	Addresses     []common.Address   `bson:"addresses,omitempty" json:"addresses,omitempty"`
	Uint256       []*big.Int         `bson:"uint256,omitempty" json:"uint256,omitempty"`
	String        []string           `bson:"string,omitempty" json:"string,omitempty"`
	Bool          []bool             `bson:"bool,omitempty" json:"bool,omitempty"`
	FeeAmount     *big.Int           `bson:"feeamount,omitempty" json:"feeamount,omitempty"`
	StartChain    string             `bson:"startchain,omitempty" json:"startchain,omitempty"`
	Destination   common.Address     `bson:"destination,omitempty" json:"destination,omitempty"`
	EndChain      string             `bson:"endchain,omitempty" json:"endchain,omitempty"`
	SignedCount   int                `bson:"signedcount,omitempty" json:"signedcount,omitempty"`
	Signers       [][]byte           `bson:"signers,omitempty" json:"signers,omitempty"` // use crypto.UnmarshalPubkey(bytes) to get public key from bytes
	R             [][32]byte         `bson:"r,omitempty" json:"r,omitempty"`
	S             [][32]byte         `bson:"s,omitempty" json:"s,omitempty"`
	V             []uint8            `bson:"v,omitempty" json:"v,omitempty"`
	H             [][32]byte         `bson:"h,omitempty" json:"h,omitempty"`
	SignerTime    *big.Int           `bson:"signertime,omitempty" json:"signertime,omitempty"`
	Confirmed     bool               `bson:"confirmed,omitempty" json:"confirmed,omitempty"`
}

type TransactionRepository interface {
	InsertTransaction(event Transaction) (*mongo.InsertOneResult, error)
	FindTransactions() ([]Transaction, error)
	UpdateTransaction(filter interface{}, transaction Transaction) (*mongo.UpdateResult, error)
	FindTransaction(filter interface{}, opts ...*options.FindOneOptions) (Transaction, error)
	FindTransactionWithTransform(transform interface{}, filter interface{}, opts ...*options.FindOneOptions) error
}

type TransactionRepositoryImpl struct {
	db *mongo.Database
}

const COLLECTION_NAME = "transaction"

func (transactionRepo *TransactionRepositoryImpl) InsertTransaction(transaction Transaction) (*mongo.InsertOneResult, error) {
	return transactionRepo.db.Collection(COLLECTION_NAME).InsertOne(context.TODO(), transaction)
}

func (transactionRepo *TransactionRepositoryImpl) FindTransactions() ([]Transaction, error) {
	ctx := context.Background()
	var transactions []Transaction
	cur, err := transactionRepo.db.Collection(COLLECTION_NAME).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var transaction Transaction
		err := cur.Decode(&transaction)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (transactionRepo *TransactionRepositoryImpl) FindTransaction(filter interface{}, opts ...*options.FindOneOptions) (Transaction, error) {
	trx := Transaction{}
	err := transactionRepo.db.Collection(COLLECTION_NAME).FindOne(context.TODO(), filter, opts...).Decode(&trx)
	return trx, err
}

func (transactionRepo *TransactionRepositoryImpl) FindTransactionWithTransform(transform interface{}, filter interface{}, opts ...*options.FindOneOptions) error {
	err := transactionRepo.db.Collection(COLLECTION_NAME).FindOne(context.TODO(), filter, opts...).Decode(&transform)
	return err
}

func (transactionRepo *TransactionRepositoryImpl) UpdateTransaction(filter interface{}, inTrx Transaction) (*mongo.UpdateResult, error) {
	return transactionRepo.db.Collection(COLLECTION_NAME).ReplaceOne(context.TODO(), filter, inTrx)
}

func NewTransactionRepository(db *mongo.Database, auth auth.IAuth) TransactionRepository {
	rep := new(TransactionRepositoryImpl)
	rep.db = db
	return rep
}
