package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Price       float64
	Discount    float64
	Description string
	Supplier    string
}
