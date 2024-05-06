package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"invoice/process/src/model/entity"
	"log"
	"os"
)

func CreateConnection() *gorm.DB {
	dns := os.Getenv("DB_HOST")
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error: ", err)
	}
	err = db.AutoMigrate(&entity.Order{})
	if err != nil {
		return nil
	}
	return db
}
