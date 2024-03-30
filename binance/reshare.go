package binance

import (
	"telegraph/config"
	"telegraph/log"
	"telegraph/tools"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/resharing"
	"github.com/bnb-chain/tss-lib/v2/tss"
)

var (
	reshareParty tss.Party
)

func Reshare(conf config.Config) {
	log.Info("Resharing started")
	outCh := make(chan tss.Message)
	endCh := make(chan keygen.LocalPartySaveData)
	var ctx *tss.PeerContext
	if params != nil {
		ctx = tss.NewPeerContext(params.Parties().IDs())
	} else {
		i := thisPartyID.Index
		parties := SortedPartyIDs[:i]
		parties = append(parties, SortedPartyIDs[i+1:]...) // removing self party from SortedPartyIDs
		ctx = tss.NewPeerContext(parties)
		params = tss.NewParameters(tss.EC(), ctx, thisPartyID, len(SortedPartyIDs), conf.Threshold)
	}
	log.Debug("ECDSAPub: ", localPartySaveData.ECDSAPub)
	newCtx := tss.NewPeerContext(SortedPartyIDs)
	resharingParams := tss.NewReSharingParameters(tss.EC(), ctx, newCtx, thisPartyID,
		params.PartyCount(), params.Threshold(), len(SortedPartyIDs), conf.Threshold)

	params = tss.NewParameters(tss.EC(), newCtx, thisPartyID, len(SortedPartyIDs), conf.Threshold)

	reshareParty = resharing.NewLocalParty(resharingParams, localPartySaveData, outCh, endCh)
	go func() {
		err := reshareParty.Start()
		if err != nil {
			panic(err)
		}
	}()
	processResharing(outCh, endCh)
}

func processResharing(outCh chan tss.Message, endCh chan keygen.LocalPartySaveData) {
	for {
		select {
		case msg := <-outCh:
			go doTssJob(msg, "reshare")
		case data := <-endCh:
			log.Info("Reshare completed")
			localPartySaveData = data
			log.Debug("Writing key file")
			err := tools.WriteFile(keyFile, localPartySaveData)
			if tools.Check(err) {
				log.Debug("Key file written successfully")
			}
			err = tools.WriteFile(paramsFile, ExportParams(params))
			if tools.Check(err) {
				log.Debug("Params file written successfully")
			}
			err = tools.WriteFile(urlMapFile, URLMap)
			if tools.Check(err) {
				log.Debug("urlMap file written successfully")
			}
			return
		}
	}
}
