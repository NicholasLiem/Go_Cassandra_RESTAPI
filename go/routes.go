package main

import (
	"github.com/NicholasLiem/Go_Cassandra_RESTAPI/Handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", Handlers.StatusHandler).Methods("GET")
	return r
}
