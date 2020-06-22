package handlers

import (
	"net/http"

	"github.com/mleyb/divisionloadout-api/data"
	"github.com/mleyb/divisionloadout-api/model"
)

func (handler *Handler) BuildUpdateHandler(w http.ResponseWriter, r *http.Request) {
	id := getQueryParam(r, "id")

	build := &model.Build{}

	_, err := data.Update(id, build)
	if err != nil {
		respondWithError(w, 500, err.Error())
	}

	respondWithOk(w)
}
