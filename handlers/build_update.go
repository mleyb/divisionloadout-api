package handlers

import (
	"net/http"

	"github.com/mleyb/divisionloadout-api/data"
	"github.com/mleyb/divisionloadout-api/model"
)

func (handler *Handler) BuildUpdateHandler(w http.ResponseWriter, r *http.Request) {
	build := &model.Build{}

	_, err := data.Update(build)
	if err != nil {
		respondWithError(w, 500, err.Error())
	}

	respondWithOk(w)
}
