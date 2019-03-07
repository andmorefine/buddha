// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// stats client HTTP transport
//
// Command:
// $ goa gen github.com/tonouchi510/goa2-sample/design

package client

import (
	"context"
	"net/http"

	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// Client lists the stats service endpoint HTTP clients.
type Client struct {
	// UserNumber Doer is the HTTP client used to make requests to the user_number
	// endpoint.
	UserNumberDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the stats service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		UserNumberDoer:      doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// UserNumber returns an endpoint that makes HTTP requests to the stats service
// user_number server.
func (c *Client) UserNumber() goa.Endpoint {
	var (
		decodeResponse = DecodeUserNumberResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUserNumberRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UserNumberDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("stats", "user_number", err)
		}
		return decodeResponse(resp)
	}
}