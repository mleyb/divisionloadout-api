package handlers

import (
	"net/http"
)

// BuildAllHandler returns all builds
func (handler *Handler) BuildAllHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	w.Write([]byte("TODO"))
}
