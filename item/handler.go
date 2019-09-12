package item

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"../domain"
)

// Handler handler for item requests
type Handler struct {
	service domain.Service
}

// NewHandler returns new handler
func NewHandler(service domain.Service) *Handler {
	return &Handler{service: service}
}

// Create POST request handler
func (handler *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var body domain.ItemRequest
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item, err := handler.service.CreateItem(body.Content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// GetItem GET request handler
func (handler *Handler) GetItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("helloo from get")
	vars := mux.Vars(r)
	idAsString := vars["id"]
	if len(idAsString) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(idAsString)
	item, err := handler.service.FindItemByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

// GetAllItems GET all items request handler
func (handler *Handler) GetAllItems(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(handler.service.FindAll())

}

// DeleteItem DELETE delete item request handler
func (handler *Handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idAsString := vars["id"]
	if len(idAsString) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(idAsString)
	err := handler.service.DeleteByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// CompleteItem PUT update item request handler
func (handler *Handler) CompleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idAsString := vars["id"]
	if len(idAsString) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(idAsString)
	item, err := handler.service.CompleteItem(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}
