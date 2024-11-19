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
	return r.db.WithContext(ctx).Create(&transaction).Error
}

func (r *TransactionRepository) Update(ctx context.Context, transaction entity.Transaction) error {
	return r.db.WithContext(ctx).Save(&transaction).Error
}

func (r *TransactionRepository) GetByHash(ctx context.Context, hash string) (*entity.Transaction, error) {
	var transaction entity.Transaction
	err := r.db.WithContext(ctx).Where("hash = ?", hash).First(&transaction).Error
	return &transaction, err
}

func (r *TransactionRepository) GetByWalletID(ctx context.Context, walletID string) ([]*entity.Transaction, error) {
	var transactions []*entity.Transaction
	err := r.db.WithContext(ctx).Where("wallet_id = ?", walletID).Find(&transactions).Error
	return transactions, err
}

func (r *TransactionRepository) GetByAddress(ctx context.Context, address string) ([]*entity.Transaction, error) {
	var transactions []*entity.Transaction
	err := r.db.WithContext(ctx).Where("address = ?", address).Find(&transactions).Error
	return transactions, err
}
