package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DBUserName = "your_username"
	DBPassword = "your_password"
	DBName     = "your_db_name"
	DBHost     = "localhost"
	DBPort     = "5432" // Default PostgreSQL port
)

func Connect() *gorm.DB {
	dsn := "host=" + DBHost +
		" user=" + DBUserName +
		" password=" + DBPassword +
		" dbname=" + DBName +
		" port=" + DBPort +
		" sslmode=disable" 
		
	// Open a connection to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	return db
}