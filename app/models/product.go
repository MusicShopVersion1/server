package models

import (
	"gorm.io/gorm"
	"time"
)

//							                             1-n	  n-n     n-n     n-n     n-n    n-n
//#Name, Price, Thumbnail, Description, Year, Quality, Category, Brand, Country, Format, Genre, Style

type Product struct {
	Id          uint    `json:"ID"          gorm:"primary_key"`
	Name        string  `json:"name"        gorm:"unique;not null"          validate:"required,min=4,max=128"`
	Quantity    int     `json:"quantity"    gorm:"not null;default:0"       validate:""`
	Price       float32 `json:"price"       gorm:"not null"                 validate:"required"`
	Discount    float32 `json:"discount"    gorm:"not null;default:0.0"     validate:""`
	Thumbnail   string  `json:"thumbnail"   gorm:""                         validate:""`
	Description string  `json:"description" gorm:""                         validate:""`
	Year        string  `json:"year"        gorm:""                         validate:"len=4"`
	Quality     string  `json:"quality"     gorm:""                         validate:""`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `            gorm:"index"`

	Galleries []Gallery `json:"galleries"       gorm:"foreignKey:ProductId;references:Id"` //Product 1-n Gallery
	Brands    []Brand   `json:"brands"          gorm:"many2many:product_brands"`           //Product n-n Brand
}

func (currProduct *Product) UpdateStruct(newProduct *Product) {
	currProduct.Name = newProduct.Name
	currProduct.Quantity = newProduct.Quantity
	currProduct.Price = newProduct.Price
	currProduct.Discount = newProduct.Discount
	currProduct.Thumbnail = newProduct.Thumbnail
	currProduct.Description = newProduct.Description
	currProduct.Year = newProduct.Year
	currProduct.Quality = newProduct.Quality
}

type Gallery struct {
	Id        uint   `json:"ID"          gorm:"primary_key"`
	Thumbnail string `json:"thumbnail"   gorm:"not null"             validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	ProductId uint
}
