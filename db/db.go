package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	dsn := "host=localhost user=postgres password=1907 dbname=go_banners port=5432"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db = database
	migrate()
}

func migrate() {
	db.AutoMigrate(&Banner{})
	db.AutoMigrate(&Auth{})
}
