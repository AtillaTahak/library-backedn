package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/atillatahak/library-backend/models"
)

func SetupDatabase() *gorm.DB {

	dsn := "postgresql://user:password@localhost:5432/database_name?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.Book{})

	return db
}
