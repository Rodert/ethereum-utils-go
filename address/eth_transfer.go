package address

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"runtime/debug"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Service service
type Service struct {
	Client *ethclient.Client
}

// Options service options
type Options struct {
	ProviderURI string
}

// New returns new service
func New(opts *Options) (*Service, error) {
	if opts.ProviderURI == "" {
		return nil, errors.New("ethereum provider uri is required")
	}
	client, err := ethclient.Dial(opts.ProviderURI)
	if err != nil {
		return nil, err
	}
	return &Service{
		Client: client,
	}, nil
}

func (client *Service) TransferEth(privateKey string, _toAddress string, amount *big.Int, chainID int) (*types.Transaction, error) {
	toAddress := common.HexToAddress(_toAddress)
	if chainID == 0 {
		chainID = 1
	}

	signer := types.NewEIP155Signer(big.NewInt(int64(chainID)))
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		debug.PrintStack()
		fmt.Printf("\n%+v\n", err)
		return &types.Transaction{}, err
	}
	fromAddress, err := client.GetPublicAddressFromPrivateKey(key)
	if err != nil {
		debug.PrintStack()
		fmt.Printf("\n%+v\n", err)
		return &types.Transaction{}, err
	}
	nonce, err := client.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		debug.PrintStack()
		fmt.Printf("\n%+v\n", err)
		return &types.Transaction{}, err
	}
	gasLimit := uint64(121000) // standard limit for sending
	gasPrice, err := client.Client.SuggestGasPrice(context.Background())
	if err != nil {
		debug.PrintStack()
		fmt.Printf("\n%+v\n", err)
		return &types.Transaction{}, err
	}

	tx, err := types.SignTx(types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil), signer, key)
	if err != nil {
		debug.PrintStack()
		fmt.Printf("\n%+v\n", err)
		return tx, err
	}

	err = client.SendTx(tx)
	if err != nil {
		debug.PrintStack()
		fmt.Printf("\n%+v\n", err)
		return tx, err
	}
	return tx, nil
}

// GetPublicAddressFromPrivateKey returns public address from private key
func (client *Service) GetPublicAddressFromPrivateKey(priv *ecdsa.PrivateKey) (common.Address, error) {
	var address common.Address
	pub := priv.Public()
	pubECDSA, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		return address, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address = crypto.PubkeyToAddress(*pubECDSA)
	return address, nil
}

// SendTx send a transaction to the network
func (s *Service) SendTx(tx *types.Transaction) error {
	err := s.Client.SendTransaction(context.Background(), tx)
	if err != nil {
		return err
	}
	return nil
}
