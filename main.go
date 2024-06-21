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
	client, err := ethclient.Dial("http://127.0.0.1:8545/")
	if err != nil {
		panic(err)
	}

	// 部署合约
	contractAddress := deployContract(client)
	//连接合约，输出合约余额
	contract := connectContract(client, contractAddress)
	balance, err := contract.GetContractBalance(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Println("合约余额:", balance, "wei")
}

func connectContract(client *ethclient.Client, address common.Address) *bank.Bank {
	conn, err := bank.NewBank(common.HexToAddress(address.Hex()), client)
	if err != nil {
		panic(err)
	}
	return conn
}

func deployContract(client *ethclient.Client) common.Address {
	auth := getAccountAuth(client, "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d")

	deployedContractAddress, _, _, err := bank.DeployBank(auth, client)
	if err != nil {
		panic(err)
	}

	fmt.Println(deployedContractAddress.Hex())

	return deployedContractAddress
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
	auth.Value = big.NewInt(1000000000000000000) // wei
	auth.GasLimit = uint64(3000000)              // units
	auth.GasPrice = big.NewInt(875000000)

	return auth
}
