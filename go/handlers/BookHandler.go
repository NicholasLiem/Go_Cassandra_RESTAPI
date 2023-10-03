package handlers

import (
	"encoding/json"
	"github.com/NicholasLiem/Go_Cassandra_RESTAPI/datastore"
	"github.com/NicholasLiem/Go_Cassandra_RESTAPI/models"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateBookHandler(rw http.ResponseWriter, r *http.Request) {
	var book models.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		http.Error(rw, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	query := "INSERT INTO books (id, title, author, isbn) VALUES (?, ?, ?, ?)"
	if err := datastore.Session.Query(query, book.ID, book.Title, book.Author, book.ISBN).Exec(); err != nil {
		http.Error(rw, "Failed to create book"+err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}

func GetBookHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	var book models.Book
	query := "SELECT id, title, author, isbn FROM books WHERE id = ?"
	if err := datastore.Session.Query(query, bookID).Scan(&book.ID, &book.Title, &book.Author, &book.ISBN); err != nil {
		http.Error(rw, "Book not found", http.StatusNotFound)
		return
	}

	jsonData, err := json.Marshal(book)
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	_, err = rw.Write(jsonData)
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func GetAllBooksHandler(rw http.ResponseWriter, r *http.Request) {
	var books []models.Book
	query := "SELECT id, title, author, isbn FROM books"

	iter := datastore.Session.Query(query).Iter()
	defer iter.Close()

	for {
		var book models.Book
		if !iter.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN) {
			break
		}
		books = append(books, book)
	}

	if err := iter.Close(); err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(books)
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	_, err = rw.Write(jsonData)
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func UpdateBookHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	var updatedBook models.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedBook); err != nil {
		http.Error(rw, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	query := "UPDATE books SET title = ?, author = ?, isbn = ? WHERE id = ?"
	if err := datastore.Session.Query(query, updatedBook.Title, updatedBook.Author, updatedBook.ISBN, bookID).Exec(); err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
}

func DeleteBookHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	query := "DELETE FROM books WHERE id = ?"
	if err := datastore.Session.Query(query, bookID).Exec(); err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusNoContent)
}
