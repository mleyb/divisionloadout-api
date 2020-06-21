package router

import (
	"github.com/gorilla/mux"
	"github.com/mleyb/divisionloadout-api/handlers"
)

// New creates a new router
func New() *mux.Router {
	r := mux.NewRouter()

	handler := handlers.New()

	r.HandleFunc("/loadout", handler.LoadoutHandler).Methods("GET")

	return r
}
