package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Conn string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := config.Conn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}
