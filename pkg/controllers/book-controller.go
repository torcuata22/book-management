package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/torcuata22/book-management/pkg/models"
	"github.com/torcuata22/book-management/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res) //write the response, a json version of newBooks found in the db
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails) //marshal converts a go struct to json (this is res)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{} //CreateBook is a model type Book
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook() // there is also a function CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	//create a book using the model
	updateBook := &models.Book{}
	//read request body, unmarshall it
	utils.ParseBody(r, updateBook) //we are passing the new data for the book in the request
	vars := mux.Vars(r)
	//get the book id
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	//get the book by id in the database and then update book based on the data passed in the request
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Title != "" {
		bookDetails.Title = updateBook.Title
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Description != "" {
		bookDetails.Description = updateBook.Description
	}
	if updateBook.Publisher != "" {
		bookDetails.Publisher = updateBook.Publisher
	}
	if updateBook.Status != "" {
		bookDetails.Status = updateBook.Status
	}
	if updateBook.Price != 0 {
		bookDetails.Price = updateBook.Price
	}
	if updateBook.Quantity != 0 {
		bookDetails.Quantity = updateBook.Quantity
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	deletedBook := models.DeleteBook(ID)
	res, _ := json.Marshal(deletedBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
