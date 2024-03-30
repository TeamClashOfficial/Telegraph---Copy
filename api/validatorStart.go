package api

import (
	"io"
	"net/http"
	"telegraph/binance"
	"telegraph/config"
	"telegraph/log"
	"telegraph/tools"
)

func validateRequest(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(io.LimitReader(r.Body, 1048576))
	tools.Check(err)
	var clientConf config.Config
	if !clientConf.LoadFromJSON(body, false) {
		w.WriteHeader(http.StatusUnauthorized)
		return nil, err
	}
	if clientConf.PartyPassword != config.Conf.PartyPassword {
		w.WriteHeader(http.StatusForbidden)
		return nil, err
	}
	if len(binance.SortedPartyIDs) <= config.Conf.Threshold {
		msg := "Threshold is less than the length of sorted party IDs"
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(msg))
		return nil, err
	}
	return body, nil
}
func Start(w http.ResponseWriter, r *http.Request) {
	body, err := validateRequest(w, r)
	if err != nil {
		return
	}
	if config.Conf.IsGenesis {
		for moniker, sIp := range binance.URLMap {
			if moniker != config.Conf.Moniker {
				resp, err := tools.API("POST", sIp.IP+"/validator/start", body)
				if tools.Check(err) {
					log.Debug("Status ", resp.StatusCode, " from /validator/start request to ", sIp.IP)
				}
				_ = resp.Body.Close()
			}
		}
	}
	go binance.Keygen(config.Conf.Threshold)
	_, _ = w.Write([]byte("Keygen started"))
}
