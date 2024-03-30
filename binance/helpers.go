package binance

import (
	"encoding/json"
	"math/big"
	"net/http"
	"path/filepath"
	"telegraph/config"
	"telegraph/log"
	"telegraph/tools"
	"time"

	"github.com/bnb-chain/tss-lib/v2/tss"
)

const dataDir = "data"

var (
	keyFile       = filepath.Join(dataDir, "key.data")
	paramsFile    = filepath.Join(dataDir, "params.data")
	urlMapFile    = filepath.Join(dataDir, "url.data")
	preParamsFile = filepath.Join(dataDir, "preParams.data")
)

var (
	PartyIDs  []*tss.PartyID
	CheckSign bool
)

type SortedIP struct {
	IP    string
	Index int
}
type Parameters struct {
	PartyID             *tss.PartyID
	Parties             tss.SortedPartyIDs
	PartyCount          int
	Threshold           int
	SafePrimeGenTimeout time.Duration
}

func ExportParams(params *tss.Parameters) *Parameters {
	return &Parameters{
		PartyID:             params.PartyID(),
		Parties:             params.Parties().IDs(),
		PartyCount:          params.PartyCount(),
		Threshold:           params.Threshold(),
		SafePrimeGenTimeout: params.SafePrimeGenTimeout(),
	}
}

func ImportParams(p *Parameters) *tss.Parameters {
	ctx := tss.NewPeerContext(p.Parties)
	return tss.NewParameters(tss.EC(), ctx, p.PartyID, p.PartyCount, p.Threshold)
}

func Init() error {
	conf := config.Conf
	err := tools.ReadFile(keyFile, &localPartySaveData)
	if err != nil {
		if conf.IsGenesis {
			thisPartyID = GetPartyID(conf, 0)
			PartyIDs = append(PartyIDs, thisPartyID)
		} else { // keyFile does not exist, send node info to genesis server to start keygen
			log.Info("Sending node info to genesis server")
			resp, err := tools.API("POST", conf.GenesisIP+"/validator/params", conf.ToBytes())
			if !tools.Check(err) || resp.StatusCode != http.StatusOK {
				log.Debug("Failed to send node info to genesis server, Error: ", err)
			}
			_ = resp.Body.Close()
		}
		return err
	}
	pparams := new(Parameters)
	err = tools.ReadFile(paramsFile, pparams)
	if tools.Check(err) {
		params = ImportParams(pparams)
		thisPartyID = params.PartyID()
		PartyIDs = params.Parties().IDs()
		SortedPartyIDs = PartyIDs
	} else {
		return err
	}
	err = tools.ReadFile(urlMapFile, &URLMap)
	if !tools.Check(err) {
		return err
	}
	return nil
}

func GetPartyID(conf config.Config, index ...int) *tss.PartyID {
	bigIntKey := new(big.Int)
	bigIntKey, ok := bigIntKey.SetString(conf.Key, 10)
	if !ok {
		log.Error("SetString: error")
		return nil
	}
	partyID := tss.NewPartyID(conf.ID, conf.Moniker, bigIntKey)
	if len(index) == 1 {
		partyID.Index = index[0]
	}
	return partyID
}

func AppendPartyID(partyID *tss.PartyID) []byte {
	PartyIDs = append(PartyIDs, partyID)
	SortedPartyIDs = tss.SortPartyIDs(PartyIDs)
	log.Debug("SortedPartyIDs length: ", len(SortedPartyIDs))
	partyIDsBytes, err := json.Marshal(SortedPartyIDs)
	tools.Check(err)

	return partyIDsBytes
}

func ReplacePartyIDs(index int, partyIDBytes []byte) {
	err := json.Unmarshal(partyIDBytes, &SortedPartyIDs)
	tools.Check(err)
	PartyIDs = SortedPartyIDs
	thisPartyID = SortedPartyIDs[index]
}
