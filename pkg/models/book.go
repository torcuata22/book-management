package models

import (
	"github.com/torcuata22/book-management/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title       string  `gorm:"json: name"`
	Author      string  `gorm:"json: author"`
	Description string  `gorm:"json: description"`
	Publisher   string  `gorm:"json: publisher"`
	Status      string  `gorm:"json: status"`
	Price       float64 `gorm:"json: price"`
	Quantity    int     `gorm:"json: quantity"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// model functions needed to talk to the database:
func (b *Book) CreateBook() *Book { //receive b of type Book and return a type Book
	//db.Model(b).NewRecord(b) //NewRecord is not valid with gorm 1.25
	db.Create(&b)
	return b
}
