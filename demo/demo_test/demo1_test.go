package demotest_test

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	USDCContractABI = `[{"constant":false,"inputs":[{"name":"newImplementation","type":"address"}],"name":"upgradeTo","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"newImplementation","type":"address"},{"name":"data","type":"bytes"}],"name":"upgradeToAndCall","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":true,"inputs":[],"name":"implementation","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"newAdmin","type":"address"}],"name":"changeAdmin","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"admin","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"_implementation","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"payable":true,"stateMutability":"payable","type":"fallback"},{"anonymous":false,"inputs":[{"indexed":false,"name":"previousAdmin","type":"address"},{"indexed":false,"name":"newAdmin","type":"address"}],"name":"AdminChanged","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"implementation","type":"address"}],"name":"Upgraded","type":"event"}]`
)

func TestDemo(t *testing.T) {
	// Connect to an Ethereum node using Infura
	// client, err := ethclient.Dial(fmt.Sprintf("https://goerli.infura.io/v3/%s", os.Getenv("INFURA_PROJECT_ID")))
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("123")
	// Load the ABI for the ERC-20 token contract
	tokenAddress := common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984")
	abi, err := abi.JSON(strings.NewReader(USDCContractABI))
	if err != nil {
		log.Fatal(err)
	}
	contract := bind.NewBoundContract(tokenAddress, abi, client, client, client)

	// Create a new transaction
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress("0x65f06f9ea7FDc6EB73fd3Ffd543b6bb1ff3C3E72")) // from
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("1234")
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("12345")
	privateKey, err := crypto.HexToECDSA("ffb16b7d73b676c9689accb5c68622cb0960ebc00b6c14a80b75eacafa90a229")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("12346")
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(1)) // from KEY
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// toAddress := common.HexToAddress("0xB605e22c573fADFD76C90969BB80AfDD3C355b1e") // to
	// amount := big.NewInt(1000000000000000000)                                      // in wei
	fmt.Println("123467")
	key, _ := crypto.HexToECDSA("ffb16b7d73b676c9689accb5c68622cb0960ebc00b6c14a80b75eacafa90a229")
	transactOpts, err := bind.NewKeyedTransactorWithChainID(key, new(big.Int).SetInt64(1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%+v\n", transactOpts)

	tx, err := contract.Transfer(transactOpts)
	// tx, err := contract.Transfer(auth, toAddress, amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%+v\n", tx)
	fmt.Println("123468")
	fmt.Printf("Tx Hash: %s", tx.Hash().Hex())
}
