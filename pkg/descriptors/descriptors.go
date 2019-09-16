// +nirvana:api=descriptors:"Descriptor"

package descriptors

import (
	v1 "github.com/caicloud/nirvana-template-project/pkg/descriptors/v1"
	"github.com/caicloud/nirvana-template-project/pkg/middlewares"

	def "github.com/caicloud/nirvana/definition"
)

// Descriptor returns a combined descriptor for APIs of all versions.
func Descriptor() def.Descriptor {
	return def.Descriptor{
		Consumes:    []string{def.MIMEJSON},
		Produces:    []string{def.MIMEJSON},
		Description: "APIs",
		Path:        "/apis",
		Middlewares: middlewares.Middlewares(),
		Children: []def.Descriptor{
			v1.Descriptor(),
		},
	}
}
