package main

import (
	"fmt"
	"github.com/NicholasLiem/Go_Cassandra_RESTAPI/handlers"
	gorillahandler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serverPort := os.Getenv("PORT")

	r := SetupRoutes()
	headers := gorillahandler.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := gorillahandler.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := gorillahandler.AllowedOrigins([]string{"*"})
	fmt.Println("Running server on port " + serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, gorillahandler.CORS(headers, methods, origins)(r)))
}

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.StatusHandler).Methods("GET")
	r.HandleFunc("/api/books", handlers.CreateBookHandler).Methods("POST")
	r.HandleFunc("/api/books/all", handlers.GetAllBooksHandler).Methods("GET")
	r.HandleFunc("/api/books/{id}", handlers.GetBookHandler).Methods("GET")
	r.HandleFunc("/api/books/{id}", handlers.UpdateBookHandler).Methods("PUT")
	r.HandleFunc("/api/books/{id}", handlers.DeleteBookHandler).Methods("DELETE")
	return r
}

/**
 * Creating keyspace
 */

/**
 * Creating table
 */
//
//CREATE TABLE book(
//	id TEXT PRIMARY KEY,
//	title TEXT,
//	author TEXT,
//	isbn TEXt
//	);
//
//CREATE KEYSPACE IF NOT EXISTS restfulapi
//WITH replication = {
//	'class': 'SimpleStrategy',
//	'replication_factor': 2
//};
