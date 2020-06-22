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

	r.HandleFunc("/build", handler.BuildCreateHandler).Methods("POST")

	r.HandleFunc("/build", handler.BuildUpdateHandler).Methods("PUT")

	r.HandleFunc("/build/{id}", handler.BuildDeleteHandler).Methods("DELETE")

	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	r.Use(handlers.CORS(headersOk, originsOk, methodsOk))

	return r
}
