package model

import (
	"eea/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var globalDB *gorm.DB

func InitDB() {
	db, err := gorm.Open(mysql.Open(config.Configs.Mysql), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		logrus.Fatal(err.Error())
	}
	globalDB = db
	logrus.Println("db connected success")
	err = globalDB.AutoMigrate(
		&User{},
		&Balance{},
	)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Println("auto migrate success")
}

func GetDB() *gorm.DB {
	return globalDB.Session(&gorm.Session{})
}
