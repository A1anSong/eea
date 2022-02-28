package controller

import (
	"eea/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetBalance admin
func SetBalance(c *gin.Context) {
	userID := c.Param("uid")
	nID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "param error"})
		return
	}
	var balance model.Balance
	err = c.Bind(&balance)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "param error"})
		return
	}

	db := model.GetDB()
	var old model.Balance
	ret := db.Where("user_id=?", nID).Find(&old)
	if ret.Error != nil {
		if ret.Error == gorm.ErrRecordNotFound {
			results := db.Create(&balance)
			if results.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("db error:%s", results.Error.Error())})
				return
			}
			return
		}
	}
	old.Balance = balance.Balance
	old.Currency = balance.Currency
	ret = db.Save(&old)
	if ret.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("db error:%s", ret.Error.Error())})
		return
	}
}

// SetInvertStrategy admin
func SetInvertStrategy(c *gin.Context) {

}

// WithDrawConfim admin
func WithDrawConfim(c *gin.Context) {
	idStr := c.Param("id")
	nID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "param error"})
		return
	}
	err = doTransfer(nID, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("data error:%s", err.Error())})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

// TransferInConfim admin
func TransferInConfim(c *gin.Context) {
	idStr := c.Param("id")
	nID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "param error"})
		return
	}
	err = doTransfer(nID, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("data error:%s", err.Error())})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func UpdateUserInfo(c *gin.Context) {
	var user model.User
	err := c.Bind(&user)
	if err != nil || user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "param error"})
		return
	}
	db := model.GetDB()
	var oldUser model.User
	err = db.Find(&oldUser, user.ID).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("user not found:%s", err.Error())})
		return
	}
	oldUser.Role = user.Role
	oldUser.Status = user.Status
	oldUser.AuthLevel = user.AuthLevel
	err = db.Save(&oldUser).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("save user error:%s", err.Error())})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func doTransfer(nID int64, isIn bool) (err error) {
	var transfer model.Transfer
	db := model.GetDB()
	ret := db.First(&transfer, nID)
	if ret.Error != nil {
		err = fmt.Errorf("no transfer found: %d, %s", nID, ret.Error.Error())
		return
	}
	if transfer.Status != model.TransferInit {
		err = fmt.Errorf("transfer status error %s", transfer.Status)
		return
	}
	// maybe add lock
	var balance model.Balance
	ret = db.Where("user_id", transfer.UserID).Find(&balance)
	if ret.Error != nil && ret.Error != gorm.ErrRecordNotFound {
		err = fmt.Errorf("get balance failed: %w", ret.Error)
		return
	}
	balance.UserID = transfer.UserID
	balance.Currency = transfer.Currency
	if isIn {
		if transfer.Type != model.TransferTypeIn {
			err = fmt.Errorf("transfer type not match")
			return
		}
		err = balance.AddBalance(transfer.Amount)
	} else {
		if transfer.Type != model.TransferTypeOut {
			err = fmt.Errorf("transfer type not match")
			return
		}
		err = balance.SubBalance(transfer.Amount)
	}
	if err != nil {
		err = fmt.Errorf("add/sub balance failed: %s", err.Error())
		return
	}
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	transfer.Status = model.TransferSuccess
	if balance.Model == nil {
		err = tx.Create(&balance).Error
	} else {
		err = tx.Save(&balance).Error
	}
	if err != nil {
		err = fmt.Errorf("save balance failed: %w", err)
		return
	}
	err = tx.Save(&transfer).Error
	if err != nil {
		err = fmt.Errorf("save transfer failed: %w", err)
		return
	}
	return
}
