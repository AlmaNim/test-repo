package db

import (
	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Banner struct { //представляет таблицу баннеров в бд со структурой, соответствующей полям таблицы бд
	gorm.Model //Встраивание gorm.Model добавляет поля ID, CreatedAt, UpdatedAt, DeletedAt
	FeatureId  int
	Tags       pq.Int32Array `gorm:"type:integer[]"`
	Content    datatypes.JSON
	IsActive   bool
}

type Auth struct { //представляет таблицу аутентификационных токенов для пользователей с соответствующей структурой полей
	gorm.Model
	Token string
	Role  string
}
