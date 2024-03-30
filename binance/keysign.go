package binance

import (
	"math/big"
	"sync"
	"telegraph/log"

	"github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	keysignParty tss.Party
)

type keysignInfo struct {
	party   *tss.Party
	partyID *tss.PartyID
	outCh   chan tss.Message
	endCh   chan common.SignatureData
}

func Keysign(msg *big.Int) ([]byte, error) {
	log.Info("Keysign Started")

	outCh := make(chan tss.Message)
	endCh := make(chan common.SignatureData)
	keysignParty = signing.NewLocalParty(msg, params, localPartySaveData, outCh, endCh)
	info := keysignInfo{
		party:   &keysignParty,
		partyID: thisPartyID,
		outCh:   outCh,
		endCh:   endCh,
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go startKeysign(info, wg, msg)
	wg.Wait()
	return []byte("Keysign finished"), nil
}

func startKeysign(info keysignInfo, ksWg *sync.WaitGroup, msg *big.Int) {
	defer ksWg.Done()
	go func() {
		err := (*info.party).Start()
		if err != nil {
			panic(err)
		}
	}()
	processKeysign(info.outCh, info.endCh, msg)
}

func processKeysign(outCh chan tss.Message, endCh chan common.SignatureData, msg *big.Int) {
	for {
		select {
		case msg := <-outCh:
			go doTssJob(msg, "keysign")
		case data := <-endCh:
			pk, err := crypto.SigToPub(crypto.Keccak256Hash(msg.Bytes()).Bytes(), append(data.Signature, 0))
			if err != nil {
				log.Fatal("crypto.Ecrecover:", err)
			}
			log.Info("KeySign completed, Public Key: ", pk)
			CheckSign = true
			return
		}
	}
}
