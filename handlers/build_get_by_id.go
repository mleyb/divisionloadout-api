package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mleyb/divisionloadout-api/data"
)

// BuildGetByIdHandler returns a Build given the Id, or 404 if not found
func (handler *Handler) BuildGetByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	build, err := data.BuildGetById(id)

	if err != nil {
		respondWithError(w, 500, err.Error())
	}

	if build == nil {
		respondWithNotFound(w)
	}

	respondWithJSON(w, 200, build)
}
