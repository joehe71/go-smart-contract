package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	bank "go-contracts/api"
	"log"
	"math/big"
)

func GetAccountInfo(client *ethclient.Client) {
	address := common.HexToAddress("0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC")
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Failed to get account balance: %v", err)
	}
	fmt.Println("账户余额:", ToEth(balance))

	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatalf("Failed to get account nonce: %v", err)
	}
	fmt.Println("账户nonce:", nonce)

}

func ConnectContract(client *ethclient.Client, address common.Address) *bank.Bank {
	conn, err := bank.NewBank(common.HexToAddress(address.Hex()), client)
	if err != nil {
		panic(err)
	}
	return conn
}

func DeployContract(client *ethclient.Client) common.Address {
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

func ToEth(wei *big.Int) *big.Float {
	ethValue := new(big.Float).SetInt(wei)
	ethValue.Quo(ethValue, big.NewFloat(1e18))
	return ethValue
}
