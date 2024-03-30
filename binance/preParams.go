package binance

import (
	"encoding/hex"
	"encoding/json"
	"os"
	"telegraph/log"
	"telegraph/tools"
	"time"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
)

var (
	preParams *keygen.LocalPreParams
)

func LoadPreParams() error {
	file, err := os.ReadFile(preParamsFile)
	if err != nil {
		log.Info("Generating Pre Params. This can take a while...")
		preParams, err = keygen.GeneratePreParams(2 * time.Minute)
		if err != nil {
			return err
		}
		if err = tools.WriteFile(preParamsFile, preParams); err != nil {
			return err
		}
	} else {
		var params keygen.LocalPreParams
		val, err := hex.DecodeString(string(file))
		if !tools.Check(err) {
			return err
		}
		err = json.Unmarshal(val, &params)
		if !tools.Check(err) {
			return err
		}
		preParams = &params
	}
	return nil
}
