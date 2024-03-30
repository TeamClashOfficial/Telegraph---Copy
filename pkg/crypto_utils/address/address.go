package address

import (
	"regexp"
	"telegraph/log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetAddress() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	log.Info("SAVE BUT DO NOT SHARE THIS (Private Key):", hexutil.Encode(privateKeyBytes))
	// Store private key in file
	// Store public key in DB
}

func CheckAddress(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}
