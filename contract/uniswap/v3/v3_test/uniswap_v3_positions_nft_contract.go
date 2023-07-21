package v3_test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	InfuraNode1 = "https://mainnet.infura.io/v3/d6624913c27d4f5a9540f071767b4d49"
	Addr1       = ""
)

// 通过 ETH 购买 USDT
func TestETHToUSDTByUniswap(t *testing.T) {
	client, err := ethclient.Dial(InfuraNode1)
	_ = client
	if err != nil {
		fmt.Println(err)
		return
	}
	// addr1 := common.HexToAddress(Addr1)
	/*
		addr1 := common.HexToAddress(Addr1)
		uvpnct, err := v3.NewUniswapV3PositionsNftContractTransactor(addr1, client)
		if err != nil {
			fmt.Println(err)
		}

		auth := bind.NewKeyedTransactor("privateKey")

		transactOpts := bind.TransactOpts{
			From:     auth.From,
			Signer:   auth,
			GasLimit: uint64(300000),
			GasPrice: big.NewInt(21000000000),
		}

		_ = v3.UniswapV3PositionsNftContractTransactorSession{
			Contract:     uvpnct,
			TransactOpts: transactOpts,
		}
	*/
}
