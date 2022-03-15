package config

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Configs = Config{
	Log: LogConfig{
		Level: "info",
		File:  "server.log",
	},
}

type Config struct {
	Domain string
	Mysql  string
	Jwt    struct {
		Expire time.Duration
		Secret string
	}
	RSA string
	Log LogConfig
}

type LogConfig struct {
	Level string
	File  string
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
