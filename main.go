//go:generate goagen bootstrap -d github.com/kawabet/goa-react-ts/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/kawabet/goa-react-ts/controllers/app"
)

func main() {
	// Create service
	service := goa.New("Chat API")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "message" controller
	c := NewMessageController(service)
	app.MountMessageController(service, c)
	// Mount "room" controller
	c2 := NewRoomController(service)
	app.MountRoomController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
