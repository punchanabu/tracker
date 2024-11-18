package vo

import (
	"math/big"
)

type Money struct {
	Amount   *big.Int
	Currency string
}

func NewMoney(amount *big.Int, currency string) Money {
	return Money{
		Amount:   amount,
		Currency: currency,
	}
}

func (m Money) Add(money Money) Money {
	if m.Currency != money.Currency {
		panic("cannot add different currency together!")
	}

	return Money{
		Amount:   new(big.Int).Add(m.Amount, money.Amount),
		Currency: m.Currency,
	}
}

func (m Money) Sub(money Money) Money {
	if m.Currency != money.Currency {
		panic("cannot subtract different currency together!")
	}

	return Money{
		Amount:   new(big.Int).Sub(m.Amount, money.Amount),
		Currency: m.Currency,
	}
}

func (m Money) IsZero() bool {
	return m.Amount.Sign() == 0
}
