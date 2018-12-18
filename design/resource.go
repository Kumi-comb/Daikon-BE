package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("linkcode", func() {
	BasePath("/linkcode")

	Action("linkcode", func() {
		Description("リンクコード生成")
		Routing(GET("/:uniqueCode"))
		Params(func() {
			Param("uniqueCode", String, "Widget unique code")
		})

		Response(OK, "text/plain")
	})
})
