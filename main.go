package main

import (
	"errors"
	"kang-by-xoverse/api/rest"
	"kang-by-xoverse/api/utils"

	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func generateWallet() (string, string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	privateKeyHex := hexutil.Encode(privateKeyBytes)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		return "", "", "", errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyHex := hexutil.Encode(publicKeyBytes)

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return privateKeyHex, publicKeyHex, address, nil
}

func main() {
	// fmt.Println(generateWallet())

	utils.LoadDotEnv()
	rdb, close := utils.GetRedisClient()

	rest.RunRestServer(rdb)

	close()
}
