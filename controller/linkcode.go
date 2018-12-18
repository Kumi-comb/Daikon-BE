package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/Kumi-comb/Daikon-BE/app"
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
func (c *LinkcodeController) Linkcode(ctx *app.LinkcodeLinkcodeContext) error {
	// LinkcodeController_Linkcode: start_implement

	// Put your logic here
	b := sha256.Sum256([]byte(ctx.UniqueCode + time.Now().Format("20060102150405")))
	s := hex.EncodeToString(b[:])[:4]
	fmt.Printf(time.Now().Format("200601021510405"))

	return ctx.OK([]byte(s))
	// LinkcodeController_Linkcode: end_implement
}
