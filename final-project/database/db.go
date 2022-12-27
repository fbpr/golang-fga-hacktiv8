package database

import (
	"fmt"
	"log"

	"final-project/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	user     = "developer"
	password = "supersecretpassword"
	dbPort   = "5432"
	dbName   = "mygram"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	err = db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})

	if err != nil {
		log.Fatal("Error migrating models to database: ", err)
	}
}

func GetDB() *gorm.DB {
	return db
}
