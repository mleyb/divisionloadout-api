package handlers

import (
	"net/http"
)

// LoadoutHandler is a handler for loadouts
func (handler *Handler) LoadoutHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
