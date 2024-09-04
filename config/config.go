package config

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func DB() *gorm.DB {
	return db
}
