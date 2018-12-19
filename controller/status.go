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

	// Put your logic here

	return nil
	// StatusController_Get: end_implement
}

// Settings runs the settings action.
func (c *StatusController) Settings(ctx *app.SettingsStatusContext) error {
	// StatusController_Settings: start_implement

	// Put your logic here

	res := &app.Settings{}
	return ctx.OK(res)
	// StatusController_Settings: end_implement
}
