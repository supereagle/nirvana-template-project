package v1

import (
	"github.com/caicloud/nirvana-template-project/pkg/middlewares"

	def "github.com/caicloud/nirvana/definition"
)

// descriptors describe APIs of current version.
var descriptors []def.Descriptor

// register registers descriptors.
func register(ds ...def.Descriptor) {
	descriptors = append(descriptors, ds...)
}

// Descriptor returns a combined descriptor for current version.
func Descriptor() def.Descriptor {
	return def.Descriptor{
		Consumes:    []string{def.MIMEJSON},
		Produces:    []string{def.MIMEJSON},
		Description: "v1 APIs",
		Path:        "apis/v1",
		Middlewares: middlewares.Middlewares(),
		Children:    descriptors,
	}
}
