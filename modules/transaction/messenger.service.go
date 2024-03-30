package transactionmodule

import (
	"encoding/json"
	"io"
	"net/http"
	"telegraph/log"
	validatormodule "telegraph/modules/validator"
	"telegraph/tools"
)

func SendTransaction(t Transaction, v validatormodule.Validator) error {
	body, err := json.Marshal(t)
	tools.Check(err)

	log.Debug("Sending transaction to node: ", v.Domain+"/transaction/create")
	resp, err := tools.API("POST", v.Domain+"/transaction/create", body)
	if tools.Check(err) {
		log.Debug("Status ", resp.StatusCode, " from /transaction/create request to ", v.Domain)
	} else {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		_, err := io.Copy(io.Discard, resp.Body)
		if err != nil {
			return err
		}

	} else {
		// The status is not Created. print the error.
		log.Error("Failed to Send Transaction: ", resp.Status)
	}
	return nil
}
