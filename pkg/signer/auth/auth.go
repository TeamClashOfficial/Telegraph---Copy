package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"net/url"
	portabi "telegraph/pkg/crypto_utils/port/evm/abi"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IAuth interface {
	createClient(evmHttpUrl string) *ethclient.Client
	GetAuth(portContract common.Address, publicKey common.Address) (*AuthObj, error)
	SetEVMClient(evmHttpUrl string)
}

type AuthImpl struct {
	client *ethclient.Client
}

type AuthObj struct {
	Client   *ethclient.Client
	Auth     *bind.TransactOpts
	Instance *portabi.Portabi
}

func (authInstance *AuthImpl) createClient(evmHttpUrl string) *ethclient.Client {
	fmt.Println("evmHttpUrl: ", evmHttpUrl)

	client, err := ethclient.Dial(evmHttpUrl)
	if err != nil {
		log.Fatal("Unable to establish Eth Client", err)
	}
	return client
}

func getOptsFromKeyServer(chainID *big.Int) (*bind.TransactOpts, error) {
	baseURL := "http://localhost:7045/signer/transaction-opts"
	params := url.Values{}
	params.Add("chainID", chainID.String())
	url := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Opts resp err: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Opts body err: ", err)
		return nil, err
	}

	fmt.Println(body)

	var opts *bind.TransactOpts
	err = json.Unmarshal(body, &opts)
	if err != nil {
		fmt.Println("Opts unmarshal err: ", err)
		return nil, err
	}

	return opts, nil
}

func (authInstance *AuthImpl) GetAuth(portContract common.Address, publicAddress common.Address) (*AuthObj, error) {
	chainID, err := authInstance.client.NetworkID(context.Background())
	fmt.Println("chainID: ", chainID)
	if err != nil {
		return nil, err
	}
	auth, err := getOptsFromKeyServer(chainID)
	if err != nil {
		fmt.Println("Auth err: ", auth)
		return nil, err
	}
	fmt.Println("Opts auth: ", auth)
	gasPrice, err := authInstance.client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("Gas price err: ", gasPrice)
		return nil, err
	}
	fmt.Println("gasPrice: ", gasPrice)
	fromAddress := publicAddress
	nonce, err := authInstance.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}
	fmt.Println("Nonce: ", nonce)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	toAddress := portContract
	instance, err := portabi.NewPortabi(toAddress, authInstance.client)
	if err != nil {
		return nil, err
	}
	return &AuthObj{
		Client:   authInstance.client,
		Auth:     auth,
		Instance: instance,
	}, nil
}

func (authInstance *AuthImpl) SetEVMClient(evmHttpUrl string) {
	authInstance.client = authInstance.createClient(evmHttpUrl)
}

func NewAuth(client *ethclient.Client) IAuth {
	return &AuthImpl{
		client,
	}
}
