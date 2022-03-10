package model

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

const (
	TransferInit    = "init"
	TransferSuccess = "success"
	TransferFailed  = "failed"

	TransferTypeIn  = "transerIn"
	TransferTypeOut = "transerOut"
)

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
