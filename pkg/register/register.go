package register

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"telegraph/config"
	"telegraph/log"
	networkmodule "telegraph/modules/network"
	portabi "telegraph/pkg/crypto_utils/port/evm/abi"
	"telegraph/pkg/db/validator"
	"telegraph/pkg/signer/auth"
	"telegraph/pkg/wallet"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	diContainer "telegraph/container"
)

type RegisterRequest struct {
	EntryFeeAddress string `json:"entry_fee_address"`
}

type RegisterImpl struct {
	networkService networkmodule.NetworkService
}

func CheckForCoreContract(conf config.Config, networkServiceInstance networkmodule.NetworkService) {
	savedNetworks, err := networkServiceInstance.GetNetworks()
	if err != nil {
		log.Fatal("networks.GetNetworks error:", err)
	}
	if len(savedNetworks) == 0 {
		log.Error("No networks to listen to")
		return
	}

	// loop through networks and check if one matches conf.CoreContract.Name
	hasCoreContract := false
	for _, network := range savedNetworks {
		if strings.EqualFold(network.Name, conf.CoreContract.Name) {
			log.Info("Found matching network")
			hasCoreContract = true
		}
	}

	if !hasCoreContract {
		err := networkServiceInstance.AddNetwork(networkmodule.Network{
			ChainID:         conf.CoreContract.ChainID,
			ChainType:       conf.CoreContract.ChainType,
			Name:            conf.CoreContract.Name,
			ContractAddress: conf.CoreContract.ContractAddress,
			EVMWSSURL:       conf.CoreContract.EvmWSSURL,
			EVMHTTPURL:      conf.CoreContract.EvmHTTPURL,
		})
		if err != nil {
			log.Fatal("network.AddNetwork error:", err)
		}
	}

}

func populateValidators() {

}

