package model

import (
	"fmt"

	"gorm.io/gorm"
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
	UserID      uint
	Type        string
	Bank        string
	BankAccount string
	Currency    string
	Amount      int64
	Status      string
}

func (t *Transfer) CheckValid() (err error) {
	if t.Bank == "" || t.BankAccount == "" || t.Currency == "" || t.Amount == 0 {
		err = fmt.Errorf("param empty")
		return
	}
	return
}

func (b *Transfer) GetAmount() (value int64, err error) {
	value = b.Amount
	return
}
