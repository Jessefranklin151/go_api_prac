package config

import (
	"github.com/Jessefranklin151/go_api_prac/go_api_prac/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initializePostgresDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=root dbname=racha-pjb port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	db, openErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if openErr != nil {
		return nil, openErr
	}

	migrateErr := db.AutoMigrate(&schemas.PlayerGormModel{})
	if migrateErr != nil {
		return nil, migrateErr
	}

	return db, nil
}
