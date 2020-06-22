package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

// Handler defines a requst handler
type Handler struct {
}

// New creates a new request handler
func New() *Handler {
	handler := &Handler{}
	return handler
}

func getQueryParam(r *http.Request, name string) string {
	vars := mux.Vars(r)

	value := vars[name]

	return value
}

func respondWithJSON(w http.ResponseWriter, status int, content interface{}) {
	body, err := json.Marshal(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(body)
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	respondWithJSON(w, status, map[string]string{"error": message})
}

func respondWithOk(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func respondWithBadRequest(w http.ResponseWriter) {
	w.WriteHeader(400)
}

func respondWithNotFound(w http.ResponseWriter) {
	w.WriteHeader(404)
}

func respondWithCreated(w http.ResponseWriter, id string) {
	respondWithJSON(w, 201, map[string]string{"id": id})
}
