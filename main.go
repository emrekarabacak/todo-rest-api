package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GetAll())
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if len(vars["id"]) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(vars["id"])

	item, err := Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Add(string(body)))
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if len(vars["id"]) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(vars["id"])

	status := make(map[string]bool, 0)
	json.NewDecoder(r.Body).Decode(&status)

	updatedItem, err := Update(id, status["completed"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedItem)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if len(vars["id"]) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(vars["id"])

	err := Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", AddItem).Methods("POST")
	router.HandleFunc("/", GetAllItems).Methods("GET")
	router.HandleFunc("/{id}", GetItem).Methods("GET")
	router.HandleFunc("/{id}", UpdateItem).Methods("PUT")
	router.HandleFunc("/{id}", DeleteItem).Methods("DELETE")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
