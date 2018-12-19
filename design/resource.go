package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// JWT Components
var JWT = JWTSecurity("jwt", func() {
	Header("Authorizaion")
	Scope("api:access", "API Access")
})

var _ = Resource("jwt", func() {
	BasePath("/jwt")

	Action("signin", func() {
		Description("ID/Passペアで認証を行う")
		Routing(POST("/signin"))

		Payload(func() {
			Member("userScreenName")
			Member("password")
		})

		Response(NoContent, func() {
			Headers(func() {
				Header("Authorization", String, "Generated JWT")
			})
		})
		Response(Forbidden)
	})
})

var _ = Resource("accounts", func() {
	BasePath("/accounts")

	Action("createAccount", func() {
		Description("アカウントの作成を行う")
		Routing(POST(""))

		Payload(func() {
			Member("screenName")
			Member("password")
		})

		Response(OK)
		Response(BadRequest)
	})
})

var _ = Resource("devices", func() {
	BasePath("/devices")

	Action("deviceList", func() {
		Description("デバイス一覧を返す")
		Routing(GET(""))
		Security(JWT, func() {
			Scope("api:access")
		})

		Response(OK, ArrayOf(String))
	})

	Action("generateLinkCode", func() {
		Description("ユニークコードを使用してリンクコードを生成する")
		Routing(GET("/:uniqueCode"))

		Params(func() {
			Param("uniqueCode", String, "Device unique code")
		})

		Response(OK)
	})

	Action("linkDevice", func() {
		Description("リンクコードを使用してユニークコードとアカウントを鎖付けする")
		Routing(POST("/link/:linkCode"))
		Security(JWT, func() {
			Scope("api:access")
		})

		Params(func() {
			Param("linkCode", String, "LinkCode")
		})

		Response(OK)
		Response(Forbidden)
	})
})

var _ = Resource("status", func() {
	BasePath("/status")

	Action("get", func() {
		Description("ユニークコードを基にデバイスが本来あるべき状態のデータを取得する")
		Routing(GET("/:uniqueCode"))

		Params(func() {
			Param("uniqueCode", String, "Device unique code")
		})

		Response(OK)
		Response(NotFound)
	})

	Action("settings", func() {
		Description("デバイスの設定を定義する")
		Routing(POST("/:uniqueCode/settings"))
		Security(JWT, func() {
			Scope("api:access")
		})

		Params(func() {
			Param("uniqueCode", String, "Device unique code")
		})
		Payload(func() {
			Member("userScreenName")
			Member("password")
			Member("continueTimes")
			Member("lines")
		})

		Response(OK, SettingsMediaType)
		Response(NotFound)
	})
})

var _ = Resource("resources", func() {
	BasePath("/resources")

	Security(JWT, func() {
		Scope("api:access")
	})

	Action("supportLine", func() {
		Description("対応している路線とその固有IDを返す")
		Routing(GET("/supports/lines"))

		Response(OK, ArrayOf(LineInformationMediaType))
	})
})
