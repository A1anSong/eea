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
	*gorm.Model
	UserID      uint
	Type        string
	Bank        string
	BankAccount string
	Currency    string
	Amount      int64
	Status      string
	User        User
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

func GetTransferList(status string, offset, limit int) (transfers []Transfer, total int64, err error) {
	var t Transfer
	db := GetDB().Model(&t)
	if status != "" {
		db = db.Where("Status=?", TransferInit)
	}
	db = db.Preload("User")
	err = db.Count(&total).Error
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	err = db.Offset(offset).Limit(limit).Find(&transfers).Error
	return
}
