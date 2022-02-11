package controller

import (
	"eea/config"
	"eea/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var globalDB *gorm.DB

func ConnectDB() {
	dsn := config.DBUser + ":" +
		config.DBPassword + "@tcp(" +
		config.DBAddress + ":" +
		config.DBPort + ")/" +
		config.DBDatabase + "?charset=" +
		config.DBCharset + "&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	globalDB = db

	log.Println("db connected success")

	err = globalDB.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("auto migrate success")
}

func GetDB() *gorm.DB {
	return globalDB.Session(&gorm.Session{})
}
