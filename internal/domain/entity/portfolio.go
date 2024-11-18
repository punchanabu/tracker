package entity

import "time"

type Portfolio struct {
	ID        string
	UserID    string
	Name      string
	Wallets   []Wallet
	CreatedAt time.Time
	UpdatedAt time.Time
}
