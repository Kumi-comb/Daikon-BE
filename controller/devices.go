package controller

import (
	"github.com/Kumi-comb/Daikon-BE/app"
	"github.com/Kumi-comb/Daikon-BE/libs"
	"github.com/goadesign/goa"
)

// DevicesController implements the devices resource.
type DevicesController struct {
	*goa.Controller
}

// NewDevicesController creates a devices controller.
func NewDevicesController(service *goa.Service) *DevicesController {
	return &DevicesController{Controller: service.NewController("DevicesController")}
}

// DeviceList runs the deviceList action.
func (c *DevicesController) DeviceList(ctx *app.DeviceListDevicesContext) error {
	// DevicesController_DeviceList: start_implement

	// fmt.Printf("aaa")

	// // Put your logic here
	// // Retrieve the token claims
	// token := jwt.ContextJWT(ctx)
	// if token == nil {
	// 	return fmt.Errorf("JWT token is missing from context") // internal error
	// }
	// claims := token.Claims.(jwtgo.MapClaims)

	// // Use the claims to authorize
	// subject := claims["sub"].(string)

	// if subject == "AccessToken" {
	// 	// A real app would probably use an "Unauthorized" response here
	// 	// return ctx.Unauthorized()
	// }

	// deviceList := []string{
	// 	"a",
	// }
	return nil
	// DevicesController_DeviceList: end_implement
}

// GenerateLinkCode runs the generateLinkCode action.
func (c *DevicesController) GenerateLinkCode(ctx *app.GenerateLinkCodeDevicesContext) error {
	// DevicesController_GenerateLinkCode: start_implement

	// Put your logic here
	l := libs.CreateLinkCode(ctx.UniqueCode)
	return ctx.OK([]byte(l))
	// DevicesController_GenerateLinkCode: end_implement
}

// LinkDevice runs the linkDevice action.
func (c *DevicesController) LinkDevice(ctx *app.LinkDeviceDevicesContext) error {
	// DevicesController_LinkDevice: start_implement

	// Put your logic here
	u := libs.GetUniqueCode(ctx.LinkCode)
	return ctx.OK([]byte(u))
	// DevicesController_LinkDevice: end_implement
}
