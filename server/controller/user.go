package controller

import (
	"eea/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	user, err := util.User(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"msg": "no user found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
