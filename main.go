//go:generate goagen bootstrap -d github.com/Kumi-comb/Daikon-BE/design

package main

import (
	"log"

	"github.com/Kumi-comb/Daikon-BE/app"
	"github.com/Kumi-comb/Daikon-BE/controller"
	"github.com/Kumi-comb/Daikon-BE/security"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("Daikon-BE")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount security middlewares (goadesign/examples/security/main.go)
	jwtMiddleware, err := security.NewJWTMiddleware()
	if err != nil {
		log.Panic(err)
	}
	app.UseJWTMiddleware(service, jwtMiddleware)

	// Mount "accounts" controller
	c := controller.NewAccountsController(service)
	app.MountAccountsController(service, c)
	// Mount "devices" controller
	c2 := controller.NewDevicesController(service)
	app.MountDevicesController(service, c2)
	// Mount "jwt" controller
	c3, err := controller.NewJWTController(service)
	if err != nil {
		log.Panic(err)
	}
	app.MountJWTController(service, c3)
	// Mount "resources" controller
	c4 := controller.NewResourcesController(service)
	app.MountResourcesController(service, c4)
	// Mount "status" controller
	c5 := controller.NewStatusController(service)
	app.MountStatusController(service, c5)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
