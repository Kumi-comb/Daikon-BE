package controller

import (
	"github.com/Kumi-comb/Daikon-BE/app"
	"github.com/goadesign/goa"
)

// ResourcesController implements the resources resource.
type ResourcesController struct {
	*goa.Controller
}

// NewResourcesController creates a resources controller.
func NewResourcesController(service *goa.Service) *ResourcesController {
	return &ResourcesController{Controller: service.NewController("ResourcesController")}
}

// SupportLine runs the supportLine action.
func (c *ResourcesController) SupportLine(ctx *app.SupportLineResourcesContext) error {
	// ResourcesController_SupportLine: start_implement

	// Put your logic here

	return nil
	// ResourcesController_SupportLine: end_implement
}
