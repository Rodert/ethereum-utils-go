package demotest_test

import (
	"context"
	"fmt"
	"math/big"
	"mysolidity/json"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestDemoQuery(t *testing.T) {
	//合约地址
	contractAddress := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	//websocket监听
	client, err := ethclient.Dial("ws://127.0.0.1:8545/ws")
	if err != nil {
		fmt.Errorf("could not connect to local node: %v", err)
		return
	}
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(1), //生产环境中，从0开始，查询后修改区块记录，下一次就从后一个有记录的区块数开始
		ToBlock:   big.NewInt(100),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	erc20, _ := json.NewErc20(contractAddress, client)
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		fmt.Errorf("err:%v\n", err)
		return
	}
	for _, vLog := range logs {
		if len(vLog.Topics) == 0 {
			continue
		}
		event := vLog.Topics[0].Hex()
		if event == TransferEvent() { //对对应的事件进行对应的处理
			fmt.Println(vLog.Data)
			data, err := erc20.ParseTransfer(vLog)
			if err != nil {
				fmt.Errorf("err:%v\n", err)
				continue
			}
			fmt.Println(data.From.Hex(), data.To.Hex(), data.Value.Int64(), data.Raw.Data)
		}
	}
}
func TransferEvent() string {
	event := crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")).Hex()
	return event
}
