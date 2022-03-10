package model

import (
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"strconv"
)

type Balance struct {
	*gorm.Model `json:"-"`
	UserID      int64
	Currency    string
	Balance     string
}

func (b *Balance) GetBalance() (value float64, err error) {
	value, err = strconv.ParseFloat(b.Balance, 64)
	return
}

func (b *Balance) AddBalance(amount string) (err error) {
	var old decimal.Decimal
	if b.Balance != "" {
		old, err = decimal.NewFromString(b.Balance)
	}
	if err != nil {
		return
	}
	add, err := decimal.NewFromString(amount)
	if err != nil {
		return
	}
	ret := old.Add(add)
	b.Balance = ret.String()
	return
}

func (b *Balance) SubBalance(amount string) (err error) {
	var old decimal.Decimal
	if b.Balance != "" {
		old, err = decimal.NewFromString(b.Balance)
	}
	if err != nil {
		return
	}
	add, err := decimal.NewFromString(amount)
	if err != nil {
		return
	}
	ret := old.Sub(add)
	if ret.IsNegative() {
		err = fmt.Errorf("balance not enough")
		return
	}
	b.Balance = ret.String()
	return
}
