package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := SetupRoutes()
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:" + os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Listening server on port: " + os.Getenv("PORT"))
	log.Fatal(srv.ListenAndServe())
}
