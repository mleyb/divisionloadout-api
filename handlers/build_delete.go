package handlers

import (
	"net/http"

	"github.com/mleyb/divisionloadout-api/data"
)

// DeleteHandler deletes a build
func (handler *Handler) BuildDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := getQueryParam(r, "id")
	
	err := data.Delete(id)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusInternalServerError)
	}

	respondWithOk(w)
}
