package address

import (
	"fmt"
	"math/big"
	"testing"
)

const (
	ChainIDETHMain = 1
	ChainIDLocal   = 1337
)

var client *Service

func init() {
	service, err := New(&Options{
		// ProviderURI: "wss://cloudflare-eth.com/ws",
		// ProviderURI: "https://mainnet.infura.io/v3/d6624913c27d4f5a9540f071767b4d49",
		ProviderURI: "http://127.0.0.1:8545",
	})
	if err != nil {
		panic(err)
	}

	client = service
}

// 以太坊转 ETH
func TestTransferEth(t *testing.T) {
	t.Parallel()
	// t.Skip("Skipping TransferEth") // 用来拦截操作，防止你误操作
	amount := big.NewInt(0)
	// testrpc
	privateKey := "2fdb1c3b367fa0ae7fb174092385616430abd4a67b6f3753f0856afcbe9a1d83"
	toAddress := "0xe5b072d5320dcf3ee3aae76f29704a09dd6fce51"
	tx, err := client.TransferEth(privateKey, toAddress, amount, ChainIDLocal)
	if err != nil {
		t.Errorf("Error transfering eth, got error %s:", err)
	}

	fmt.Printf("%+v\n", tx)
	t.Log(tx)
}
