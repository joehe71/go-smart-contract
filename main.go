package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-contracts/api"
	"math/big"
)

func main() {
	// address of etherum env
	client, err := ethclient.Dial("http://127.0.0.1:8545/")
	if err != nil {
		panic(err)
	}

	// create auth and transaction package for deploying smart contract
	auth := getAccountAuth(client, "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")

	//deploying smart contract
	deployedContractAddress, _, _, err := bank.DeployBank(auth, client) //api is redirected from api directory from our contract go file
	if err != nil {
		panic(err)
	}

	fmt.Println(deployedContractAddress.Hex()) // print deployed contract address

	conn, err := bank.NewBank(common.HexToAddress(deployedContractAddress.Hex()), client)
	fmt.Println(conn.GetContractBalance(&bind.CallOpts{}))
}

func getAccountAuth(client *ethclient.Client, accountAddress string) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(accountAddress)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("nounce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(1000000000000000000) // in wei
	auth.GasLimit = uint64(3000000)              // in units
	auth.GasPrice = big.NewInt(875000000)

	return auth
}
