package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mleyb/divisionloadout-api/data"
)

// DeleteHandler deletes a build
func (handler *Handler) BuildDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	err := data.Delete(id)
	if err != nil {
		respondWithError(w, 500, err.Error())
	}

	respondWithOk(w)
}
