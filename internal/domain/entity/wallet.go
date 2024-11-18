package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/punchanabu/portfolio-tracker/internal/domain/vo"
)

type Wallet struct {
	ID        uuid.UUID
	Address   vo.Address
	Chain     string
	Balance   vo.Money
	CreatedAt time.Time
	UpdatedAt time.Time
}
