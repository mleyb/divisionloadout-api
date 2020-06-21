package handlers

import (
	"encoding/json"
	"net/http"
)

// Handler defines a requst handler
type Handler struct {
}

// New creates a new request handler
func New() *Handler {
	handler := &Handler{}
	return handler
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

func respondWithNotFound(w http.ResponseWriter) {
	w.WriteHeader(404)
}
