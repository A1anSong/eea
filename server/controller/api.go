package controller

import (
	"eea/config"
	"eea/model"
	"eea/util"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserCookie struct {
	Id     uint             `json:"id"`
	Role   model.UserRole   `json:"role"`
	Status model.UserStatue `json:"status"`
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	base64password := c.PostForm("password")
	remember := c.PostForm("remember")
	if email == "" || base64password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "input error"})
		return
	}
	bytePassword, _ := base64.StdEncoding.DecodeString(base64password)
	password, err := util.RSADecrypt(bytePassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	user, err := model.GetUser(email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
		return
	}
	if string(password) != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "invalid password"})
	}
	user.Login()
	token, _ := util.GenToken(user)
	userCookie := UserCookie{
		Id:     user.ID,
		Role:   user.Role,
		Status: user.Status,
	}
	userCookieJson, _ := json.Marshal(&userCookie)
	maxAge := 0
	if remember == "true" {
		maxAge = int(config.Configs.Jwt.Expire.Seconds())
	}
	c.SetCookie("eea_token", token, maxAge, "/", config.Configs.Domain, false, true)
	c.SetCookie("user_info", string(userCookieJson), maxAge, "/", config.Configs.Domain, false, false)
	c.JSON(http.StatusOK, gin.H{"msg": "Login success as " + user.Role})
}
