package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/achimonchi/belajar_restapi_mux/src/modules/profile/controller"
	"github.com/achimonchi/belajar_restapi_mux/src/modules/profile/model"
	"github.com/achimonchi/belajar_restapi_mux/src/modules/profile/repository"

	"github.com/achimonchi/belajar_restapi_mux/config"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("REST API with Postgres Golang and Mux")
	handleRequest()
}

func handleRequest() {

	// init route
	r := mux.NewRouter()

	// route handler
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
func getBooks(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetPostgresDB()

	if err != nil {
		fmt.Println(err)
	}

	bookRepository := repository.NewBookRepository(db)

	books, err := controller.GetAll(bookRepository)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
	// fmt.Println("Hit API Book")
}
func getBook(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetPostgresDB()

	if err != nil {
		fmt.Println(err)
	}

	bookRepository := repository.NewBookRepository(db)

	params := mux.Vars(r)

	book, err := controller.GetByID(params["id"], bookRepository)
	fmt.Println(book)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetPostgresDB()

	if err != nil {
		fmt.Println(err)
	}

	bookRepository := repository.NewBookRepository(db)
	var book = model.NewBook()

	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		fmt.Println(err)
	}

	err = controller.SaveBook(book, bookRepository)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetPostgresDB()

	if err != nil {
		fmt.Println(err)
	}

	bookRepository := repository.NewBookRepository(db)
	var book = model.NewBook()

	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("Hit UpdateBook")

	err = controller.Update(book, bookRepository)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetPostgresDB()

	if err != nil {
		fmt.Println(err)
	}

	params := mux.Vars(r)

	bookRepository := repository.NewBookRepository(db)
	err = controller.Delete(params["id"], bookRepository)

}
