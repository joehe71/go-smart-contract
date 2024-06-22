package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-contracts/eth"
	"go-contracts/ton"
)

func main() {
	callEth()
}

func callTon() {
	api := ton.GetApi()
	ton.GetWallet(api)
}

func callEth() {
	client, err := ethclient.Dial("http://127.0.0.1:8545/")
	if err != nil {
		panic(err)
	}

	// 部署合约
	contractAddress := eth.DeployContract(client)
	//连接合约，输出合约余额
	contract := eth.ConnectContract(client, contractAddress)
	balance, err := contract.GetContractBalance(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Println("合约余额:", eth.ToEth(balance))

	//查询账户
	eth.GetAccountInfo(client)
}
