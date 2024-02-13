package main

import (
	"context"
	"fmt"
	"log"

	"github.com/NickTaporuk/go-defi/bclient"
	"github.com/NickTaporuk/go-defi/config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// cfg, err := config.LoadConfig("config.yaml")
	cfg := &config.Config{
		Blockchain: config.Blockchain{
			ProviderType: "direct",
			RPC:          "https://eth.llamarpc.com",
		},
	}
	client, err := cfg.EthClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	bc, err := bclient.NewClient(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	// Get price on uniswap
	price, err := bc.EthDaiPrice()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(price)
}
