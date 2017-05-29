// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "Chat API": Application User Types
//
// Command:
// $ goagen
// --design=github.com/m0a-mystudy/goa-chat/design
// --out=$(GOPATH)/src/github.com/m0a-mystudy/goa-chat
// --version=v1.2.0-dirty

package app

import (
	"github.com/goadesign/goa"
	"time"
	"unicode/utf8"
)

// messagePayload user type.
type messagePayload struct {
	Body         *string    `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty"`
	GoogleUserID *string    `form:"googleUserID,omitempty" json:"googleUserID,omitempty" xml:"googleUserID,omitempty"`
	PostDate     *time.Time `form:"postDate,omitempty" json:"postDate,omitempty" xml:"postDate,omitempty"`
}

// Finalize sets the default values for messagePayload type instance.
func (ut *messagePayload) Finalize() {
	var defaultPostDate, _ = time.Parse(time.RFC3339, "1978-06-30T10:00:00+09:00")
	if ut.PostDate == nil {
		ut.PostDate = &defaultPostDate
	}
}

// Validate validates the messagePayload type instance.
func (ut *messagePayload) Validate() (err error) {
	if ut.Body == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "body"))
	}
	if ut.PostDate == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "postDate"))
	}
	if ut.Body != nil {
		if utf8.RuneCountInString(*ut.Body) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.body`, *ut.Body, utf8.RuneCountInString(*ut.Body), 1, true))
		}
	}
	if ut.Body != nil {
		if utf8.RuneCountInString(*ut.Body) > 400 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.body`, *ut.Body, utf8.RuneCountInString(*ut.Body), 400, false))
		}
	}
	return
}

// Publicize creates MessagePayload from messagePayload
func (ut *messagePayload) Publicize() *MessagePayload {
	var pub MessagePayload
	if ut.Body != nil {
		pub.Body = *ut.Body
	}
	if ut.GoogleUserID != nil {
		pub.GoogleUserID = ut.GoogleUserID
	}
	if ut.PostDate != nil {
		pub.PostDate = *ut.PostDate
	}
	return &pub
}

// MessagePayload user type.
type MessagePayload struct {
	Body         string    `form:"body" json:"body" xml:"body"`
	GoogleUserID *string   `form:"googleUserID,omitempty" json:"googleUserID,omitempty" xml:"googleUserID,omitempty"`
	PostDate     time.Time `form:"postDate" json:"postDate" xml:"postDate"`
}

// Validate validates the MessagePayload type instance.
func (ut *MessagePayload) Validate() (err error) {
	if ut.Body == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "body"))
	}

	if utf8.RuneCountInString(ut.Body) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.body`, ut.Body, utf8.RuneCountInString(ut.Body), 1, true))
	}
	if utf8.RuneCountInString(ut.Body) > 400 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.body`, ut.Body, utf8.RuneCountInString(ut.Body), 400, false))
	}
	return
}

// roomPayload user type.
type roomPayload struct {
	// Date of creation
	Created *time.Time `form:"created,omitempty" json:"created,omitempty" xml:"created,omitempty"`
	// description of room
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// ID of room
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of room
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the roomPayload type instance.
func (ut *roomPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Description == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if ut.Description != nil {
		if utf8.RuneCountInString(*ut.Description) > 400 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.description`, *ut.Description, utf8.RuneCountInString(*ut.Description), 400, false))
		}
	}
	return
}

// Publicize creates RoomPayload from roomPayload
func (ut *roomPayload) Publicize() *RoomPayload {
	var pub RoomPayload
	if ut.Created != nil {
		pub.Created = ut.Created
	}
	if ut.Description != nil {
		pub.Description = *ut.Description
	}
	if ut.ID != nil {
		pub.ID = ut.ID
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	return &pub
}

// RoomPayload user type.
type RoomPayload struct {
	// Date of creation
	Created *time.Time `form:"created,omitempty" json:"created,omitempty" xml:"created,omitempty"`
	// description of room
	Description string `form:"description" json:"description" xml:"description"`
	// ID of room
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of room
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the RoomPayload type instance.
func (ut *RoomPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if utf8.RuneCountInString(ut.Description) > 400 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.description`, ut.Description, utf8.RuneCountInString(ut.Description), 400, false))
	}
	return
}
