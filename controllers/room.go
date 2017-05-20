package controllers

import (
	"database/sql"
	"io"

	"github.com/goadesign/goa"
	"github.com/kawabet/goa-react-ts/controllers/app"
	"golang.org/x/net/websocket"
)

// RoomController implements the room resource.
type RoomController struct {
	*goa.Controller
	db *sql.DB
}

// NewRoomController creates a room controller.
func NewRoomController(service *goa.Service, db *sql.DB) *RoomController {
	return &RoomController{
		Controller: service.NewController("RoomController"),
		db:         db,
	}
}

// List runs the list action.
func (c *RoomController) List(ctx *app.ListRoomContext) error {
	// RoomController_List: start_implement

	// Put your logic here

	// RoomController_List: end_implement
	res := app.RoomCollection{}
	return ctx.OK(res)
}

// Post runs the post action.
func (c *RoomController) Post(ctx *app.PostRoomContext) error {
	// RoomController_Post: start_implement

	// Put your logic here

	// RoomController_Post: end_implement
	return nil
}

// Show runs the show action.
func (c *RoomController) Show(ctx *app.ShowRoomContext) error {
	// RoomController_Show: start_implement

	// Put your logic here

	// RoomController_Show: end_implement
	res := &app.Room{}
	return ctx.OK(res)
}

// Watch runs the watch action.
func (c *RoomController) Watch(ctx *app.WatchRoomContext) error {
	c.WatchWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// WatchWSHandler establishes a websocket connection to run the watch action.
func (c *RoomController) WatchWSHandler(ctx *app.WatchRoomContext) websocket.Handler {
	return func(ws *websocket.Conn) {
		// RoomController_Watch: start_implement

		// Put your logic here

		// RoomController_Watch: end_implement
		ws.Write([]byte("watch room"))
		// Dummy echo websocket server
		io.Copy(ws, ws)
	}
}
