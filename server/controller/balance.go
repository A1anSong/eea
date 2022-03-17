package controller

import (
	"eea/model"
	"eea/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TransferInReq(c *gin.Context) {
	var param model.Transfer
	err := c.Bind(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "param error"})
		return
	}
	err = param.CheckValid()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("param error:%s", err.Error())})
		return
	}

	user, err := util.User(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"msg": fmt.Sprintf("no user found %s", err.Error())})
		return
	}
	param.UserID = user.ID
	param.Type = model.TransferTypeIn
	param.Status = model.TransferInit
	err = model.GetDB().Create(&param).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("add transfer failed:%s", err.Error())})
		return
	}
	// TODO: send a msg via telegram
	c.JSON(http.StatusOK, gin.H{"msg": "ok", "data": param})
}

func WithDrawReq(c *gin.Context) {
	var param model.Transfer
	err := c.Bind(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "param error"})
		return
	}
	if param.CheckValid() != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("param error:%s", err.Error())})
		return
	}
	user, err := util.User(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"msg": "no user found"})
		return
	}
	amount, err := param.GetAmount()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "param error"})
		return
	}
	var balance model.Balance
	ret := model.GetDB().Where("user_id=?", user.ID).Find(&balance)
	if ret.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("user balance not found %s", err.Error())})
		return
	}
	available, err := balance.GetBalance()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "user balance error"})
		return
	}
	if amount > available {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "user balance not enough"})
		return
	}

	param.UserID = user.ID
	param.Type = model.TransferTypeOut
	param.Status = model.TransferInit

	ret = model.GetDB().Create(&param)
	if ret.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("add transfer failed:%s", err.Error())})
		return
	}
	// TODO: send a msg via telegram
	c.JSON(http.StatusOK, gin.H{"msg": "ok", "data": param})
}

func Balance(c *gin.Context) {
	user, err := util.User(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"msg": "no user found"})
		return
	}
	var balance model.Balance
	err = model.GetDB().Where("user_id=?", user.ID).Find(&balance).Error
	if err != nil {
		balance.UserID = user.ID
		balance.Currency = "USD"
		balance.Balance = 0
		logrus.Errorf("user balance not found %s", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"data": balance})
}
