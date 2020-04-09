package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

}

func handleRequest() {
	// init route
	r := mux.NewRouter()

	// route handler
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
