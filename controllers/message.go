package main

import (
	"github.com/goadesign/goa"
	"github.com/kawabet/goa-react-ts/controllers/app"
)

// MessageController implements the message resource.
type MessageController struct {
	*goa.Controller
}

// NewMessageController creates a message controller.
func NewMessageController(service *goa.Service) *MessageController {
	return &MessageController{Controller: service.NewController("MessageController")}
}

// List runs the list action.
func (c *MessageController) List(ctx *app.ListMessageContext) error {
	// MessageController_List: start_implement

	// Put your logic here

	// MessageController_List: end_implement
	res := app.MessageCollection{}
	return ctx.OK(res)
}

// Post runs the post action.
func (c *MessageController) Post(ctx *app.PostMessageContext) error {
	// MessageController_Post: start_implement

	// Put your logic here

	// MessageController_Post: end_implement
	return nil
}

// Show runs the show action.
func (c *MessageController) Show(ctx *app.ShowMessageContext) error {
	// MessageController_Show: start_implement

	// Put your logic here

	// MessageController_Show: end_implement
	res := &app.Message{}
	return ctx.OK(res)
}
