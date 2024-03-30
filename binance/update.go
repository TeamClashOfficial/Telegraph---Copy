package binance

import (
	"encoding/json"
	"telegraph/log"
	"telegraph/tools"

	"github.com/bnb-chain/tss-lib/v2/tss"
)

type Update struct {
	WireBytes   []byte      `json:"wireBytes"`
	From        tss.PartyID `json:"from"`
	IsBroadcast bool        `json:"isBroadcast"`
	Operation   string      `json:"operation"`
}

func doTssJob(msg tss.Message, op string) {
	wireBytes, r, _ := msg.WireBytes()

	log.Debug("r.IsBroadcast ", r.IsBroadcast)
	if r.IsBroadcast {
		for _, sIP := range URLMap {
			log.Debug("Broadcasting to ip: ", sIP.IP)
			sendUpdate(sIP.IP, wireBytes, msg.GetFrom(), r.IsBroadcast, op)
		}
	} else {
		for _, to := range msg.GetTo() {
			sIP := URLMap[to.GetMoniker()]
			log.Debug("Unicasting to ip: ", sIP.IP)
			sendUpdate(sIP.IP, wireBytes, msg.GetFrom(), r.IsBroadcast, op)
		}
	}

}

func sendUpdate(ip string, wireBytes []byte, from *tss.PartyID, isBroadcast bool, op string) {
	update := Update{
		WireBytes:   wireBytes,
		From:        *from,
		IsBroadcast: isBroadcast,
		Operation:   op,
	}
	updateBytes, err := json.Marshal(update)
	tools.Check(err)
	log.Debug("sending update, ip: ", ip)
	resp, err := tools.API("POST", ip+"/validator/update", updateBytes)
	if tools.Check(err) {
		log.Debug("Status ", resp.StatusCode, " from /validator/update request to ", ip)
	}
	_ = resp.Body.Close()
}

func UpdateFromBytes(updateBytes []byte) {
	var update Update
	if err := json.Unmarshal(updateBytes, &update); err != nil {
		log.Error("failed to unmarshal update", err.Error())
		return
	}

	log.Debug("Recieved Update request, operation: ", update.Operation)
	var err error

	switch update.Operation {
	case "keygen":
		_, err = keygenParty.UpdateFromBytes(update.WireBytes, &update.From, update.IsBroadcast)
	case "keysign":
		_, err = keysignParty.UpdateFromBytes(update.WireBytes, &update.From, update.IsBroadcast)
	case "reshare":
		_, err = reshareParty.UpdateFromBytes(update.WireBytes, &update.From, update.IsBroadcast)
	}

	if err != nil && err.Error() != "Error is nil" {
		log.Error("failed to update", err.Error())
		return
	}
}
