package networkmodule

import (
	"context"
	"fmt"
	"telegraph/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Network struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	ChainID         int                `json:"chainid" bson:"chainid"`
	ChainType       string             `json:"chaintype" bson:"chaintype"`
	Name            string             `json:"name" bson:"name"`
	ContractAddress string             `json:"contractaddress" bson:"contractaddress"`
	EVMWSSURL       string             `json:"evmwssurl" bson:"evmwssurl"`
	EVMHTTPURL      string             `json:"evmhttpurl" bson:"evmhttpurl"`
}

type NetworkRepository interface {
	InsertNetwork(network Network) (*mongo.InsertOneResult, error)
	FindNetworks() ([]Network, error)
	UpdateNetwork(filter interface{}, network interface{}) (*mongo.UpdateResult, error)
	FindNetwork(filter interface{}, opts ...*options.FindOneOptions) (Network, error)
	DeleteNetwork(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	DeleteNetworks() error
}

type NetworkRepositoryImpl struct {
	db *mongo.Database
}

const COLLECTION_NAME = "network"

func (networkRepository *NetworkRepositoryImpl) FindNetworks() ([]Network, error) {
	results, err := networkRepository.db.Collection(COLLECTION_NAME).Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	count, err := networkRepository.db.Collection(COLLECTION_NAME).CountDocuments(context.Background(), bson.D{})
	if count == 0 {
		fmt.Printf("No chains found\n")
		return nil, err
	}

	if err != nil {
		panic(err)
	}

	var decoded []Network
	for results.Next(context.TODO()) {
		var elem Network
		err := results.Decode(&elem)
		if err != nil {
			log.Fatal("results.Decode error:", err)
		}
		decoded = append(decoded, elem)
	}
	return decoded, nil
}

func (networkRepository *NetworkRepositoryImpl) FindNetwork(filter interface{}, opts ...*options.FindOneOptions) (Network, error) {
	network := Network{}
	err := networkRepository.db.Collection(COLLECTION_NAME).FindOne(context.TODO(), filter, opts...).Decode(&network)
	return network, err
}

func (networkRepository *NetworkRepositoryImpl) InsertNetwork(network Network) (*mongo.InsertOneResult, error) {
	return networkRepository.db.Collection(COLLECTION_NAME).InsertOne(context.TODO(), network)
}

func (networkRepository *NetworkRepositoryImpl) UpdateNetwork(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return networkRepository.db.Collection(COLLECTION_NAME).UpdateOne(context.TODO(), filter, update)
}

func (networkRepository *NetworkRepositoryImpl) DeleteNetwork(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return networkRepository.db.Collection(COLLECTION_NAME).DeleteOne(context.TODO(), filter, opts...)
}

func (networkRepository *NetworkRepositoryImpl) EditNetwork(network Network) error {
	fmt.Println("NETWORK ID: ", network.ID)
	result, err := networkRepository.db.Collection(COLLECTION_NAME).UpdateOne(
		context.TODO(),
		bson.M{"_id": network.ID},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "chainid", Value: network.ChainID},
				{Key: "chaintype", Value: network.ChainType},
				{Key: "name", Value: network.Name},
				{Key: "contractaddress", Value: network.ContractAddress},
				{Key: "evmwssurl", Value: network.EVMWSSURL},
				{Key: "evmhttpurl", Value: network.EVMHTTPURL},
			}},
		},
	)
	if err != nil {
		log.Fatal("networkRepository.db.Collection('network').UpdateOne error:", err)
		return err
	}
	fmt.Printf("Updated %v Network!\n", result)
	return nil
}

func (networkRepository *NetworkRepositoryImpl) DeleteNetworks() error {
	return networkRepository.db.Collection(COLLECTION_NAME).Drop(context.TODO())
}

func NewNetworkRepository(db *mongo.Database) NetworkRepository {
	rep := new(NetworkRepositoryImpl)
	rep.db = db
	return rep
}
