package router

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	api "github.com/mleyb/divisionloadout-api/handlers"
)

// New creates a new router
func New() *mux.Router {
	r := mux.NewRouter()

	handler := api.New()

	r.HandleFunc("/build", handler.BuildAllHandler).Methods("GET")
	r.HandleFunc("/build/{id}", handler.BuildByIdHandler).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	r.Use(handlers.CORS(headersOk, originsOk, methodsOk))

	return r
}
