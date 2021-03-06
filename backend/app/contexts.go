// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "Chat API": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/m0a-mystudy/goa-chat/design
// --out=$(GOPATH)/src/github.com/m0a-mystudy/goa-chat
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"strconv"
)

// ListAccountContext provides the account list action context.
type ListAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller list action.
func NewListAccountContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListAccountContext) OK(r AccountCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.account+json; type=collection")
	if r == nil {
		r = AccountCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListAccountContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// PostAccountContext provides the account post action context.
type PostAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *MessagePayload
}

// NewPostAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller post action.
func NewPostAccountContext(ctx context.Context, r *http.Request, service *goa.Service) (*PostAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := PostAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *PostAccountContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *PostAccountContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// ShowAccountContext provides the account show action context.
type ShowAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	User string
}

// NewShowAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller show action.
func NewShowAccountContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUser := req.Params["user"]
	if len(paramUser) > 0 {
		rawUser := paramUser[0]
		rctx.User = rawUser
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowAccountContext) OK(r *Account) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.account+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ShowAccountContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowAccountContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListMessageContext provides the message list action context.
type ListMessageContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Limit  *int
	Offset *int
	RoomID int
}

// NewListMessageContext parses the incoming request URL and body, performs validations and creates the
// context used by the message controller list action.
func NewListMessageContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListMessageContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListMessageContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramLimit := req.Params["limit"]
	if len(paramLimit) > 0 {
		rawLimit := paramLimit[0]
		if limit, err2 := strconv.Atoi(rawLimit); err2 == nil {
			tmp2 := limit
			tmp1 := &tmp2
			rctx.Limit = tmp1
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("limit", rawLimit, "integer"))
		}
	}
	paramOffset := req.Params["offset"]
	if len(paramOffset) > 0 {
		rawOffset := paramOffset[0]
		if offset, err2 := strconv.Atoi(rawOffset); err2 == nil {
			tmp4 := offset
			tmp3 := &tmp4
			rctx.Offset = tmp3
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("offset", rawOffset, "integer"))
		}
	}
	paramRoomID := req.Params["roomID"]
	if len(paramRoomID) > 0 {
		rawRoomID := paramRoomID[0]
		if roomID, err2 := strconv.Atoi(rawRoomID); err2 == nil {
			rctx.RoomID = roomID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("roomID", rawRoomID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListMessageContext) OK(r MessageWithAccountCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.message_with_account+json; type=collection")
	if r == nil {
		r = MessageWithAccountCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListMessageContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// PostMessageContext provides the message post action context.
type PostMessageContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	RoomID  int
	Payload *MessagePayload
}

// NewPostMessageContext parses the incoming request URL and body, performs validations and creates the
// context used by the message controller post action.
func NewPostMessageContext(ctx context.Context, r *http.Request, service *goa.Service) (*PostMessageContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := PostMessageContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramRoomID := req.Params["roomID"]
	if len(paramRoomID) > 0 {
		rawRoomID := paramRoomID[0]
		if roomID, err2 := strconv.Atoi(rawRoomID); err2 == nil {
			rctx.RoomID = roomID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("roomID", rawRoomID, "integer"))
		}
	}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *PostMessageContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *PostMessageContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// ShowMessageContext provides the message show action context.
type ShowMessageContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	MessageID int
	RoomID    int
}

// NewShowMessageContext parses the incoming request URL and body, performs validations and creates the
// context used by the message controller show action.
func NewShowMessageContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowMessageContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowMessageContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramMessageID := req.Params["messageID"]
	if len(paramMessageID) > 0 {
		rawMessageID := paramMessageID[0]
		if messageID, err2 := strconv.Atoi(rawMessageID); err2 == nil {
			rctx.MessageID = messageID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("messageID", rawMessageID, "integer"))
		}
	}
	paramRoomID := req.Params["roomID"]
	if len(paramRoomID) > 0 {
		rawRoomID := paramRoomID[0]
		if roomID, err2 := strconv.Atoi(rawRoomID); err2 == nil {
			rctx.RoomID = roomID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("roomID", rawRoomID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowMessageContext) OK(r *Message) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.message+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ShowMessageContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowMessageContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListRoomContext provides the room list action context.
type ListRoomContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Limit  *int
	Offset *int
}

// NewListRoomContext parses the incoming request URL and body, performs validations and creates the
// context used by the room controller list action.
func NewListRoomContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListRoomContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListRoomContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramLimit := req.Params["limit"]
	if len(paramLimit) > 0 {
		rawLimit := paramLimit[0]
		if limit, err2 := strconv.Atoi(rawLimit); err2 == nil {
			tmp10 := limit
			tmp9 := &tmp10
			rctx.Limit = tmp9
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("limit", rawLimit, "integer"))
		}
	}
	paramOffset := req.Params["offset"]
	if len(paramOffset) > 0 {
		rawOffset := paramOffset[0]
		if offset, err2 := strconv.Atoi(rawOffset); err2 == nil {
			tmp12 := offset
			tmp11 := &tmp12
			rctx.Offset = tmp11
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("offset", rawOffset, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListRoomContext) OK(r RoomCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.room+json; type=collection")
	if r == nil {
		r = RoomCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListRoomContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// PostRoomContext provides the room post action context.
type PostRoomContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *RoomPayload
}

// NewPostRoomContext parses the incoming request URL and body, performs validations and creates the
// context used by the room controller post action.
func NewPostRoomContext(ctx context.Context, r *http.Request, service *goa.Service) (*PostRoomContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := PostRoomContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *PostRoomContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *PostRoomContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// ShowRoomContext provides the room show action context.
type ShowRoomContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	RoomID int
}

// NewShowRoomContext parses the incoming request URL and body, performs validations and creates the
// context used by the room controller show action.
func NewShowRoomContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowRoomContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowRoomContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramRoomID := req.Params["roomID"]
	if len(paramRoomID) > 0 {
		rawRoomID := paramRoomID[0]
		if roomID, err2 := strconv.Atoi(rawRoomID); err2 == nil {
			rctx.RoomID = roomID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("roomID", rawRoomID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowRoomContext) OK(r *Room) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.room+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ShowRoomContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowRoomContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// WatchRoomContext provides the room watch action context.
type WatchRoomContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	RoomID int
}

// NewWatchRoomContext parses the incoming request URL and body, performs validations and creates the
// context used by the room controller watch action.
func NewWatchRoomContext(ctx context.Context, r *http.Request, service *goa.Service) (*WatchRoomContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := WatchRoomContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramRoomID := req.Params["roomID"]
	if len(paramRoomID) > 0 {
		rawRoomID := paramRoomID[0]
		if roomID, err2 := strconv.Atoi(rawRoomID); err2 == nil {
			rctx.RoomID = roomID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("roomID", rawRoomID, "integer"))
		}
	}
	return &rctx, err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *WatchRoomContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}
