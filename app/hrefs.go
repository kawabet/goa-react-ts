// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "Chat API": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/kawabet/goa-react-ts/design
// --out=$(GOPATH)/src/github.com/kawabet/goa-react-ts
// --version=v1.2.0-dirty

package app

import (
	"fmt"
	"strings"
)

// MessageHref returns the resource href.
func MessageHref(roomID, messageID interface{}) string {
	paramroomID := strings.TrimLeftFunc(fmt.Sprintf("%v", roomID), func(r rune) bool { return r == '/' })
	parammessageID := strings.TrimLeftFunc(fmt.Sprintf("%v", messageID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/rooms/%v/messages/%v", paramroomID, parammessageID)
}

// RoomHref returns the resource href.
func RoomHref(roomID interface{}) string {
	paramroomID := strings.TrimLeftFunc(fmt.Sprintf("%v", roomID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/rooms/%v", paramroomID)
}
