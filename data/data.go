package data

import (
	"github.com/mleyb/divisionloadout-api/model"
)

// BuildById returns a Build given its Id
func BuildById(id string) (*model.Build, error) {
	// TODO
	return &model.Build{
		Buildid:   "1",
		Buildname: "Example",
	}, nil
}
