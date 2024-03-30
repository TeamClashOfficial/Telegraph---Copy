package networkmodule

import (
	"encoding/json"
	"fmt"
	"net/http"
	"telegraph/tools"
)

type NetworkController interface {
	GetNetworks(w http.ResponseWriter, r *http.Request)
	AddNetwork(w http.ResponseWriter, r *http.Request)
	RemoveNetwork(w http.ResponseWriter, r *http.Request)
	UpdateNetwork(w http.ResponseWriter, r *http.Request)
}

type NetworkControllerImpl struct {
	networkService NetworkService
}

func (controller *NetworkControllerImpl) GetNetworks(w http.ResponseWriter, _ *http.Request) {
	fmt.Print("inside controller")
	networks, err := controller.networkService.GetNetworks()
	print("NETWORKS: ", networks)
	// return empty array if no networks
	if err != nil {
		tools.Check(err)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_ = json.NewEncoder(w).Encode(networks)
}

func (controller *NetworkControllerImpl) AddNetwork(w http.ResponseWriter, r *http.Request) {
	var newNetwork Network

	if err := json.NewDecoder(r.Body).Decode(&newNetwork); err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	if err := controller.networkService.AddNetwork(newNetwork); err != nil {
		tools.Check(err)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = w.Write([]byte("New Network Added"))

}

func (controller *NetworkControllerImpl) RemoveNetwork(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	id := v.Get("id")

	if err := controller.networkService.RemoveNetwork(id); err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	_, _ = w.Write([]byte("Network Removed"))
}

func (controller *NetworkControllerImpl) UpdateNetwork(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NETWORKS BODY: ", r.Body)

	var newNetwork Network

	if err := json.NewDecoder(r.Body).Decode(&newNetwork); err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	if err := controller.networkService.UpdateNetwork(newNetwork); err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	_, _ = w.Write([]byte("Network Updated"))
}

func NewNetworkController(service NetworkService) NetworkController {
	controller := new(NetworkControllerImpl)
	controller.networkService = service
	return controller
}
