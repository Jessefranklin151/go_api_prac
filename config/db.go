package config

import (
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	return initializePostgresDB()
}
