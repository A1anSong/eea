package router

import (
	"eea/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	token, err := c.Cookie("eea_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "login required"})
		c.Abort()
	}
	info, err := util.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "login required"})
		c.Abort()
	}
	c.Set("user_info", info)
	c.Next()
}
