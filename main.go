package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//STRUCT BOOK
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//STRUCT AUTHOR
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//INIT BOOKS VAR AS A SLICE BOOK STRUCT
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	var isFound bool = false

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get Params ID from route

	//loop thru books
	for _, item := range books {
		if item.ID == params["id"] {
			isFound = true
			json.NewEncoder(w).Encode(item)
		}
	}
	if !isFound {
		w.WriteHeader(http.StatusNotFound)
		resp := make(map[string]string)
		resp["message"] = "Book Not Found"
		json.NewEncoder(w).Encode(resp)
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID - Not Safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"] // Mock ID - Not Safe
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
		json.NewEncoder(w).Encode(books)
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	r := mux.NewRouter()

	//MOCK DATA FOR BOOK
	books = append(books, Book{ID: "1", Isbn: "10183023", Title: "BookOne", Author: &Author{
		FirstName: "Jane", LastName: "Doe"},
	})
	books = append(books, Book{ID: "2", Isbn: "221513210", Title: "BookTwo", Author: &Author{
		FirstName: "John", LastName: "Doe"},
	})

	//Route Handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r)) //LOG FATAL IF FAILS RETURNS AN ERROR
}
