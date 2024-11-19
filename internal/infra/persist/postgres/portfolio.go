package postgres

import (
	"context"

	"github.com/punchanabu/portfolio-tracker/internal/domain/entity"
	"gorm.io/gorm"
)

type PortfolioRepository struct {
	db *gorm.DB
}

func NewPortfolioRepository(db *gorm.DB) *PortfolioRepository {
	return &PortfolioRepository{db: db}
}

func (r *PortfolioRepository) Create(ctx context.Context, portfolio *entity.Portfolio) error {
	return r.db.WithContext(ctx).Create(portfolio).Error
}

func (r *PortfolioRepository) Update(ctx context.Context, portfolio *entity.Portfolio) error {
	return r.db.WithContext(ctx).Save(portfolio).Error
}

func (r *PortfolioRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&entity.Portfolio{}, id).Error
}

func (r *PortfolioRepository) GetByID(ctx context.Context, id string) (*entity.Portfolio, error) {
	var portfolio entity.Portfolio
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&portfolio).Error
	if err != nil {
		return nil, err
	}
	return &portfolio, nil
}

func (r *PortfolioRepository) GetByUserID(ctx context.Context, userID string) (*entity.Portfolio, error) {
	// TODO: implement get by user id
	return nil, nil
}