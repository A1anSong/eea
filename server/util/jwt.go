package util

import (
	"eea/config"
	"eea/model"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
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
				Time: time.Now().Add(config.JwtExpireDuration),
			},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, eeaToken)
	return token.SignedString([]byte(config.JwtSecrect))
}

func ParseToken(tokenString string) (*EEAToken, error) {
	token, err := jwt.ParseWithClaims(tokenString, &EEAToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecrect), nil
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
