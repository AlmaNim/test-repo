package main

import (
	"go_banners/api"
	"go_banners/db"
)

func main() {
	api.Init() //инициализация апишки
	db.Init()  //инициализация бд
}
