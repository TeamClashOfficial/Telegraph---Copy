package api

import (
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"telegraph/binance"
	"telegraph/config"
	"telegraph/log"
	"telegraph/tools"
)

func Sign(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(io.LimitReader(r.Body, 1048576))
	tools.Check(err)
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
		for moniker, sIP := range binance.URLMap {
			if moniker != config.Conf.Moniker {
				resp, err := tools.API("POST", sIP.IP+"/validator/sign", body)
				if tools.Check(err) {
					log.Debug("Status ", resp.StatusCode, " from /validator/sign request to ", sIP.IP)
				}
				_ = resp.Body.Close()
			}
		}
	}
	go func() { _, _ = binance.Keysign(intMessage) }()
	_, _ = w.Write([]byte("Keysign started"))
}
