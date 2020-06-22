package data

import (
	"github.com/mleyb/divisionloadout-api/model"
)

// Create creates a new build
func Create(build *model.Build) (*model.Build, error) {
	// TODO
	build.Buildid = "ABC123"
	
	return build, nil
}
