package main

import (
	"github.com/NicholasLiem/Go_Cassandra_RESTAPI/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.StatusHandler).Methods("GET")
	return r
}
