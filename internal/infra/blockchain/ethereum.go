package blockchain

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/punchanabu/portfolio-tracker/internal/domain/vo"
)

type EthereumClient struct {
	client *ethclient.Client
}

func NewEthereumClient(nodeURL string) (*EthereumClient, error) {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		return nil, err
	}

	return &EthereumClient{client: client}, nil
}

func (e *EthereumClient) GetBalance(ctx context.Context, address vo.Address) (*big.Int, error) {
	account := common.HexToAddress(address.String())
	balance, err := e.client.BalanceAt(ctx, account, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
