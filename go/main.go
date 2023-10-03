package main

import (
	"fmt"
	gorillahandler "github.com/gorilla/handlers"
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
