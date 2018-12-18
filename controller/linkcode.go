package controller

import (
	"github.com/Kumi-comb/Daikon-BE/app"
	"github.com/Kumi-comb/Daikon-BE/libs"
	"github.com/goadesign/goa"
)

// LinkcodeController implements the linkcode resource.
type LinkcodeController struct {
	*goa.Controller
}

// NewLinkcodeController creates a linkcode controller.
func NewLinkcodeController(service *goa.Service) *LinkcodeController {
	return &LinkcodeController{Controller: service.NewController("LinkcodeController")}
}

// Linkcode runs the linkcode action.
func (c *LinkcodeController) Linkcode(ctx *app.LinkcodeContext) error {
	// LinkcodeController_Linkcode: start_implement

	s := libs.GenerateLinkCode(ctx.UniqueCode)
	return ctx.OK([]byte(s))
	// LinkcodeController_Linkcode: end_implement
}
