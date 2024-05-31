package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(con string, model interface{}) (*gorm.DB, error) {
	database, err := gorm.Open(postgres.Open(con), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := database.AutoMigrate(model); err != nil {
		return nil, err
	}

	return database, nil
}
