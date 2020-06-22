package router

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	api "github.com/mleyb/divisionloadout-api/handlers"
)

// New creates a new router
func New() *mux.Router {
	r := mux.NewRouter()

	attachHandlers(r)

	setupMiddleware(r)
	
	return r
}

func attachHandlers(r *mux.Router) {
	api := api.New()

	r.HandleFunc("/build", api.BuildGetAllHandler).Methods("GET")
	r.HandleFunc("/build/{id}", api.BuildGetByIdHandler).Methods("GET")

	r.HandleFunc("/build", api.BuildCreateHandler).Methods("POST")

	r.HandleFunc("/build/{id}", api.BuildUpdateHandler).Methods("PUT")

	r.HandleFunc("/build/{id}", api.BuildDeleteHandler).Methods("DELETE")
}

func setupMiddleware(r *mux.Router) {
	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	r.Use(handlers.CORS(headersOk, originsOk, methodsOk))
}