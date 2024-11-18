package service

import (
	"context"

	"github.com/punchanabu/portfolio-tracker/internal/domain/entity"
	"github.com/punchanabu/portfolio-tracker/internal/domain/repository"
	"github.com/punchanabu/portfolio-tracker/internal/domain/vo"
)

type BalanceService struct {
	walletRepo      repository.WalletRepository
	transactionRepo repository.TransactionRepository
}

func NewBalanceService(walletRepo repository.WalletRepository, transactionRepo repository.TransactionRepository) *BalanceService {
	return &BalanceService{
		walletRepo:      walletRepo,
		transactionRepo: transactionRepo,
	}
}

func (s *BalanceService) GetWalletBalance(ctx context.Context, wallet entity.Wallet) (vo.Money, error) {
	return vo.Money{}, nil
}

func (s* BalanceService) UpdateWalletBalance(ctx context.Context, wallet entity.Wallet) error {
	return nil
}
