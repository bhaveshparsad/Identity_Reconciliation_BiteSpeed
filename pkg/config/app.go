package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DBUserName = "fluxkart_unv2_user"
	DBPassword = "5pf4CX9VACERLA6RrYi5iJsivO7t8j5u"
	DBName     = "fluxkart_unv2"
	DBHost     = "dpg-co2ldssf7o1s73cl3000-a"
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