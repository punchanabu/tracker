package repository

import (
	"context"
	"math/big"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/punchanabu/portfolio-tracker/internal/domain/vo"
)

type BlockchainRepository interface {
	GetBalance(ctx context.Context, address vo.Address) (*big.Int, error)
	GetTransaction(ctx context.Context, hash string) (*types.Transaction, error)
	GetTransactionReceipt(ctx context.Context, hash string) (*types.Receipt, error)
	GetLatestBlock(ctx context.Context) (uint64, error)
	Close()
}