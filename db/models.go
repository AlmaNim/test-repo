package db

import (
	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Banner struct {
	gorm.Model
	FeatureId int
	Tags      pq.Int32Array `gorm:"type:integer[]"`
	Content   datatypes.JSON
	IsActive  bool
}

type Auth struct {
	gorm.Model
	Token string
	Role  string
}
