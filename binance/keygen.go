package binance

import (
	"sync"
	"telegraph/log"
	"telegraph/tools"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/tss"
	s256k1 "github.com/btcsuite/btcd/btcec"
)

var (
	thisPartyID        *tss.PartyID
	keygenParty        tss.Party
	SortedPartyIDs     tss.SortedPartyIDs
	URLMap             map[string]SortedIP
	params             *tss.Parameters
	localPartySaveData keygen.LocalPartySaveData
)

func Keygen(threshold int) {
	log.Info("Keygen started")
	tss.SetCurve(s256k1.S256())

	partyCount := len(SortedPartyIDs)
	wg := &sync.WaitGroup{}
	wg.Add(partyCount)
	go GenerateNewKey(thisPartyID, *preParams, wg, threshold)
	wg.Wait()
}

func GenerateNewKey(partyID *tss.PartyID, preParams keygen.LocalPreParams, wg *sync.WaitGroup, threshold int) {
	defer wg.Done()
	ctx := tss.NewPeerContext(SortedPartyIDs)
	params = tss.NewParameters(tss.EC(), ctx, partyID, len(SortedPartyIDs), threshold)
	outCh := make(chan tss.Message)
	endCh := make(chan keygen.LocalPartySaveData)
	keygenParty = keygen.NewLocalParty(params, outCh, endCh, preParams)

	go func() {
		err := keygenParty.Start()
		if err != nil {
			panic(err)
		}
	}()
	processKeyGen(outCh, endCh)
}

func processKeyGen(outCh chan tss.Message, endCh chan keygen.LocalPartySaveData) {
	for {
		select {
		case msg := <-outCh:
			go doTssJob(msg, "keygen")
		case data := <-endCh:
			log.Info("Keygen completed")
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
