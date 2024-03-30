package validatormodule

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"telegraph/binance"
	"telegraph/config"
	"telegraph/log"
	"telegraph/tools"
)

type ValidatorController interface {
	GetMyValidator(w http.ResponseWriter, r *http.Request)
	AddValidatorParams(w http.ResponseWriter, r *http.Request)
	UpdateValidatorParams(w http.ResponseWriter, r *http.Request)
	StartValidator(w http.ResponseWriter, r *http.Request)
	UpdateValidator(w http.ResponseWriter, r *http.Request)
	SignValidator(w http.ResponseWriter, r *http.Request)
	ReshareValidator(w http.ResponseWriter, r *http.Request)
}

type ValidatorControllerImpl struct {
	validatorService ValidatorService
}

func (controller *ValidatorControllerImpl) GetMyValidator(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetMyValidator")
	validator, err := controller.validatorService.GetMyValidator()
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// w.Write([]byte(err.Error()))
		_ = json.NewEncoder(w).Encode(nil)
		return
	}

	_ = json.NewEncoder(w).Encode(validator)

}

func (controller *ValidatorControllerImpl) AddValidatorParams(w http.ResponseWriter, r *http.Request) {
	var newValidator Validator

	if err := json.NewDecoder(r.Body).Decode(&newValidator); err != nil {
		tools.Check(err)
		_, _ = w.Write([]byte("Unable to read request body"))
		return
	}
	// log newValidator
	log.Debug("New Validator: ", newValidator)

	// if !config.Conf.IsGenesis {
	// 	w.Write([]byte("POST method only allowed on Genesis nodes"))
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	var clientConf config.Config

	// if !clientConf.LoadFromJson(body, true) {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	// if clientConf.PartyPassword != config.Conf.PartyPassword {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	return
	// }

	// partyIDsBytes := binance.AppendPartyID(binance.GetPartyID(clientConf))
	// partyCount := len(binance.SortedPartyIDs)

	err := controller.validatorService.AddParams(clientConf, newValidator)

	// api.SyncUrlMap.Store(clientConf.Moniker, binance.SortedIP{IP: clientConf.IP, Index: partyCount - 1})
	// go updateParams(partyIDsBytes, partyCount)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Validator Params Added"))

}

func (controller *ValidatorControllerImpl) UpdateValidatorParams(w http.ResponseWriter, r *http.Request) {
	var info infoStruct

	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if config.Conf.IsGenesis {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	binance.ReplacePartyIDs(info.Index, info.PartyIDsBytes)

	err := controller.validatorService.ReplaceAndUpdateParams(info)

	binance.URLMap = info.URLMap
	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Validator Params Updated"))
}

func (controller *ValidatorControllerImpl) StartValidator(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	var clientConf config.Config

	if !clientConf.LoadFromJSON(body, false) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if clientConf.PartyPassword != config.Conf.PartyPassword {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if len(binance.SortedPartyIDs) <= config.Conf.Threshold {
		msg := "Threshold is less than the length of sorted party IDs"
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(msg))
		return
	}

	if config.Conf.IsGenesis {
		err = controller.validatorService.StartValidator(body, binance.URLMap)
		if err != nil {
			log.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	// go binance.Keygen(config.Conf.Threshold)
	_, _ = w.Write([]byte("Validator Started"))

}

func (controller *ValidatorControllerImpl) SignValidator(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	messageStruct := struct {
		Message string `json:"message"`
	}{}

	err = json.Unmarshal(body, &messageStruct)
	if err != nil || messageStruct.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intMessage := new(big.Int)
	intMessage, ok := intMessage.SetString(messageStruct.Message, 10)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("SetString: error"))
		return
	}

	if config.Conf.IsGenesis {
		err = controller.validatorService.Sign(body, binance.URLMap)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	go func() { _, _ = binance.Keysign(intMessage) }()
	_, _ = w.Write([]byte("Keysign started"))
}

func (controller *ValidatorControllerImpl) ReshareValidator(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var clientConf config.Config
	if !clientConf.LoadFromJSON(body, false) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if clientConf.PartyPassword != config.Conf.PartyPassword {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if len(binance.SortedPartyIDs) <= config.Conf.Threshold {
		msg := "Threshold is less than the length of sorted party IDs"
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(msg))
		return
	}

	if config.Conf.IsGenesis {
		err = controller.validatorService.Reshare(body, binance.URLMap)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	go binance.Reshare(config.Conf)
	_, _ = w.Write([]byte("Keygen started"))

}

func (controller *ValidatorControllerImpl) UpdateValidator(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	go binance.UpdateFromBytes(bytes)
	_, _ = w.Write([]byte("Update started"))
}

func NewValidatorController(validatorService ValidatorService) ValidatorController {
	return &ValidatorControllerImpl{
		validatorService: validatorService,
	}
}
