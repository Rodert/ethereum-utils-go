package erc20test_test

import (
	"context"
	"crypto/ecdsa"
	"ethereum-utils-go/contract/erc20"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	// "github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

const (
	USDTAddr = "0xdac17f958d2ee523a2206206994597c13d831ec7"
)

// 查询代币余额
// 验证
func TestERC20Balance(t *testing.T) {
	addr := "0xda98f2EfF1D348bD58B6E5d636d8f123DF3AF535"
	erc20.ERC20Balance(addr)
}

func TestDemo1(t *testing.T) {
	// import (
	// 	"context"
	// 	"crypto/ecdsa"
	// 	"fmt"
	// 	"log"
	// 	"math/big"

	// 	"github.com/ethereum/go-ethereum"
	// 	"github.com/ethereum/go-ethereum/common"
	// 	"github.com/ethereum/go-ethereum/common/hexutil"
	// 	"github.com/ethereum/go-ethereum/core/types"
	// 	"github.com/ethereum/go-ethereum/crypto"
	// 	"github.com/ethereum/go-ethereum/crypto/sha3"
	// 	"github.com/ethereum/go-ethereum/ethclient"
	// )

	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc

}

func TestERC20Transfer(t *testing.T) {
	addr := "0xda98f2EfF1D348bD58B6E5d636d8f123DF3AF535"
	rawurl := "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
	privateKeyString := "0xffb16b7d73b676c9689accb5c68622cb0960ebc00b6c14a80b75eacafa90a229"

	client, err := ethclient.Dial(rawurl)
	if err != nil {
		fmt.Println(err)
		fmt.Println("123")
		return
	}

	contractNameTransactor, err := erc20.NewContractNameTransactor(common.HexToAddress(addr), client)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 构建调用参数
	privateKey, _ := crypto.HexToECDSA(privateKeyString)
	opts := bind.NewKeyedTransactor(privateKey)
	opts.Value = big.NewInt(0)     // USDT不需要value
	opts.GasLimit = uint64(300000) // 设定gas限制

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	callOpt := &bind.TransactOpts{
		// From:    common.Address{},
		From:    fromAddress,
		Context: context.Background(),
	}
	transaction, err := contractNameTransactor.Transfer(callOpt, common.HexToAddress(USDTAddr), big.NewInt(1000000))
	if err != nil {
		fmt.Println(err)
		fmt.Println("1234")
		return
	}
	fmt.Println("12345")
	fmt.Printf("\n%+v\n", transaction)

}
