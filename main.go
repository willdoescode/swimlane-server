package main

import (
	"log"
	h "main/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/create", h.CreateItem).Methods("POST")
	router.HandleFunc("/{id}", h.GetManyItems).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", cors.Default().Handler(router)))
}
