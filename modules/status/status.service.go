package statusmodule

import (
	"crypto"
	"fmt"
	"math/big"
	networkmodule "telegraph/modules/network"
	"telegraph/pkg/signer/auth"
	"telegraph/pkg/wallet"
)

type Status struct {
	PublicKey        string        `json:"publicKey"`
	Signer           crypto.Signer `json:"signer"`
	MyNetworkTime    *big.Int      `json:"myNetworkTime"`
	TotalNetworkTime *big.Int      `json:"totalNetworkTime"`
}

type StatusService interface {
	GetStatus() (*Status, error)
}

type StatusServiceImpl struct {
	wallet         wallet.Wallet
	auth           auth.IAuth
	networkService networkmodule.NetworkService
}

func (service *StatusServiceImpl) GetStatus() (*Status, error) {
	// conf := config.Conf
	coreNetwork, err := service.networkService.GetNetworkByName("ETH")
	if err != nil {
		fmt.Println("Error getting coreNetwork: ", err)
		return nil, err
	}
	service.auth.SetEVMClient(coreNetwork.EVMHTTPURL)
	pubAddress, err := service.wallet.GetPublicAddress()
	if err != nil {
		fmt.Println("Error getting public address: ", err)
		return nil, err
	}
	// string to common.Address
	// portAddress := common.HexToAddress(coreNetwork.ContractAddress)
	// authObj, err := service.auth.GetAuth(portAddress, pubAddress)
	// if err != nil {
	// 	fmt.Println("Error getting auth object: ", err)
	// }

	// signerInfo, err := authObj.Instance.Signers(&bind.CallOpts{}, pubAddress)
	// if err != nil {
	// 	fmt.Println("Error getting myNetworkTime: ", err)
	// 	return nil, err
	// }
	// fmt.Println("myNetworkTime: ", signerInfo.SignerTime)

	// totalNetworkTime, err := authObj.Instance.NodeStartTime(&bind.CallOpts{})
	// if err != nil {
	// 	fmt.Println("Error getting totalNetworkTime: ", err)
	// 	return nil, err
	// }

	status := &Status{
		PublicKey: pubAddress.String(),
		// MyNetworkTime:    signerInfo.SignerTime,
		// TotalNetworkTime: totalNetworkTime,
	}
	return status, nil
}

func NewStatusService(
	auth auth.IAuth,
	networkService networkmodule.NetworkService,
	wallet wallet.Wallet,
) StatusService {
	return &StatusServiceImpl{
		auth:           auth,
		wallet:         wallet,
		networkService: networkService,
	}
}
