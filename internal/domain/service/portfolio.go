package service

import (
	"context"
	"errors"
	"time"

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
	myPortfolio := &entity.Portfolio{
		ID:     uuid.New(),
		UserID: userID,
		Name:   name,
	}

	err := s.portfolioRepo.Create(ctx, myPortfolio)

	return myPortfolio, err
}

func (s *PortfolioService) AddWalletToPortfolio(ctx context.Context, portfolioID uuid.UUID, wallet *entity.Wallet) error {
	portfolio, err := s.portfolioRepo.GetByID(ctx, portfolioID)
	if err != nil {
		return err
	}

	existingWallets, err := s.walletRepo.GetByAddress(ctx, wallet.Address)
	if existingWallets != nil {
		return errors.New("wallet already exists")
	}

	wallet.ID = uuid.New()
	wallet.CreatedAt = time.Now()
	wallet.UpdatedAt = time.Now()

	if err := s.walletRepo.Create(ctx, wallet); err != nil {
		return err
	}

	portfolio.Wallets = append(portfolio.Wallets, *wallet)
	return s.portfolioRepo.Update(ctx, portfolio)
}
