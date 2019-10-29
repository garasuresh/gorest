package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Year string `json:"year"`
}

var books []Book

func handleRequests() {

	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Print("Server started on 8081 port")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	
	books = append(books, 
		Book{Id: "1", Title: "java", Author: "sun", Year: "2019"},
		Book{Id: "2", Title: "Goroutines", Author: "Mr. Goroutine", Year: "2011"},
		Book{Id: "3", Title: "Golang routers", Author: "Mr. Router", Year: "2012"})
	
	handleRequests()
}

/*
* Get all the books 
*/
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

/*
* Get book by id
*/
func getBook(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	for _, book := range books{
		// fmt.Printf("%+v\n", book)
		if book.Id == params["id"] {
			json.NewEncoder(w).Encode(book)
		}
	}
}

/*
* Add book to existing boosk
*/
func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_= json.NewDecoder(r.Body).Decode(&book)

	books = append(books, book)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
	
}

/*
* Update book 
*/
func updateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	errorstatus := true
	params := mux.Vars(r)

	_ = json.NewDecoder(r.Body).Decode(&book)

	for i,b := range books {
		if b.Id == params["id"] {
			errorstatus = false
			books[i] = book
		}
	}


	if !errorstatus {
		json.NewEncoder(w).Encode(books)
	} else {
		w.WriteHeader(http.StatusNotFound);
		w.Header().Set("Content-Type", "application/json")
		var jsonStr = `{"error":"Requested object not found."}`
		fmt.Fprintf(w, jsonStr)
	}
}

/*
* Delete book 
*/
func deleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete object")
	params := mux.Vars(r)

	for i, book := range books {
		if book.Id == params["id"] {
			books = append(books[:i], books[i+1:]...)
		}
	}
	
	json.NewEncoder(w).Encode(books)
}