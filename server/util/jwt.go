package util

import (
	"eea/config"
	"eea/model"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type EEAToken struct {
	UserId uint
	Role   model.UserRole
	Status model.UserStatue
	jwt.RegisteredClaims
}

func GenToken(user *model.User) (string, error) {
	eeaToken := EEAToken{
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "EEACash",
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(config.Configs.Jwt.Expire),
			},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, eeaToken)
	return token.SignedString([]byte(config.Configs.Jwt.Secret))
}

func ParseToken(tokenString string) (*EEAToken, error) {
	token, err := jwt.ParseWithClaims(tokenString, &EEAToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Configs.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if eeaToken, ok := token.Claims.(*EEAToken); ok && token.Valid {
		return eeaToken, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

func User(c *gin.Context) (u *model.User, err error) {
	value, ok := c.Get("user_info")
	if !ok {
		err = fmt.Errorf("no user found")
		return
	}
	u, ok = value.(*model.User)
	if !ok {
		err = fmt.Errorf("user data error")
		return
	}
	return
}
