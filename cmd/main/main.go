package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	//"gorm.io/driver/postgres"
	"github.com/torcuata22/book-management/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
