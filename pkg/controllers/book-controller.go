package controllers

import (
	"net/http"
)

var (
	CreateBook  = func(w http.ResponseWriter, r *http.Request) {}
	GetBook     = func(w http.ResponseWriter, r *http.Request) {}
	GetBookById = func(w http.ResponseWriter, r *http.Request) {}
	UpdateBook  = func(w http.ResponseWriter, r *http.Request) {}
	DeleteBook  = func(w http.ResponseWriter, r *http.Request) {}
)
