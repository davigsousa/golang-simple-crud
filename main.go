package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Models
type ResponseMessage struct {
	Message string `json:"message"`
}

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
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(10000000)) // Not safe
	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	bookID := params["id"]

	bookIndex := -1
	for index, current := range books {
		if current.ID == bookID {
			bookIndex = index
			break
		}
	}

	if bookIndex == -1 {
		json.NewEncoder(w).Encode(&ResponseMessage{
			Message: "Book does not exist",
		})
		return
	}

	books = append(books[:bookIndex], books[bookIndex+1:]...)
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Initialize new router
	r := mux.NewRouter()

	// Initialize mock data
	books = append(books, Book{ID: "1", Isbn: "38427439", Title: "Book 1", Author: &Author{Firstname: "Davi", Lastname: "Sousa"}})
	books = append(books, Book{ID: "2", Isbn: "38427439", Title: "Book 2", Author: &Author{Firstname: "Davi", Lastname: "Sousa"}})
	books = append(books, Book{ID: "3", Isbn: "38427439", Title: "Book 3", Author: &Author{Firstname: "Davi", Lastname: "Sousa"}})

	// Define Routes
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Set listener
	log.Default().Println("Server is listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
