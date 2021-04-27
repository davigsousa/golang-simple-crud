package main

import (
	"log"
	"net/http"

	"github.com/davigsousa/golang-simple-crud/controllers/books"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize new router
	r := mux.NewRouter()

	// Initialize mock data
	books.InitializeMockData()

	// Define Routes
	r.HandleFunc("/api/books", books.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", books.GetBook).Methods("GET")
	r.HandleFunc("/api/books", books.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", books.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", books.DeleteBook).Methods("DELETE")

	// Set listener
	log.Default().Println("Server is listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
