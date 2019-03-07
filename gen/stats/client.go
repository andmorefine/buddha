// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// stats client
//
// Command:
// $ goa gen github.com/tonouchi510/goa2-sample/design

package stats

import (
	"context"

	goa "goa.design/goa"
)

// Client is the "stats" service client.
type Client struct {
	UserNumberEndpoint goa.Endpoint
}

// NewClient initializes a "stats" service client given the endpoints.
func NewClient(userNumber goa.Endpoint) *Client {
	return &Client{
		UserNumberEndpoint: userNumber,
	}
}

// UserNumber calls the "user_number" endpoint of the "stats" service.
func (c *Client) UserNumber(ctx context.Context) (res *Statsuser, err error) {
	var ires interface{}
	ires, err = c.UserNumberEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*Statsuser), nil
}