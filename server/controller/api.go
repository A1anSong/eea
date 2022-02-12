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

	db := GetDB()
	var user model.User
	db.First(&user, model.User{
		Email:    email,
		Password: password,
	})
	user.LastLogin = time.Now()
	db.Save(&user)

	token, _ := util.GenToken(&user)

	userCookie := UserCookie{
		Id:     user.ID,
		Role:   user.Role,
		Status: user.Status,
	}

	userCookieJson, _ := json.Marshal(&userCookie)

	c.SetCookie("eea_token", token, int(config.JwtExpireDuration.Seconds()), "/", config.ServerDomain, false, true)
	c.SetCookie("user_id", string(userCookieJson), int(config.JwtExpireDuration.Seconds()), "/", config.ServerDomain, false, false)

	c.JSON(http.StatusOK, user)
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
