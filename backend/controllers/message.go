package controllers

import (
	"database/sql"
	"time"

	"github.com/goadesign/goa"
	"github.com/kawabet/goa-react-ts/backend/app"
	"github.com/kawabet/goa-react-ts/backend/models"
)

// MessageController implements the message resource.
type MessageController struct {
	*goa.Controller
	db *sql.DB
}

// NewMessageController creates a message controller.
func NewMessageController(service *goa.Service, db *sql.DB) *MessageController {
	return &MessageController{
		Controller: service.NewController("MessageController"),
		db:         db,
	}
}

// List runs the list action.
func (c *MessageController) List(ctx *app.ListMessageContext) error {
	res := app.MessageCollection{}
	messages, err := models.MessagesByRoomID(c.db, ctx.RoomID)
	if err != nil {
		return err
	}
	for _, m := range messages {
		res = append(res, ToMessageMedia(m))
	}
	return ctx.OK(res)
}

// Post runs the post action.
func (c *MessageController) Post(ctx *app.PostMessageContext) error {
	m := models.Message{
		RoomID:    ctx.RoomID,
		AccountID: ctx.Payload.AccountID,
		Body:      ctx.Payload.Body,
		Postdate:  time.Now(),
	}
	err := m.Insert(c.db)
	if err != nil {
		return ctx.BadRequest()
	}
	return ctx.Created()
}

// Show runs the show action.
func (c *MessageController) Show(ctx *app.ShowMessageContext) error {
	// MessageController_Show: start_implement

	// Put your logic here

	// MessageController_Show: end_implement
	res := &app.Message{}
	return ctx.OK(res)
}

// ToMessageMedia はmodelsの構造体をgoaの構造体に変換します
func ToMessageMedia(message *models.Message) *app.Message {
	ret := app.Message{
		AccountID: message.AccountID,
		Body:      message.Body,
		PostDate:  message.Postdate,
	}
	return &ret
}
