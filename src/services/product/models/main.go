package models

import "gorm.io/gorm"

type Brand struct {
    gorm.Model
    Name string `gorm:"type:varchar(100);unique;not null"` 
}

type Category struct {
    gorm.Model
    Name string `gorm:"type:varchar(100);unique;not null"` 
}

type Product struct {
	gorm.Model

	Name string `gorm:"type:varchar(100);not null"`

	BrandId uint `gorm:"index;not null"`
	Brand Brand

	Description string `gorm:"type:text;not null"`

	Price float64 `gorm:"type:numeric(10,2);not null"`

	Discount int `gorm:"type:int"`

	IsAvailable bool `gorm:"type:bool;not null"`

	CategoryId uint `gorm:"index;not null"`

	Category Category
}


