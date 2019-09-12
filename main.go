package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"./item"
)

func main() {

	repository := item.NewInMemoryRepository()
	service := item.NewItemService(repository)
	handler := item.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/", handler.Create).Methods("POST")
	router.HandleFunc("/", handler.GetAllItems).Methods("GET")
	router.HandleFunc("/{id}", handler.GetItem).Methods("GET")
	router.HandleFunc("/{id}", handler.DeleteItem).Methods("DELETE")
	router.HandleFunc("/{id}", handler.CompleteItem).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))
}
