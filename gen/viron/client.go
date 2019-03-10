// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// viron client
//
// Command:
// $ goa gen github.com/andmorefine/buddha/design

package viron

import (
	"context"

	goa "goa.design/goa"
)

// Client is the "viron" service client.
type Client struct {
	AuthtypeEndpoint  goa.Endpoint
	VironMenuEndpoint goa.Endpoint
}

// NewClient initializes a "viron" service client given the endpoints.
func NewClient(authtype, vironMenu goa.Endpoint) *Client {
	return &Client{
		AuthtypeEndpoint:  authtype,
		VironMenuEndpoint: vironMenu,
	}
}

// Authtype calls the "authtype" endpoint of the "viron" service.
func (c *Client) Authtype(ctx context.Context) (res VironAuthtypeCollection, err error) {
	var ires interface{}
	ires, err = c.AuthtypeEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(VironAuthtypeCollection), nil
}

// VironMenu calls the "viron_menu" endpoint of the "viron" service.
func (c *Client) VironMenu(ctx context.Context) (res *VironMenu, err error) {
	var ires interface{}
	ires, err = c.VironMenuEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*VironMenu), nil
}
