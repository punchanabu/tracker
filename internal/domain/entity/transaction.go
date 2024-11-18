package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/punchanabu/portfolio-tracker/internal/domain/vo"
)

type Transaction struct {
	ID        uuid.UUID
	WalletID  uuid.UUID
	Hash      string
	From      vo.Address
	To        vo.Address
	Amount    vo.Money
	Status    TransactionStatus
	TimeStamp time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TransactionStatus string

const (
	TransactionStatusPending TransactionStatus = "PENDING"
	TransactionStatusSuccess TransactionStatus = "SUCCESS"
	TransactionStatusFailed  TransactionStatus = "FAILED"
)
