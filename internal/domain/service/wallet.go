package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/punchanabu/portfolio-tracker/internal/domain/entity"
	"github.com/punchanabu/portfolio-tracker/internal/domain/repository"
	"github.com/punchanabu/portfolio-tracker/internal/domain/vo"
)

type WalletService struct {
	walletRepo     repository.WalletRepository
	portfolioRepo  repository.PortfolioRepository
	balanceService *BalanceService
}

func NewWalletService(walletRepo repository.WalletRepository, portfolioRepo repository.PortfolioRepository, balanceService *BalanceService) *WalletService {
	return &WalletService{
		walletRepo:     walletRepo,
		portfolioRepo:  portfolioRepo,
		balanceService: balanceService,
	}
}

func (s *WalletService) AddWalletToPortfolio(ctx context.Context, portfolioID uuid.UUID, wallet *entity.Wallet) error {
	portfolio, err := s.portfolioRepo.GetByID(ctx, portfolioID)
	if err != nil {
		return err
	}

	existingWallets, err := s.walletRepo.GetByAddress(ctx, wallet.Address)
	if existingWallets != nil {
		return err
	}

	wallet.ID = uuid.New()
	wallet.CreatedAt = time.Now()
	wallet.UpdatedAt = time.Now()

	if err := s.walletRepo.Create(ctx, wallet); err != nil {
		return err
	}

	portfolio.Wallets = append(portfolio.Wallets, *wallet)
	if err := s.portfolioRepo.Update(ctx, portfolio); err != nil {
		return err
	}

	return nil
}

func (s *WalletService) GetWalletBalance(ctx context.Context, walletID uuid.UUID) (vo.Money, error) {
	wallet, err := s.walletRepo.GetByID(ctx, walletID)
	if err != nil {
		return vo.Money{}, err
	}
	return s.balanceService.CalculateWalletBalance(ctx, wallet)
}
