package blockchain

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

// Retrieve the balance of an Ethereum address
func (e *EthereumClient) GetBalance(ctx context.Context, address vo.Address) (*big.Int, error) {
	account := common.HexToAddress(address.String())
	balance, err := e.client.BalanceAt(ctx, account, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

// Retrive transaction details by hash
func (e *EthereumClient) GetTransaction(ctx context.Context, hash string) (*types.Transaction, error) {
	txHash := common.HexToHash(hash)
	tx, isPending, err := e.client.TransactionByHash(ctx, txHash)
	if err != nil {
		return nil, err
	}
	if isPending {
		return nil, ethereum.NotFound
	}
	return tx, nil
}

// Retrieve transaction receipt by hash
func (e *EthereumClient) GetTransactionReceipt(ctx context.Context, hash string) (*types.Receipt, error) {
	txHash := common.HexToHash(hash)
	return e.client.TransactionReceipt(ctx, txHash)
}

// Retreive the lastest block number
func (e *EthereumClient) GetLatestBlock(ctx context.Context) (uint64, error) {
	return e.client.BlockNumber(ctx)
}

// Retrive transaction for an given address
func (e *EthereumClient) GetTransactionByAddress(ctx context.Context, address vo.Address, startBlock, endBlock uint64) ([]*types.Transaction, error) {
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(startBlock),
		ToBlock:   new(big.Int).SetUint64(endBlock),
		Addresses: []common.Address{common.HexToAddress(address.String())},
	}

	logs, err := e.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	var transactions []*types.Transaction
	for _, log := range logs {
		tx, _, err := e.client.TransactionByHash(ctx, log.TxHash)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}

	return transactions, nil
}

func (e *EthereumClient) Close() {
	e.client.Close()
}
