package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/achimonchi/belajar_restapi_mux/config"
	"github.com/achimonchi/belajar_restapi_mux/src/modules/profile/model"
	"github.com/achimonchi/belajar_restapi_mux/src/modules/profile/repository"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("REST API with Postgres Golang and Mux")
	handleRequest()
}

var books []model.Books
var book model.NewBook

func handleRequest() {
	db, err := config.GetPostgresDB()

	if err != nil {
		fmt.Println(err)
	}

	bookRepository := repository.NewBookRepository(db)
	// test := "Haai"
	// init route
	r := mux.NewRouter()

	// route handler
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{
		Message: "Hello World",
	})
	// fmt.Println(&test)
}

func getBook(w http.ResponseWriter, r *http.Request) {

}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}
