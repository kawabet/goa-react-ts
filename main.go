//go:generate goagen bootstrap -d github.com/kawabet/goa-react-ts/design

package main

import (
	"database/sql"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Create service
	service := goa.New("Chat API")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	db, err := sql.Open("mysql", "root:root@/dev?parseTime=true")
	if err != nil {
		service.LogError("startup", "err", err)
	}

	// Mount "message" controller
	c := NewMessageController(service, db)
	app.MountMessageController(service, c)
	// Mount "room" controller
	c2 := NewRoomController(service, db)
	app.MountRoomController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
