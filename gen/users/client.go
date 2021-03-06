// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// users client
//
// Command:
// $ goa gen github.com/andmorefine/buddha/design

package users

import (
	"context"

	goa "goa.design/goa"
)

// Client is the "users" service client.
type Client struct {
	ListEndpoint   goa.Endpoint
	ShowEndpoint   goa.Endpoint
	AddEndpoint    goa.Endpoint
	UpdateEndpoint goa.Endpoint
	RemoveEndpoint goa.Endpoint
}

// NewClient initializes a "users" service client given the endpoints.
func NewClient(list, show, add, update, remove goa.Endpoint) *Client {
	return &Client{
		ListEndpoint:   list,
		ShowEndpoint:   show,
		AddEndpoint:    add,
		UpdateEndpoint: update,
		RemoveEndpoint: remove,
	}
}

// List calls the "list" endpoint of the "users" service.
func (c *Client) List(ctx context.Context) (res StoredUserCollection, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(StoredUserCollection), nil
}

// Show calls the "show" endpoint of the "users" service.
// Show may return the following errors:
//	- "not_found" (type *NotFound): user not found
//	- error: internal error
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *StoredUser, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StoredUser), nil
}

// Add calls the "add" endpoint of the "users" service.
func (c *Client) Add(ctx context.Context, p *User) (res int64, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int64), nil
}

// Update calls the "update" endpoint of the "users" service.
func (c *Client) Update(ctx context.Context, p *UpdatePayload) (res *StoredUser, err error) {
	var ires interface{}
	ires, err = c.UpdateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StoredUser), nil
}

// Remove calls the "remove" endpoint of the "users" service.
// Remove may return the following errors:
//	- "not_found" (type *NotFound): user not found
//	- error: internal error
func (c *Client) Remove(ctx context.Context, p *RemovePayload) (err error) {
	_, err = c.RemoveEndpoint(ctx, p)
	return
}
