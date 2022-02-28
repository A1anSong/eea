package model

import (
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

const (
	TransferInit    = "init"
	TransferSuccess = "success"
	TransferFailed  = "failed"

	TransferTypeIn  = "transerIn"
	TransferTypeOut = "transerOut"
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

type Transfer struct {
	*gorm.Model `json:"-"`
	UserID      int64
	Type        string
	Bank        string
	BankAccount string
	Currency    string
	Amount      string
	Status      string
}

func (t *Transfer) CheckValid() (err error) {
	if t.Bank == "" || t.BankAccount == "" || t.Currency == "" || t.Amount == "" {
		err = fmt.Errorf("param empty")
		return
	}
	value, err := strconv.ParseFloat(t.Amount, 64)
	if err != nil {
		err = fmt.Errorf("amount format error:%w", err)
		return
	}
	if value == 0 {
		err = fmt.Errorf("amount can't be 0")
	}
	return
}

func (b *Transfer) GetAmount() (value float64, err error) {
	value, err = strconv.ParseFloat(b.Amount, 64)
	return
}
