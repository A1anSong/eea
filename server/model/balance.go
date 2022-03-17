package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Balance struct {
	*gorm.Model
	UserID   uint   `gorm:"uniqueIndex:udx_balance,sort:desc"`
	Currency string `gorm:"uniqueIndex:udx_balance,sort:desc"`
	Balance  int64  // money = Balance/10000
	User     User
}

func (b *Balance) GetBalance() (value int64, err error) {
	value = b.Balance
	return
}

func (b *Balance) AddBalance(amount int64) (err error) {
	b.Balance += amount
	return
}

func (b *Balance) SubBalance(amount int64) (err error) {
	if b.Balance < amount {
		err = fmt.Errorf("balance not enough")
		return
	}
	b.Balance -= amount
	return
}

func GetBalanceList(email string, offset, limit int) (balances []Balance, total int64, err error) {

	var b Balance
	db := GetDB().Model(&b)
	if email != "" {
		db2 := GetDB()
		var userIDs []uint
		err = db2.Raw("select id from users where email like ?", fmt.Sprintf("%%%s%%", email)).Scan(&userIDs).Error
		if err != nil {
			return
		}
		db = db.Where("user_id in ?", userIDs)
	}
	db = db.Preload("User")
	err = db.Count(&total).Error
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	err = db.Offset(offset).Limit(limit).Find(&balances).Error
	return
}
