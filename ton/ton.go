package ton

import (
	"context"
	"fmt"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"strings"
)

func GetApi() ton.APIClientWrapped {
	client := liteclient.NewConnectionPool()

	configUrl := "https://ton-blockchain.github.io/testnet-global.config.json"
	err := client.AddConnectionsFromConfigUrl(context.Background(), configUrl)
	if err != nil {
		panic(err)
	}
	return ton.NewAPIClient(client).WithRetry()
}

func GetWallet(api ton.APIClientWrapped) {
	words := strings.Split("mix case firm submit inch outdoor enjoy decide trigger tobacco master giggle neck art novel owner physical winner kitten recall eyebrow neutral predict demise", " ")

	w, err := wallet.FromSeed(api, words, wallet.V3)
	if err != nil {
		panic(err)
	}

	block, err := api.CurrentMasterchainInfo(context.Background())
	if err != nil {
		panic(err)
	}
	balance, err := w.GetBalance(context.Background(), block)
	if err != nil {
		panic(err)
	}

	fmt.Println("balance", balance)
}
