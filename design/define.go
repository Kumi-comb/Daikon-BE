package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Daikon-BE", func() { // API defines the microservice endpoint and
	Title("Daikon backend system")                    // other global properties. There should be one
	Description("API for train delay display widget") // and exactly one API definition appearing in
	Scheme("http")                                    // the design.
	Host("localhost:8080")
})
