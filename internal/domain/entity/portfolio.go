package entity

import (
	"time"
	"github.com/google/uuid"
)

type Portfolio struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Name      string
	Wallets   []Wallet
	CreatedAt time.Time
	UpdatedAt time.Time
}
