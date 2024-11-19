package postgres

import (
	"context"

	"github.com/punchanabu/portfolio-tracker/internal/domain/entity"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(ctx context.Context, transaction entity.Transaction) error {
	// TODO: implement create transaction
	return nil
}

func (r *TransactionRepository) Update(ctx context.Context, transaction entity.Transaction) error {
	// TODO: implement update transaction
	return nil
}

func (r *TransactionRepository) GetByHash(ctx context.Context, hash string) (*entity.Transaction, error) {
	// TODO: implement get by hash
	return nil, nil
}

func (r *TransactionRepository) GetByWalletID(ctx context.Context, walletID string) ([]*entity.Transaction, error) {
	// TODO: implement get by wallet id
	return nil, nil
}

func (r *TransactionRepository) GetByAddress(ctx context.Context, address string) ([]*entity.Transaction, error) {
	return nil, nil
}
