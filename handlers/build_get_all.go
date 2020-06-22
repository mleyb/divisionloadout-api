package handlers

import (
	"net/http"

	"github.com/mleyb/divisionloadout-api/data"
)

// BuildGetAllHandler returns all builds
func (handler *Handler) BuildGetAllHandler(w http.ResponseWriter, r *http.Request) {

	builds, err := data.BuildGetAll()

	if err != nil {
		respondWithError(w, 500, err.Error())
	}

	respondWithJSON(w, 200, builds)
}
