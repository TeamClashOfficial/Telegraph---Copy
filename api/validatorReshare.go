package api

import (
	"net/http"
	"telegraph/binance"
	"telegraph/config"
	"telegraph/log"
	"telegraph/tools"
)

func Reshare(w http.ResponseWriter, r *http.Request) {
	body, err := validateRequest(w, r)
	if err != nil {
		return
	}
	if config.Conf.IsGenesis {
		for moniker, sIP := range binance.URLMap {
			if moniker != config.Conf.Moniker {
				resp, err := tools.API("POST", sIP.IP+"/validator/reshare", body)
				if tools.Check(err) {
					log.Debug("Status ", resp.StatusCode, " from /validator/reshare request to ", sIP.IP)
				}
				_ = resp.Body.Close()
			}
		}
	}
	go binance.Reshare(config.Conf)
	_, _ = w.Write([]byte("Keygen started"))
}
