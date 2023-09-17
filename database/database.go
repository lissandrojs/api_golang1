package database

import (
	"api_golang1/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "user=your_user dbname=your_db sslmode=disable")
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Product{})

	DB = db
	return db, nil
}
