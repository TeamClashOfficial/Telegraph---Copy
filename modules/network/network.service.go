package networkmodule

import (
	"fmt"
	"telegraph/log"
	"telegraph/tools"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NetworkService interface {
	GetNetworks() ([]Network, error)
	AddNetwork(network Network) error
	RemoveNetwork(id string) error
	UpdateNetwork(network Network) error
	GetNetworkByName(name string) (Network, error)
	RemoveAllNetworks() error
}

type NetworkServiceImpl struct {
	networkRepo NetworkRepository
}

func (service *NetworkServiceImpl) GetNetworks() ([]Network, error) {
	fmt.Println("in service")
	networks, err := service.networkRepo.FindNetworks()
	fmt.Println("NETWORKS: 8888888 ", networks)
	// if no networks, return empty array
	if err != nil {
		tools.Check(err)
		return []Network{}, nil
	}
	if len(networks) == 0 {
		return []Network{}, nil
	}
	return networks, nil
}

func (service *NetworkServiceImpl) AddNetwork(network Network) error {
	// create new objectID
	objID := primitive.NewObjectID()
	network.ID = objID
	// check if the network already exists
	filter := bson.D{primitive.E{Key: "contractaddress", Value: network.ContractAddress}}
	opts := options.FindOne().SetSort(bson.D{primitive.E{Key: "contractaddress", Value: 1}})
	_, err := service.networkRepo.FindNetwork(filter, opts)
	if err == nil { // network already exists in the database
		return err
	}

	insertResult, err := service.networkRepo.InsertNetwork(network)
	if err != nil {
		log.Fatal("networkRepository.db.Collection('network').InsertOne error:", err)
		return err
	}
	log.Debug("Inserted a new Network: ", insertResult.InsertedID)
	return nil
}

func (service *NetworkServiceImpl) RemoveNetwork(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	deleteResult, err := service.networkRepo.DeleteNetwork(bson.M{"_id": objID})
	if err != nil {
		log.Fatal("networkRepository.db.Collection('network').DeleteOne error:", err)
	}
	log.Debug("Removed Network: ", deleteResult)
	return err
}

func (service *NetworkServiceImpl) UpdateNetwork(network Network) error {
	result, err := service.networkRepo.UpdateNetwork(bson.M{"_id": network.ID}, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "chainid", Value: network.ChainID},
			{Key: "chaintype", Value: network.ChainType},
			{Key: "name", Value: network.Name},
			{Key: "contractaddress", Value: network.ContractAddress},
			{Key: "evmwssurl", Value: network.EVMWSSURL},
			{Key: "evmhttpurl", Value: network.EVMHTTPURL},
		}},
	})
	if err != nil {
		log.Fatal("networkRepository.db.Collection('network').UpdateOne error:", err)
		return err
	}
	fmt.Printf("Updated %v Network!\n", result)
	return nil
}

func (service *NetworkServiceImpl) GetNetworkByName(name string) (Network, error) {
	log.Debug("GetNetworkByName: ", name)
	var network Network
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	opts := options.FindOne().SetSort(bson.D{primitive.E{Key: "name", Value: 1}})
	network, err := service.networkRepo.FindNetwork(filter, opts)
	if err != nil {
		log.Fatal("networkRepository.db.Collection('network').FindOne error:", err)
	}

	if network == (Network{}) {
		return network, fmt.Errorf("Network not found")
	}
	return network, err
}

func (service *NetworkServiceImpl) RemoveAllNetworks() error {
	return service.networkRepo.DeleteNetworks()
}

func NewNetworkService(networkRepo NetworkRepository) NetworkService {
	service := new(NetworkServiceImpl)
	service.networkRepo = networkRepo
	return service
}
