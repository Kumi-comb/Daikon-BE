package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// LineInformationMediaType is a one of support lines.
var LineInformationMediaType = MediaType("application/vnd.line.information", func() {
	attributes := func() {
		Attribute("ID")
		Attribute("CorporationName")
		Attribute("LineName")
	}
	Attributes(attributes)
	View("default", attributes)
})

// SettingsMediaType is a setting.
var SettingsMediaType = MediaType("appliaction/vnd.settings", func() {
	attributes := func() {
		Attribute("lines", ArrayOf(String))
		Attribute("times", ArrayOf(HashOf(String, String)))
	}
	Attributes(attributes)
	View("default", attributes)
})
