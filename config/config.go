package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func ConnectDatabase() { // ConnectDatabase подлючает к базе данных imarket_db
	dsn := "host=localhost user=postgres password=fkla5283 dbname= imarket_db port=5432 sslmode=disable TimeZone=UTC"
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return
	}
}
