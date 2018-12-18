//go:generate goagen bootstrap -d github.com/Kumi-comb/Daikon-BE/design

package main

import (
	"github.com/Kumi-comb/Daikon-BE/app"
	"github.com/Kumi-comb/Daikon-BE/controller"
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

	// Mount "linkcode" controller
	c := controller.NewLinkcodeController(service)
	app.MountLinkcodeController(service, c)

	// Mount "status" controller
	c := controller.NewStatusController(service)
	app.MountStatusController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
