package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func ConnectDatabase(config *Config) (*gorm.DB, error) {
	//Define Data source
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode,
	)
	//Connect to DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}
