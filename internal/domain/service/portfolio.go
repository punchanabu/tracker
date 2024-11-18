package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/punchanabu/portfolio-tracker/internal/domain/entity"
	"github.com/punchanabu/portfolio-tracker/internal/domain/repository"
)

type PortfolioService struct {
	portfolioRepo repository.PortfolioRepository
	walletRepo    repository.WalletRepository
}

func NewPortfolioService(portfolioRepo repository.PortfolioRepository, walletRepo repository.WalletRepository) *PortfolioService {
	return &PortfolioService{
		portfolioRepo: portfolioRepo,
		walletRepo:    walletRepo,
	}
}

func (s *PortfolioService) CreatePortfolio(ctx context.Context, userID uuid.UUID, name string) (*entity.Portfolio, error) {
	return nil, nil
}

func (s *PortfolioService) AddWalletToPortfolio(ctx context.Context, portfolioID uuid.UUID, walletID uuid.UUID) error {
	return nil
}