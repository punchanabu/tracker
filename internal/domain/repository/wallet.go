package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/punchanabu/portfolio-tracker/internal/domain/entity"
	"github.com/punchanabu/portfolio-tracker/internal/domain/vo"
)

type WalletRepository interface {
	Create(ctx context.Context, wallet *entity.Wallet) error
	Update(ctx context.Context, wallet *entity.Wallet) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Wallet, error)
	GetByAddress(ctx context.Context, address vo.Address) (*entity.Wallet, error)
	GetByPortfolioID(ctx context.Context, portfolioID uuid.UUID) ([]*entity.Wallet, error)
}
