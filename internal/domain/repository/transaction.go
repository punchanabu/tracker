package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/punchanabu/portfolio-tracker/internal/domain/entity"
	"github.com/punchanabu/portfolio-tracker/internal/domain/vo"
)

type TransactionRepository interface {
	Create(ctx context.Context, transaction *entity.Transaction) error
	Update(ctx context.Context, transaction *entity.Transaction) error
	GetByHash(ctx context.Context, hash string) (*entity.Transaction, error)
	GetByWalletID(ctx context.Context, walletID uuid.UUID) ([]*entity.Transaction, error)
	GetByAddress(ctx context.Context, address vo.Address) ([]*entity.Transaction, error)
}
