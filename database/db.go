package database

import (
	"assignment2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "root"
	dbname   = "assignment2FGA"
	port     = 5432
	db       *gorm.DB
	err      error
)

func StartDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database : ", err)
	}

	db.Debug().AutoMigrate(&models.Order{}, &models.Item{})
}

func GetDB() *gorm.DB {
	return db
}
