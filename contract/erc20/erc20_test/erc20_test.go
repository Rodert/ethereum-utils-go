package erc20test_test

import (
	"ethereum-utils-go/contract/erc20"
	"fmt"
	"testing"
)

const (
	USDTToken = "0xdac17f958d2ee523a2206206994597c13d831ec7"
	Token1    = "0x16e4cd71004b84097a14e6beb9ec8b5093280abf"
	Addr1     = "0x564286362092D8e7936f0549571a803B203aAceD"
	Addr2     = "0xbe0eb53f46cd790cd13851d5eff43d12404d33e8"
	Addr3     = "0xf977814e90da44bfa03b6295a0616a897441acec"
	Addr4     = "0xD999a5234A591EEC281b91deF68C6B6d32e174e6"
)

const (
	// 建议替换自己 project_id，每日使用次数有限
	InfuraNode1 = "https://mainnet.infura.io/v3/d6624913c27d4f5a9540f071767b4d49"
)

// 查询代币余额
// 验证
func TestERC20Balance(t *testing.T) {
	addr := "0xda98f2EfF1D348bD58B6E5d636d8f123DF3AF535"
	erc20.ERC20Balance(addr)
}

// 经过测试不是所有代币都可以拿到：比如：0x16e4cd71004b84097a14e6beb9ec8b5093280abf
// 获取代币信息
func TestGetTokenInfo(t *testing.T) {
	name, symbol, decimals := erc20.GetToken(USDTToken, InfuraNode1)
	fmt.Printf("name: %s\n", name)
	fmt.Printf("symbol: %s\n", symbol)
	fmt.Printf("decimals: %v\n", decimals)
}

func TestGetTokenInfoBalance(t *testing.T) {
	balanceWei, balanceETH := erc20.GetTokenBalance(USDTToken, Addr3, InfuraNode1)
	fmt.Printf("balanceWei: %v\n", balanceWei)
	fmt.Printf("balanceETH: %f\n", balanceETH)
}
