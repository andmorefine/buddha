package design

import (
	. "goa.design/goa/dsl"
)

var _ = Service("swagger", func() {
	Description("The swagger service serves the API swagger definition.")

	Files("/v1/swagger.json", "../../gen/http/openapi.json", func() {
		Description("JSON document containing the API swagger definition")
	})
})