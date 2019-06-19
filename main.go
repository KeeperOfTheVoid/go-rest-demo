package main

import (
	"encoding/json"
	"log"
	"net/http"
	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID		string	`json:"id"`
	Isbn	string	`json:"isbn"`
	Title	string	`json:"title"`
	Author	*Author	`json:"author"`
}

// Author Struct
type Author struct {
	Firstname		string	`json:"firstname"`
	Lastname		string	`json:"lastname"`
}

/*
 * A slice is a variable length array
 * A regular array has to have a defined length
 */
// Init books variable as a slice Book struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get parameters
	params := mux.Vars(r) 
	// Loop through books and find with ID
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {
	
}

// Update a Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	
}
// Delete a Book

func deleteBook(w http.ResponseWriter, r *http.Request) {
	
}

func main()  {
	// Init Mux Router
	router := mux.NewRouter()

	// Mock Data
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "The Count of Monte Cristo", Author: &Author{Firstname: "Alexander", Lastname: "Dumas"}})
	books = append(books, Book{ID: "2", Isbn: "847564", Title: "Pride and Prejudice", Author: &Author{Firstname: "Jane", Lastname: "Austen"}})

	// Route Handlers / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Run Server
	log.Fatal(http.ListenAndServe(":8000", router))
}