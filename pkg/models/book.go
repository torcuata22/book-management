package models

import (
	"log"

	"github.com/torcuata22/book-management/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title       string  `gorm:"column:title" json:"title"`
	Author      string  `gorm:"column:author" json:"author"`
	Description string  `gorm:"column:description" json:"description"`
	Publisher   string  `gorm:"column:publisher" json:"publisher"`
	Status      string  `gorm:"column:status" json:"status"`
	Price       float64 `gorm:"column:price" json:"price"`
	Quantity    int     `gorm:"column:quantity" json:"quantity"`
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

func GetAllBooks() []Book { //use a slice to return a slice
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=>?", id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Find(&book) //find the book so I can return the deleted book data
	db.Delete(&book)
	return book
	//return db makes more sense
}

// TODO: FIX the Panic! in this function
// SoldBook: when a book is sold it subtracts one from the quantity
func SoldBook(id int64, boughtAmount int) Book {
	var book Book
	err := db.Where("ID=?", id).Find(&book)
	if err != nil {
		log.Println(err)
		return Book{}
	}

	log.Println("Before update:", book.Quantity)

	if book.Quantity < boughtAmount {
		// handle the case where there's not enough stock
		log.Println("Not enough stock")
		return Book{}
	}

	book.Quantity -= boughtAmount
	db.Model(&book).Update("Quantity", book.Quantity)

	log.Println("After update:", book.Quantity)

	db.Save(book)
	return book
}

//to call SoldBook: SoldBook(1,2) this would sell 2 copies of book with id 1

//update happens when we find the book id and then delete it and then create new record with new values

// Here's a quick status report:

// **Current Status:**

// * We have a `SoldBook` function that is supposed to update the quantity of a book in the database.
// * The function is not correctly updating the book quantity, and is returning an empty book object.
// * We have tried logging statements to debug the issue, but the problem persists.

// Feel free to come back to this whenever you're ready, and we can pick up where we left off. I'll do my best to help you resolve the issue and get your book management system up and running!
