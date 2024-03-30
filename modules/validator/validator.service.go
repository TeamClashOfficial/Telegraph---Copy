package validatormodule

import (
	"encoding/json"
	"fmt"
	"telegraph/api"
	"telegraph/binance"
	"telegraph/config"
	"telegraph/log"
	"telegraph/tools"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type infoStruct struct {
	URLMap        map[string]binance.SortedIP `json:"urlMap"`
	PartyIDsBytes []byte                      `json:"partyIDsBytes"`
	Index         int                         `json:"index"`
}

type ValidatorService interface {
	GetValidators() ([]Validator, error)
	GetMyValidator() (Validator, error)
	FindValidatorByAddress(address common.Address) (Validator, error)
	AddValidator(newValidator Validator) error
	AddParams(conf config.Config, newValidator Validator) error
	ReplaceAndUpdateParams(info infoStruct) error
	StartValidator(body []byte, urlMap map[string]binance.SortedIP) error
	Sign(body []byte, urlMap map[string]binance.SortedIP) error
	Reshare(body []byte, urlMap map[string]binance.SortedIP) error
}

type ValidatorServiceImpl struct {
	validatorRepository ValidatorRepository
}

func (validatorService *ValidatorServiceImpl) GetValidators() ([]Validator, error) {
	return validatorService.validatorRepository.FindValidators()
}

func (validatorService *ValidatorServiceImpl) GetMyValidator() (Validator, error) {
	var validator Validator
	filter := bson.D{primitive.E{Key: "ismine", Value: true}}

	fmt.Println("filter", filter)
	opts := options.FindOne().SetSort(bson.D{primitive.E{Key: "ismine", Value: 1}})
	validator, err := validatorService.validatorRepository.FindValidator(filter, opts)
	if err != nil {
		return validator, err
	}
	return validator, err
}

func (validatorService *ValidatorServiceImpl) FindValidatorByAddress(address common.Address) (Validator, error) {
	var validator Validator
	filter := bson.D{primitive.E{Key: "publicethaddress", Value: address}}
	opts := options.FindOne().SetSort(bson.D{primitive.E{Key: "publicethaddress", Value: 1}})
	validator, err := validatorService.validatorRepository.FindValidator(filter, opts)
	if err != nil {
		return validator, err
	}
	return validator, err
}

func (validatorService *ValidatorServiceImpl) AddValidator(newValidator Validator) error {
	validator := Validator{
		PublicETHAddress: newValidator.PublicETHAddress,
		Moniker:          newValidator.Moniker,
		LastPingTime:     time.Now(),
		Domain:           newValidator.Domain,
		IsMine:           newValidator.IsMine,
	}

	filter := bson.D{primitive.E{Key: "domain", Value: validator.Domain}}
	opts := options.FindOne().SetSort(bson.D{primitive.E{Key: "domain", Value: 1}})

	_, err := validatorService.validatorRepository.FindValidator(filter, opts)
	if err == nil { // validator already exists in the database
		return err
	}
	_, err = validatorService.validatorRepository.InsertValidator(validator)
	return err
}

func (validatorService *ValidatorServiceImpl) AddParams(clientConf config.Config, newValidator Validator) error {
	err := validatorService.AddValidator(newValidator)
	if err != nil {
		return err
	}
	return nil
}

func (validatorService *ValidatorServiceImpl) ReplaceAndUpdateParams(info infoStruct) error {

	// for moniker, sIp := range info.URLMap {
	// 	err := validatorService.AddValidator(common.Address{}, moniker, sIp.IP)
	// 	if err != nil {
	// 		log.Error(err.Error())
	// 		return err
	// 	}
	// }

	return nil
}

func (validatorService *ValidatorServiceImpl) StartValidator(body []byte, urlMap map[string]binance.SortedIP) error {
	for moniker, sIp := range urlMap {
		if moniker != config.Conf.Moniker {
			resp, err := tools.API("POST", sIp.IP+"/validator/start", body)
			if err != nil {
				return nil
			}
			_ = resp.Body.Close()
			log.Info("Status ", resp.StatusCode, " from /validator/start request to ", sIp.IP)
		}
	}
	return nil
}

func (validatorService *ValidatorServiceImpl) Sign(body []byte, urlMap map[string]binance.SortedIP) error {

	for moniker, sIp := range urlMap {
		if moniker != config.Conf.Moniker {
			resp, err := tools.API("POST", sIp.IP+"/validator/sign", body)

			if err != nil {
				return err
			}
			log.Info("Status ", resp.StatusCode, " from /validator/sign request to ", sIp.IP)
			_ = resp.Body.Close()
		}
	}
	return nil
}

func (validatorService *ValidatorServiceImpl) Reshare(body []byte, urlMap map[string]binance.SortedIP) error {
	for moniker, sIp := range urlMap {
		if moniker != config.Conf.Moniker {
			resp, err := tools.API("POST", sIp.IP+"/validator/reshare", body)
			if err != nil {
				return err
			}
			log.Info("Status ", resp.StatusCode, " from /validator/reshare request to ", sIp.IP)
			_ = resp.Body.Close()
		}
	}
	return nil
}

func updateParams(partyIDsBytes []byte, partyCount int) { /* TODO: Need to refactor this function */
	if partyCount != config.Conf.Threshold+1 {
		return
	}
	time.Sleep(time.Second)
	URLMap := api.SyncToMap(api.SyncUrlMap)
	binance.URLMap = URLMap
	log.Debug("URL Map: ", URLMap)

	dataSent := true
	api.SyncUrlMap.Range(func(_, sIp interface{}) bool {
		ip := sIp.(binance.SortedIP).IP
		index := sIp.(binance.SortedIP).Index
		if ip == config.Conf.IP {
			return true
		}
		info := infoStruct{
			PartyIDsBytes: partyIDsBytes,
			Index:         index,
			URLMap:        URLMap,
		}
		log.Debug(info.Index)
		infoBytes, err := json.Marshal(info)
		tools.Check(err)
		log.Debug("Sending PUT request to: ", ip+"/validator/params")
		resp, err := tools.API("PUT", ip+"/validator/params", infoBytes)
		if tools.Check(err) {
			log.Debug("Sent PUT request to: ", ip+"/validator/params", ", Response Status Code: ", resp.StatusCode)
		} else {
			dataSent = false
		}
		_ = resp.Body.Close()
		// index++ // TODO: innefectual assignment, functions quits before index is reused
		return true
	})
	if dataSent { // start keygen or reshare workflow
		var action string
		if api.ReshareFlag {
			action = "reshare"
		} else {
			action = "start"
		}
		resp, err := tools.API("POST", config.Conf.IP+"/validator/"+action, config.Conf.ToBytes())
		if tools.Check(err) {
			log.Debug("Sent "+action+" request to self, Response Status Code: ", resp.StatusCode)
		}
		_ = resp.Body.Close()
	}
}

func NewValidatorService(validatorRepository ValidatorRepository) ValidatorService {
	return &ValidatorServiceImpl{
		validatorRepository: validatorRepository,
	}
}
