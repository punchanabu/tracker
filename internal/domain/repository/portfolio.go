package repository

import (
	"context"

	"github.com/punchanabu/portfolio-tracker/internal/domain/entity"
)

type PortfolioRepository interface {
	Create(ctx context.Context, portfolio *entity.Portfolio) error
	Update(ctx context.Context, portfolio *entity.Portfolio) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*entity.Portfolio, error)
	GetByUserID(ctx context.Context, userID string) (*entity.Portfolio, error)
}

