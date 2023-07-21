package erc20

import (
	"ethereum-utils-go/contract/erc20/info"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 功能：获取合约信息
// 合约地址；外部地址；节点地址；
func GetToken(contract, nodeUrl string) (string, string, uint8) {
	if nodeUrl == "" {
		nodeUrl = "https://cloudflare-eth.com"
	}
	client, err := ethclient.Dial(nodeUrl)
	if err != nil {
		log.Fatal(err)
	}

	tokenAddress := common.HexToAddress(contract)
	instance, err := info.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	return name, symbol, decimals
}

func GetTokenBalance(contract, addr, nodeUrl string) (*big.Int, *big.Float) {
	if nodeUrl == "" {
		nodeUrl = "https://cloudflare-eth.com"
	}
	client, err := ethclient.Dial(nodeUrl)
	if err != nil {
		log.Fatal(err)
	}

	tokenAddress := common.HexToAddress(contract)
	instance, err := info.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(addr)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	return bal, value
}
