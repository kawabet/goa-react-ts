// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "Chat API": account Resource Client
//
// Command:
// $ goagen
// --design=github.com/kawabet/goa-react-ts/design
// --out=$(GOPATH)/src/github.com/kawabet/goa-react-ts/backend
// --version=v1.2.0-dirty

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ListAccountPath computes a request path to the list action of account.
func ListAccountPath() string {

	return fmt.Sprintf("/api/accounts")
}

// Retrieve all accunts.
func (c *Client) ListAccount(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListAccountRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListAccountRequest create the request corresponding to the list action endpoint of the account resource.
func (c *Client) NewListAccountRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PostAccountPath computes a request path to the post action of account.
func PostAccountPath() string {

	return fmt.Sprintf("/api/accounts")
}

// Create new account
func (c *Client) PostAccount(ctx context.Context, path string, payload *MessagePayload, contentType string) (*http.Response, error) {
	req, err := c.NewPostAccountRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPostAccountRequest create the request corresponding to the post action endpoint of the account resource.
func (c *Client) NewPostAccountRequest(ctx context.Context, path string, payload *MessagePayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// ShowAccountPath computes a request path to the show action of account.
func ShowAccountPath(user string) string {
	param0 := user

	return fmt.Sprintf("/api/accounts/%s", param0)
}

// Retrieve account with given id or something
func (c *Client) ShowAccount(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowAccountRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowAccountRequest create the request corresponding to the show action endpoint of the account resource.
func (c *Client) NewShowAccountRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
