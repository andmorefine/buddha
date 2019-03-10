// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// secured client
//
// Command:
// $ goa gen github.com/andmorefine/buddha/design

package secured

import (
	"context"

	goa "goa.design/goa"
)

// Client is the "secured" service client.
type Client struct {
	SigninEndpoint goa.Endpoint
}

// NewClient initializes a "secured" service client given the endpoints.
func NewClient(signin goa.Endpoint) *Client {
	return &Client{
		SigninEndpoint: signin,
	}
}

// Signin calls the "signin" endpoint of the "secured" service.
// Signin may return the following errors:
//	- "unauthorized" (type Unauthorized)
//	- error: internal error
func (c *Client) Signin(ctx context.Context, p *SigninPayload) (res *GoaJWT, err error) {
	var ires interface{}
	ires, err = c.SigninEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GoaJWT), nil
}
