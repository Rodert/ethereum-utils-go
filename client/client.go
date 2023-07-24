package client

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetClient(nodeUrl string) {
	if nodeUrl == "" {
		nodeUrl = "https://cloudflare-eth.com"
	}

	client, err := ethclient.Dial(nodeUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client
}
