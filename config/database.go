package config

import (
	"fmt"
	"log"

	"github.com/ThaksinCK/go-basic-api.git/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

var BasicApiDatabase Store

func InitializeDatabaseConnection() {
	ConnectionMasterDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		common.User,
		common.Password,
		common.DatabaseHost,
		common.DatabasePort,
		common.DatabaseName)
	db, err := gorm.Open(mysql.Open(ConnectionMasterDB), &gorm.Config{})
	if err != nil {
		log.Fatalf("[Database] Failed to connect database : %s\n", err.Error())
	}
	fmt.Printf("[Database] Successfully connected at : %s\n", ConnectionMasterDB)
	BasicApiDatabase.DB = db
}
