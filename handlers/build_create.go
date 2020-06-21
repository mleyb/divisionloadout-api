package handlers

import (
	"net/http"

	"github.com/mleyb/divisionloadout-api/data"
	"github.com/mleyb/divisionloadout-api/model"
)

func (handler *Handler) BuildCreateHandler(w http.ResponseWriter, r *http.Request) {
	build := &model.Build{}
	
	data.Create(build);

	respondWithOk(w)
}