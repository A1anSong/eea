package controller

import (
	"eea/config"
	"eea/model"
	"eea/util"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserCookie struct {
	Id     uint             `json:"id"`
	Role   model.UserRole   `json:"role"`
	Status model.UserStatue `json:"status"`
}

func Login(c *gin.Context) {
	//db := GetDB()
	//initAdmin := model.User{
	//	FirstName: "A1an",
	//	LastName:  "Song",
	//	Email:     "387805107@qq.com",
	//	Password:  "19921201",
	//	Role:      model.RoleAdmin,
	//	Status:    model.StatusActive,
	//	LastLogin: time.Now(),
	//}
	//db.Create(&initAdmin)

	email := c.PostForm("email")
	password := c.PostForm("password")
	remember := c.PostForm("remember")
	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Input error"})
		return
	}
	db := GetDB()
	var user model.User
	if result := db.Where("email = ? AND password = ?", email, password).First(&user); result.RowsAffected == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid Email address or Password"})
		return
	}
	user.LastLogin = time.Now()
	db.Save(&user)

	token, _ := util.GenToken(&user)

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

func RSADecrypt(c *gin.Context) {
	base64Password := c.PostForm("password")
	bytePassword, _ := base64.StdEncoding.DecodeString(base64Password)
	password, err := util.RSADecrypt(bytePassword)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.String(http.StatusOK, string(password))
}
