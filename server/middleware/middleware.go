package middleware

import (
	"eea/model"
	"eea/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	token, err := c.Cookie("eea_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "login required"})
		c.Abort()
		return
	}
	info, err := util.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "login required"})
		c.Abort()
		return
	}
	db := model.GetDB()
	var user model.User
	err = db.First(&user, info.UserId).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "no user found"})
		c.Abort()
		return
	}
	user.Password = ""
	c.Set("user_info", &user)
	c.Next()
}

func Admin(c *gin.Context) {
	userInfo, ok := c.Get("user_info")
	if userInfo == nil || !ok {
		c.JSON(http.StatusForbidden, gin.H{"msg": "admin required"})
		c.Abort()
		return
	}
	info, ok := userInfo.(*model.User)
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"msg": "admin required"})
		c.Abort()
		return
	}
	if info.Role != model.RoleAdmin {
		c.JSON(http.StatusForbidden, gin.H{"msg": "admin required"})
		c.Abort()
		return
	}
	c.Next()
}
