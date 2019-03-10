package design

import (
	. "goa.design/goa/dsl"
	cors "goa.design/plugins/cors/dsl"
)

// API describes the global properties of the API server.
var _ = API("buddha", func() {
	Title("SampleAPI")
	Description("goa2 sample code.")
	Version("1.0")
	Contact(func() {
		Name("andmorefine")
		Email("andmorefine@gmail.com")
		URL("https://github.com/andmorefine/buddha/issues")
	})
	Docs(func() {
		Description("wiki")
		URL("https://github.com/andmorefine/buddha/wiki")
	})

	cors.Origin("/.*localhost.*/", func() {
		cors.Headers("authorization, content-type")
		cors.Methods("GET", "POST", "PUT", "DELETE", "OPTIONS")
		cors.MaxAge(600)
	})

	Server("buddha", func() {
		Services("users", "swagger", "secured", "viron", "stats")
		Host("localhost", func() {
			Description("development host")
			URI("http://localhost:8000")
		})
	})
})
