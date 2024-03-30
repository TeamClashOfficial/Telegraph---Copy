package api

import (
	"encoding/json"
	"sync"
	"telegraph/binance"
	"telegraph/config"
	"telegraph/log"
	"telegraph/tools"
	"time"
)

type infoStruct struct {
	URLMap        map[string]binance.SortedIP `json:"urlMap"`
	PartyIDsBytes []byte                      `json:"partyIDsBytes"`
	Index         int                         `json:"index"`
}

func SyncToMap(syncMap *sync.Map) map[string]binance.SortedIP {
	sMap := make(map[string]binance.SortedIP)
	syncMap.Range(func(moniker, sIP interface{}) bool {
		sMap[moniker.(string)] = sIP.(binance.SortedIP)
		return true
	})
	return sMap
}

func MapToSync(m map[string]binance.SortedIP) *sync.Map {
	sMap := new(sync.Map)
	for k, v := range m {
		sMap.Store(k, v)
	}
	return sMap
}

// func Params(w http.ResponseWriter, r *http.Request) {
//  body, err := io.ReadAll(io.LimitReader(r.Body, 1048576))
// 	tools.Check(err)

// 	if r.Method == "POST" {
// 		if !config.Conf.IsGenesis {
// 			w.Write([]byte("POST method only allowed on Genesis nodes"))
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}

// 		var clientConf config.Config
// 		if !clientConf.LoadFromJson(body, true) {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		if clientConf.PartyPassword != config.Conf.PartyPassword {
// 			w.WriteHeader(http.StatusForbidden)
// 			return
// 		}
// 		partyIDsBytes := binance.AppendPartyID(binance.GetPartyID(clientConf))
// 		partyCount := len(binance.SortedPartyIDs)
// 		log.Debug("Partycount: ", partyCount, " clientConf.IP: ", clientConf.IP)

// 		validator.AddValidator(common.Address{}, clientConf.Moniker, clientConf.IP, false)
// 		SyncUrlMap.Store(clientConf.Moniker, binance.SortedIP{IP: clientConf.IP, Index: partyCount - 1})
// 		w.WriteHeader(http.StatusOK)
// 		go updateParams(partyIDsBytes, partyCount)
// 	} else if r.Method == "PUT" {
// 		log.Debug("Recieved PUT request from: ", r.RemoteAddr)
// 		if config.Conf.IsGenesis {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		var info infoStruct
// 		err := json.Unmarshal(body, &info)
// 		tools.Check(err)

// 		binance.ReplacePartyIDs(info.Index, info.PartyIDsBytes)

// 		for moniker, sIp := range info.URLMap {
// 			validator.AddValidator(common.Address{}, moniker, sIp.IP)
// 		}
// 		binance.URLMap = info.URLMap
// 		w.WriteHeader(http.StatusOK)
// 	}
// }

func updateParams(partyIDsBytes []byte, partyCount int) {
	if partyCount != config.Conf.Threshold+1 {
		return
	}
	time.Sleep(time.Second)
	URLMap := SyncToMap(SyncUrlMap)
	binance.URLMap = URLMap
	log.Debug("URL Map: ", URLMap)

	dataSent := true
	SyncUrlMap.Range(func(_, sIp interface{}) bool {
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
		// index++ // TODO: ineffectual assignment, function quits before index is reused
		return true
	})

	if dataSent { // start keygen or reshare workflow
		var action string
		if ReshareFlag {
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
