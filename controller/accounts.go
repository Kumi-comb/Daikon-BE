package controller

import (
	"github.com/Kumi-comb/Daikon-BE/app"
	"github.com/goadesign/goa"
)

// AccountsController implements the accounts resource.
type AccountsController struct {
	*goa.Controller
}

// NewAccountsController creates a accounts controller.
func NewAccountsController(service *goa.Service) *AccountsController {
	return &AccountsController{Controller: service.NewController("AccountsController")}
}

// CreateAccount runs the createAccount action.
func (c *AccountsController) CreateAccount(ctx *app.CreateAccountAccountsContext) error {
	// AccountsController_CreateAccount: start_implement

	// Put your logic here

	return nil
	// AccountsController_CreateAccount: end_implement
}
