package main

import (
	"log"
	// Need different name then handlers middleware
	h "main/handlers"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/create", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(h.CreateItemHandler))).Methods("POST")
	router.Handle("/move", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(h.MoveItemHandler))).Methods("POST")
	router.Handle("/{id}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(h.GetManyItemsHandler))).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", handlers.CompressHandler(cors.Default().Handler(router))))
}
