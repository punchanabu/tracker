package service

import (
	"context"
	"math/big"
	"time"

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

func (s *BalanceService) CalculateWalletBalance(ctx context.Context, wallet *entity.Wallet) (vo.Money, error) {
	transactions, err := s.transactionRepo.GetByWalletID(ctx, wallet.ID)
	if err != nil {
		return vo.Money{}, err
	}

	balance := vo.NewMoney(big.NewInt(0), wallet.Balance.Currency)

	for _, tx := range transactions {
		if tx.Status != entity.TransactionStatusSuccess {
			continue
		}

		if tx.From == wallet.Address {
			balance = balance.Sub(tx.Amount)
		}

		if tx.To == wallet.Address {
			balance = balance.Add(tx.Amount)
		}
	}

	return balance, nil
}

func (s *BalanceService) UpdateWalletBalance(ctx context.Context, wallet *entity.Wallet) error {
	balance, err := s.CalculateWalletBalance(ctx, wallet)
	if err != nil {
		return err
	}

	wallet.Balance = balance
	wallet.UpdatedAt = time.Now()

	return s.walletRepo.Update(ctx, wallet)
}
