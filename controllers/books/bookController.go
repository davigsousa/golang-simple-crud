package books

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/davigsousa/golang-simple-crud/models"
	"github.com/gorilla/mux"
)

// Create mock books slice
var books []models.Book

func InitializeMockData() {
	books = append(books, models.Book{ID: "1", Isbn: "38427439", Title: "Book 1", Author: &models.Author{Firstname: "Davi", Lastname: "Sousa"}})
	books = append(books, models.Book{ID: "2", Isbn: "38427439", Title: "Book 2", Author: &models.Author{Firstname: "Davi", Lastname: "Sousa"}})
	books = append(books, models.Book{ID: "3", Isbn: "38427439", Title: "Book 3", Author: &models.Author{Firstname: "Davi", Lastname: "Sousa"}})
}

// Controller Functions
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	json.NewEncoder(w).Encode(&models.ResponseMessage{
		Message: "Book does not exist.",
	})
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		json.NewEncoder(w).Encode(&models.ResponseMessage{
			Message: "Error on create book",
		})
		return
	}

	book.ID = strconv.Itoa(rand.Intn(10000000)) // Not safe
	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
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
		json.NewEncoder(w).Encode(&models.ResponseMessage{
			Message: "Book does not exist",
		})
		return
	}

	books = append(books[:bookIndex], books[bookIndex+1:]...)
	json.NewEncoder(w).Encode(books)
}
