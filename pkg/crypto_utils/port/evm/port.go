package evmport

import (
	"telegraph/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	portabi "telegraph/pkg/crypto_utils/port/evm/abi"
)

func CreateInstance() (*portabi.Portabi, error) {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/245bd056a0564a4680c8fb5f3d5aa09a")
	if err != nil {
		log.Fatal("ethclient.Dial", err)
	}

	address := common.HexToAddress("0x58Bcf4016b12d1157Ec12059527926f34Ec7FB48")
	instance, err := portabi.NewPortabi(address, client)
	if err != nil {
		log.Fatal("portabi.NewPortabi error:", err)
	}
	log.Debug("Instance: ", instance)
	_ = instance // we'll be using this in the next section

	return instance, nil
}

func GetName(instance *portabi.Portabi) string {
	name, err := instance.Name(nil)
	if err != nil {
		log.Fatal("instance.Name error:", err)
	}

	log.Debug(name) // "1.0"
	return name
}

func IsAddressSigner(instance *portabi.Portabi, address common.Address) (bool, error) {
	status, err := instance.IsValidSigner(nil, address)
	if err != nil {
		return false, err
	}

	log.Debug("Is this address a valid signer: ", status) // "1.0"
	return status, nil
}
