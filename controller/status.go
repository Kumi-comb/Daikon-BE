package controller

import (
	"github.com/Kumi-comb/Daikon-BE/app"
	"github.com/goadesign/goa"
)

// StatusController implements the status resource.
type StatusController struct {
	*goa.Controller
}

// NewStatusController creates a status controller.
func NewStatusController(service *goa.Service) *StatusController {
	return &StatusController{Controller: service.NewController("StatusController")}
}

// Get runs the get action.
func (c *StatusController) Get(ctx *app.GetStatusContext) error {
	// StatusController_Get: start_implement

	return ctx.OK([]byte("2F"))
	// StatusController_Get: end_implement
}
