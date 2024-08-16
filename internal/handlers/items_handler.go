package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/sangharshseth/internal/models"
)

// Dummy data
var items = []models.Item{
    {ID: "1", Name: "Item One"},
    {ID: "2", Name: "Item Two"},
}

// GetItems handles GET requests for /items
func GetItems(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

// GetItemByID handles GET requests for /items/{id}
func GetItemByID(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    for _, item := range items {
        if item.ID == id {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    http.NotFound(w, r)
}

// CreateItem handles POST requests for /items
func CreateItem(w http.ResponseWriter, r *http.Request) {
    var newItem models.Item
    if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    items = append(items, newItem)
    w.WriteHeader(http.StatusCreated)
}

// SetupRoutes sets up the routes for the API
func SetupRoutes(mux *http.ServeMux) {
    mux.HandleFunc("GET /items", GetItems)
    mux.HandleFunc("POST /items/create", CreateItem)
    mux.HandleFunc("GET /items/{id}", GetItemByID)
}