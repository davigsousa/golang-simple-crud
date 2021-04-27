package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Models
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Create mock books slice
var books []Book

// Controller Functions
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {

}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Initialize new router
	r := mux.NewRouter()

	// Initialize mock data
	books = append(books, Book{ID: "1", Isbn: "38427439", Title: "Book 1", Author: &Author{Firstname: "Davi", Lastname: "Sousa"}})
	books = append(books, Book{ID: "2", Isbn: "38427439", Title: "Book 2", Author: &Author{Firstname: "Davi", Lastname: "Sousa"}})
	books = append(books, Book{ID: "3", Isbn: "38427439", Title: "Book 3", Author: &Author{Firstname: "Davi", Lastname: "Sousa"}})
	books = append(books, Book{ID: "4", Isbn: "38427439", Title: "Book 4", Author: &Author{Firstname: "Davi", Lastname: "Sousa"}})
	books = append(books, Book{ID: "5", Isbn: "38427439", Title: "Book 5", Author: &Author{Firstname: "Davi", Lastname: "Sousa"}})
	books = append(books, Book{ID: "6", Isbn: "38427439", Title: "Book 6", Author: &Author{Firstname: "Davi", Lastname: "Sousa"}})

	// Define Routes
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Set listener
	log.Fatal(http.ListenAndServe(":8000", r))
}
