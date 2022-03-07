package model

import (
	"gorm.io/gorm"
)

type BalanceCurrency string

const (
	CurrencyUSD BalanceCurrency = "USD"
	CurrencyEUR BalanceCurrency = "EUR"
	CurrencyGBP BalanceCurrency = "GBP"
)

type Balance struct {
	*gorm.Model `json:"-"`
	Currency    BalanceCurrency `json:"currency"`
	Balance     int64           `json:"balance"`

	UserID uint
}
