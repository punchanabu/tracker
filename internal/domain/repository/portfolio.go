package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/punchanabu/portfolio-tracker/internal/domain/entity"
)

type PortfolioRepository interface {
	Create(ctx context.Context, portfolio *entity.Portfolio) error
	Update(ctx context.Context, portfolio *entity.Portfolio) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Portfolio, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.Portfolio, error)
}
