package controllers

import (
	"database/sql"
	"io"
	"time"

	"github.com/goadesign/goa"
	"github.com/kawabet/goa-react-ts/backend/app"
	"github.com/kawabet/goa-react-ts/backend/models"
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
	res := app.RoomCollection{}
	rooms, err := models.AllRooms(c.db, 100) // とりあえず１００件固定
	if err != nil {
		return err
	}
	for _, room := range rooms {
		res = append(res, ToRoomMedia(room))
	}
	return ctx.OK(res)
}

// Post runs the post action.
func (c *RoomController) Post(ctx *app.PostRoomContext) error {
	room := models.Room{
		Name:        ctx.Payload.Name,
		Description: ctx.Payload.Description,
		Created:     time.Now(),
	}
	err := room.Insert(c.db)
	if err != nil {
		return err
	}
	return ctx.Created()
}

// Show runs the show action.
func (c *RoomController) Show(ctx *app.ShowRoomContext) error {
	room, err := models.RoomByID(c.db, ctx.RoomID)
	if err != nil {
		return err
	}
	if room == nil {
		return ctx.NotFound()
	}
	res := ToRoomMedia(room)
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

// ToRoomMedia はmodelsの構造体をgoaの構造体に変換します
func ToRoomMedia(room *models.Room) *app.Room {
	ret := app.Room{
		ID:          &room.ID,
		Description: room.Description,
		Name:        room.Name,
		Created:     &room.Created,
	}
	return &ret
}
