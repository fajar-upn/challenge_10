package database

import (
	"fmt"
	"log"

	"challenge_10/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	port     = 5432
	dbname   = "bookstores_jwt"
	db       *gorm.DB
)

func Configuration() (*gorm.DB, error) {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connection to database: ", err)
	}

	return db, nil
}

func StartDB() *gorm.DB {
	db, err := Configuration()
	if err != nil {
		panic(err)
	}

	return db
}

func MigrationDB() {
	db, err := Configuration()
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(entity.Book{}, entity.User{})
}
