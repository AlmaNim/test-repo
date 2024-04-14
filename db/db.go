package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB //переменная, хранящая подключение к бд

// функция инициализирует подключение к бд с использованием gorm
func Init() {
	dsn := "host=localhost user=postgres password=1907 dbname=go_banners port=5432" //параметры подключения к бд
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database") //паника выбрасывается в случае неудачи при подключении
	}
	db = database
	migrate() //вызов функции миграции для обновления схемы бд
}

// функция выполняет миграцию бд для моделей banner и auth
func migrate() {
	db.AutoMigrate(&Banner{})
	db.AutoMigrate(&Auth{})
}
