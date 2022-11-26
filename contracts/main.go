package contracts

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	// "github.com/ethereum/go-ethereum/crypto"
	"fmt"
)

var CONTRACT_ADDRESS = "0xC1AfE0D67664570B08B6CBb10233CCD939Bb2De4"
var ABI_PATH = "abi/meta.factory.json"
var ctx = context.Background()

func createClient() *ethclient.Client {
	var (
		ctx         = context.Background()
		url         = "https://rpc-mumbai.maticvigil.com"
		client, err = ethclient.DialContext(ctx, url)
	)

	if err != nil {
		panic(err)
	}

	return client
}

// func GetMetaFactory() {
// 	client := createClient()

// 	instance, err := store.NewStore(address, client)
// 	if err != nil {
// 		panic(err)
// 	}

// 	_ = instance
// }

func CreateMeta() {

	address := common.HexToAddress("0x0f51B0aa02475f1433F9a06081F99Aa05d620672")

	client := createClient()
	balance, err := client.BalanceAt(ctx, address, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("The balance is", balance)
}

func CreateIdentity() {}

func SetIdentityMetadata() {}

func GetIdentity() {}

func CreateAsset() {}

func DeliverAsset() {}
