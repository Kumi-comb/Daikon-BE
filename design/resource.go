package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("linkcode", func() {
	BasePath("/linkcode")

	Action("create", func() {
		Description("リンクコード生成")
		Routing(GET("/:uniqueCode"))
		Params(func() {
			Param("uniqueCode", String, "Widget unique code")
		})

		Response(OK, "text/plain")
	})
})

var _ = Resource("status", func() {
	BasePath("/status")

	Action("get", func() {
		Description("ステータス取得")
		Routing(GET("/:uniqueCode"))
		Params(func() {
			Param("uniqueCode", String, "Widget unique code")
		})

		Response(OK, "text/plain")
	})
})
