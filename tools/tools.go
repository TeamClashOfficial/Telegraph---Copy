package tools

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"telegraph/log"
	"time"
)

type SignAuthorizationHeaderResponse struct {
	Message string `json:"message"`
}

func Check(err error, msg ...string) bool {
	reason := ""
	if len(msg) > 0 {
		reason = msg[0] + ": "
	}
	if err != nil {
		log.Error(reason + err.Error())
		return false
	}
	return true
}

func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func API(method, url string, body []byte) (*http.Response, error) {

	// first we send a random string to be signed and added to the authorization header
	requestUrl := "http://localhost:7045/signer/request-authorization-signature"
	reqResp, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer([]byte(time.Now().String())))
	if err != nil {
		return nil, err
	}
	reqResp.Header.Set("Content-Type", "application/json; charset=utf-8")
	sigReqResp, err := http.DefaultClient.Do(reqResp)
	if err != nil {
		return nil, err
	}
	defer sigReqResp.Body.Close()

	// Read the response body, it should be json.NewEncoder(w).Encode(map[string]string{"signature": signature})
	sigRespBody, err := io.ReadAll(sigReqResp.Body)
	if err != nil {
		return nil, err
	}
	var sigRespMap map[string]string
	err = json.Unmarshal(sigRespBody, &sigRespMap)
	if err != nil {
		return nil, err
	}
	signature := sigRespMap["signature"]

	// now we send the actual request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", signature)
	return http.DefaultClient.Do(req)
}

func ReadFile(file string, v interface{}) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	log.Info("Loading ", file, " from local file")
	val, err := hex.DecodeString(string(data))
	if err != nil {
		return err
	}
	err = json.Unmarshal(val, &v)
	if err != nil {
		return err
	}
	return nil
}

const dataDir = "data/"

func WriteFile(file string, data interface{}) error {
	if !Exists(dataDir) {
		if err := os.MkdirAll(dataDir, 0700); err != nil {
			return err
		}
	}
	dataBytes, err := json.Marshal(data)
	if !Check(err) {
		return err
	}

	err = os.WriteFile(file, []byte(hex.EncodeToString(dataBytes)), 0600)

	Check(err)
	return err
}