func IsNodeSigner(conf config.Config, wallet wallet.Wallet) bool {
	// Check if node is a signer
	client := createClient(conf)
	portAddress := common.HexToAddress(conf.CoreContract.ContractAddress)
	instance, err := portabi.NewPortabi(portAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	pubAddr, err := wallet.GetPublicAddress()
	// if err, generate 0x0 address
	if err != nil {
		pubAddr = common.HexToAddress("0x0000000000000000000000000000000000000000")
	}

	fmt.Println("Port Address: ", portAddress)
	fmt.Println("Signer Address: ", pubAddr)
	isSigner, err := instance.IsValidSigner(nil, pubAddr)
	if err != nil {
		log.Fatal(err)
	}

	return isSigner
}

func getAllSigners(conf config.Config) {
	// get all signers
	client := createClient(conf)
	portAddress := common.HexToAddress(conf.CoreContract.ContractAddress)
	instance, err := portabi.NewPortabi(portAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	//loop through signersArr until error
	arrLen := big.NewInt(0)
	signersArr := make([]common.Address, 0)
	for {
		signer, err := instance.SignersArr(&bind.CallOpts{}, arrLen)
		if err != nil {
			// break if error
			break
		}
		signersArr = append(signersArr, signer)
		if len(signersArr) == 0 {
			break
		}
		arrLen.Add(arrLen, big.NewInt(1))
	}
	if err != nil {
		log.Fatal(err)
	}

	//loop through signers
	//check if signer is in db
	//if not, add it
	for _, signer := range signersArr {
		//check if signer is in db
		//if not, add it
		// convert signer to address
		nodeOperator, err := validator.FindValidatorByAddress(signer)
		if err != nil {
			log.Fatal(err)
		}
		//if validator is not in db, add it
		if nodeOperator.Domain == "" {
			signerInfo, err := instance.Signers(&bind.CallOpts{}, signer)
			//add validator to db
			validator.AddValidator(
				nodeOperator.PublicETHAddress,
				signerInfo.Moniker,
				signerInfo.Domain,
				false)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func GetEntryFee(conf config.Config, req RegisterRequest) *big.Int {
	client := createClient(conf)
	portAddress := common.HexToAddress(conf.CoreContract.ContractAddress)
	portInstance, err := portabi.NewPortabi(portAddress, client)
	if err != nil {
		log.Fatal("Unable to create port instance", err)
	}
	tokenAddress := common.HexToAddress(req.EntryFeeAddress)
	entryFee, err := portInstance.EntryFees(&bind.CallOpts{}, tokenAddress)
	if err != nil {
		log.Fatal("Unable to get entry fee", err)
	}

	return entryFee
}

func checkForEntryFees(conf config.Config, wallet wallet.Wallet, req RegisterRequest) (*big.Int, *big.Int) {
	account, err := wallet.GetPublicAddress()
	if err != nil {
		log.Fatal(err)
	}
	client := createClient(conf)
	tokenAddress := common.HexToAddress(req.EntryFeeAddress)
	tokenInstance, err := portabi.NewErc20(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	coinBalance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal("client.BalanceAt error:", err)
	}
	tokenBalance, err := tokenInstance.BalanceOf(&bind.CallOpts{}, account)
	if err != nil {
		log.Fatal("client.BalanceOf error:", err)
	}

	return coinBalance, tokenBalance
}

func StartRegistration(conf config.Config, wallet wallet.Wallet, req RegisterRequest, network networkmodule.NetworkService) {
	// ask for entry fee address
	fmt.Println("Please enter the address of the entry fee token:")
	entryFeeAddress := req.EntryFeeAddress
	_, err := fmt.Scanln(&entryFeeAddress)
	if err != nil {
		log.Fatal("Unable to read entry fee address", err)
	}
	// wait for address to have enough funds(Coin + tokens)
	account, err := wallet.GetPublicAddress()
	if err != nil {
		log.Fatal(err)
	}
	client := createClient(conf)
	var coinBalance, tokenBalance = big.NewInt(0), big.NewInt(0)

	tokenAddress := common.HexToAddress(entryFeeAddress)
	tokenInstance, err := portabi.NewErc20(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	portAddress := common.HexToAddress(conf.CoreContract.ContractAddress)
	portInstance, err := portabi.NewPortabi(portAddress, client)
	if err != nil {
		log.Fatal("Unable to create port instance", err)
	}

	entryFee, err := portInstance.EntryFees(&bind.CallOpts{}, tokenAddress)
	if err != nil {
		log.Fatal("Unable to get entry fee", err)
	}

	// loop every 10s while coinBalance is 0 and tokenBalance equal to or greater than entryFee
	for coinBalance.Cmp(big.NewInt(0)) == 0 || tokenBalance.Cmp(entryFee) < 0 {
		coinBalance, err = client.BalanceAt(context.Background(), account, nil)
		if err != nil {
			log.Fatal("client.BalanceAt error:", err)
		}
		fmt.Println("Coin balance:", coinBalance)
		tokenBalance, err = tokenInstance.BalanceOf(&bind.CallOpts{}, account)
		if err != nil {
			log.Fatal("client.BalanceOf error:", err)
		}
		fmt.Println("Token balance:", tokenBalance)
		time.Sleep(10 * time.Second)
	}

	coreNetwork, err := network.GetNetworkByName("ETH")
	if err != nil {
		log.Fatal("Unable to get core network", err)
	}

	// call addSigner
	makeEntryPayment(coreNetwork, client, tokenAddress, tokenInstance, portInstance, conf, wallet)
}

func makeEntryPayment(coreNetwork networkmodule.Network, client *ethclient.Client, tokenAddress common.Address, tokenInstance *portabi.Erc20, portInstance *portabi.Portabi, conf config.Config, wallet wallet.Wallet) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Gas Price:", gasPrice)
	container := diContainer.GetContainer()

	if err := container.Invoke(func(auth auth.IAuth) {
		auth.SetEVMClient(conf.CoreContract.EvmHTTPURL)
		pubAddress, err := wallet.GetPublicAddress()
		if err != nil {
			log.Fatal(err)
		}

		portContract := common.HexToAddress(coreNetwork.ContractAddress)
		txAuth, err := auth.GetAuth(portContract, pubAddress)
		if err != nil {
			log.Fatal(err)
		}

		// approve token transfer
		feeAmount, err := portInstance.EntryFees(&bind.CallOpts{}, tokenAddress)
		if err != nil {
			log.Fatal("Unable to get entry fee amount", err)
		}
		fmt.Println("Fee amount:", feeAmount)
		fmt.Println("Approving token transfer", txAuth.Auth)
		approveTx, err := tokenInstance.Approve(txAuth.Auth, common.HexToAddress(conf.CoreContract.ContractAddress), feeAmount)
		if err != nil {
			log.Fatal("Unable to approve token transfer", err)
		}
		log.Info("Approve transaction sent: ", approveTx.Hash().Hex())

		// wait for transaction to be mined
		_, err = bind.WaitMined(context.Background(), client, approveTx)
		if err != nil {
			log.Fatal("Unable to wait for approve transaction to be mined", err)
		}

		// call addSigner
		signerAddress, err := wallet.GetPublicAddress()
		if err != nil {
			log.Fatal(err)
		}
		pubAddress, err = wallet.GetPublicAddress()
		if err != nil {
			log.Fatal(err)
		}
		txAuth, err = auth.GetAuth(portContract, pubAddress)
		if err != nil {
			log.Fatal(err)
		}

		addSignerTx, err := portInstance.AddSigner(
			txAuth.Auth,
			signerAddress,
			tokenAddress,
			conf.IP,
			conf.Moniker,
			[]uint8{},
			[][32]byte{},
			[][32]byte{},
			[][32]byte{})
		if err != nil {
			log.Fatal("Unable to add signer", err)
		}
		log.Info("Add signer transaction sent: ", addSignerTx.Hash().Hex())

		// wait for transaction to be mined
		_, err = bind.WaitMined(context.Background(), client, addSignerTx)
		if err != nil {
			log.Fatal("Unable to wait for add signer transaction to be mined", err)
		}

		// log error if transaction failed
		receipt, err := client.TransactionReceipt(context.Background(), addSignerTx.Hash())
		if err != nil {
			log.Fatal("Unable to get transaction receipt", err)
		}
		if receipt.Status == 0 {
			log.Fatal("Transaction failed", err)
		}
	}); err != nil {
		log.Fatal(err)
	}
}

func createClient(conf config.Config) *ethclient.Client {
	client, err := ethclient.Dial(conf.CoreContract.EvmHTTPURL)
	if err != nil {
		log.Fatal("Unable to establish Eth Client", err)
	}
	return client
}

func askForConfirmation(message string) bool {
	var response string
	fmt.Println(message)
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}

	switch strings.ToLower(response) {
	case "y", "yes":
		return true
	case "n", "no":
		return false
	default:
		fmt.Println("Please type (y)es or (n)o and then press enter:")
		return askForConfirmation(message)
	}
}
