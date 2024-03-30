package config

import (
	"encoding/json"
	"os"
	"strconv"
	"telegraph/log"

	env "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

var Conf Config

func init() {
	NewConf()
	log.Init(Conf.LogLevel)
}

func NewConf() {
	Conf = Config{}

	var err error

	if test := os.Getenv("TEST"); test != "" {
		err = godotenv.Load("/tmp/.telegraph.env")
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		log.Fatal(err)
	}

	err = env.Parse(&Conf)

	os.Setenv("SOFTHSM2_CONF", Conf.HSM.ConfigPath)

	if err != nil {
		panic(err)
	}
}

type Config struct {
	MainPortAddress string `env:"MAIN_PORT_ADDRESS" json:"main_port_address"`
	FinishedSetup   bool   `env:"FINISHED_SETUP" json:"finished_setup"`
	PublicKey       string `env:"PUBLIC_KEY" json:"public_key"`
	IsGenesis       bool   `env:"IS_GENESIS" json:"is_genesis"`
	ID              string `env:"ID,required" json:"id"`
	Moniker         string `env:"MONIKER,required" json:"moniker"`
	Key             string `env:"KEY,required" json:"key"`
	GenesisIP       string `env:"GENESIS_IP" json:"genesis_ip"`
	HokkETH         string `env:"HOKK_ETH" json:"hokk_eth"`
	IP              string `env:"IP,required" json:"ip"`
	PORT            string `env:"PORT,required" json:"port"`
	PartyPassword   string `env:"PARTY_PASSWORD,required" json:"party_password"`
	Threshold       int    `env:"THRESHOLD" json:"threshold"`
	LogLevel        uint32 `env:"LOG_LEVEL" json:"log_level"`
	DbURL           string `env:"DB_URL" json:"db_url"`
	TLSCert         string `env:"TLS_CERT" json:"tls_cert"`
	TLSKey          string `env:"TLS_KEY" json:"tls_key"`
	MaxConn         int    `env:"MAX_CONN" json:"max_conn"`
	HSM             struct {
		ConfigPath      string `env:"HSM_CONFIG_PATH" json:"config_path"`
		Path            string `env:"HSM_PATH" json:"path"`
		TokenLabel      string `env:"HSM_TOKEN_LABEL" json:"token_label"`
		Pin             string `env:"HSM_PIN" json:"pin"`
		UseGCMIVFromHSM bool   `env:"HSM_USE_GCM_IV" json:"use_gcm_iv"`
		CKAID           string `env:"HSM_CKA_ID" json:"cka_id"`
		CKALabel        string `env:"HSM_CKA_LABEL" json:"cka_label"`
	} `json:"hsm"`
	CoreContract struct {
		ContractAddress string `env:"CONTRACT_ADDRESS" json:"contract_address"`
		ChainID         int    `env:"CHAIN_ID" json:"chain_id"`
		ChainType       string `env:"CHAIN_TYPE" json:"chain_type"`
		Name            string `env:"NAME" json:"name"`
		EvmWSSURL       string `env:"EVM_WSS_URL" json:"evm_wssurl"`
		EvmHTTPURL      string `env:"EVM_HTTP_URL" json:"evm_httpurl"`
	} `json:"core_contract"`
}

func (config *Config) Validate() bool {
	if config.ID == "" || config.Moniker == "" ||
		config.Key == "" || config.PartyPassword == "" || config.IP == "" {
		return false
	}
	return true
}

func (config *Config) Load(configBytes []byte) error {
	return json.Unmarshal(configBytes, config)
}

func (config *Config) LoadFromJSON(body []byte, validate bool) bool {
	_ = json.Unmarshal(body, config)
	if validate {
		return config.Validate()
	}
	return true
}

func (config *Config) ToBytes() []byte {
	configBytes, _ := json.Marshal(config)
	return configBytes
}

func (config *Config) Save() error {
	var m = make(map[string]string)

	m["MAIN_PORT_ADDRESS"] = Conf.MainPortAddress
	m["FINISHED_SETUP"] = strconv.FormatBool(Conf.FinishedSetup)
	m["PUBLIC_KEY"] = Conf.PublicKey
	m["IS_GENESIS"] = strconv.FormatBool(Conf.IsGenesis)
	m["ID"] = Conf.ID
	m["MONIKER"] = Conf.Moniker
	m["KEY"] = Conf.Key
	m["GENESIS_IP"] = Conf.GenesisIP
	m["HOKK_ETH"] = Conf.HokkETH
	m["IP"] = Conf.IP
	m["PORT"] = Conf.PORT
	m["PARTY_PASSWORD"] = Conf.PartyPassword
	m["THRESHOLD"] = strconv.Itoa(Conf.Threshold)
	m["LOG_LEVEL"] = strconv.FormatUint(uint64(Conf.LogLevel), 10)
	m["DB_URL"] = Conf.DbURL
	m["TLS_CERT"] = Conf.TLSCert
	m["TLS_KEY"] = Conf.TLSKey
	m["MAX_CONN"] = strconv.Itoa(Conf.MaxConn)
	m["HSM_CONFIG_PATH"] = Conf.HSM.ConfigPath
	m["HSM_PATH"] = Conf.HSM.Path
	m["HSM_TOKEN_LABEL"] = Conf.HSM.TokenLabel
	m["HSM_PIN"] = Conf.HSM.Pin
	m["HSM_USE_GCM_IV"] = strconv.FormatBool(Conf.HSM.UseGCMIVFromHSM)
	m["HSM_CKA_ID"] = Conf.HSM.CKAID
	m["HSM_CKA_LABEL"] = Conf.HSM.CKALabel
	m["CONTRACT_ADDRESS"] = Conf.CoreContract.ContractAddress
	m["CHAIN_ID"] = strconv.Itoa(Conf.CoreContract.ChainID)
	m["CHAIN_TYPE"] = Conf.CoreContract.ChainType
	m["NAME"] = Conf.CoreContract.Name
	m["EVM_WSS_URL"] = Conf.CoreContract.EvmWSSURL
	m["EVM_HTTP_URL"] = Conf.CoreContract.EvmHTTPURL

	return godotenv.Write(m, ".env")
}
