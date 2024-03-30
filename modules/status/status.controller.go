package statusmodule

import (
	"encoding/json"
	"net/http"
)

type StatusController interface {
	Status(w http.ResponseWriter, r *http.Request)
}

type StatusControllerImpl struct {
	service StatusService
}

func (controller *StatusControllerImpl) Status(w http.ResponseWriter, r *http.Request) {
	result, err := controller.service.GetStatus()
	if err != nil {
		http.Error(w, "Error getting status: "+err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error marshaling JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonResp)
}

func NewStatusController(statusService StatusService) StatusController {
	return &StatusControllerImpl{
		service: statusService,
	}
}
