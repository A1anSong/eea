package config

import "time"

var (
	Configs Config
)

type Config struct {
	Domain string
	Mysql  string
	Jwt    struct {
		Expire time.Duration
		Secret string
	}
	RSA string
}
