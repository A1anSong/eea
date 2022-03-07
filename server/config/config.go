package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var Configs Config

type Config struct {
	Domain string
	Mysql  string
	Jwt    struct {
		Expire time.Duration
		Secret string
	}
	RSA string
}

func InitConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigName("server")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal(err.Error())
	}
	err = viper.Unmarshal(&Configs)
	if err != nil {
		logrus.Fatal(err.Error())
	}
}
