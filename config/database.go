package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbGlobal *gorm.DB

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE"),
	)
	if dbGlobal != nil {
		return dbGlobal
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("init db failed, %s\n", err)
	}
	dbGlobal = db.Debug()
	return dbGlobal
}

func MigrateDB(model ...interface{}) error {
	if dbGlobal != nil {
		return dbGlobal.AutoMigrate(model...)
	}
	return errors.New("db not initialized")
}
