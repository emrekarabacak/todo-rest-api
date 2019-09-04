package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(GetAll())
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adding new Item")
	json.NewEncoder(w).Encode(Add("New Item added"))
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/add", AddItem).Methods("GET")
	router.HandleFunc("/home", HomeHandler).Methods("GET")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
