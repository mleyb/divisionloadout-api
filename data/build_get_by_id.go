package data

import (
	"github.com/mleyb/divisionloadout-api/model"
)

// BuildGetById returns a Build given its Id
func BuildGetById(id string) (*model.Build, error) {
	// TODO
	return &model.Build{
		Buildid:   "1",
		Buildname: "Example",
	}, nil
}
