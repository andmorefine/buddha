// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// secured HTTP server types
//
// Command:
// $ goa gen github.com/andmorefine/goa2-sample/design

package server

import (
	secured "github.com/andmorefine/goa2-sample/gen/secured"
	securedviews "github.com/andmorefine/goa2-sample/gen/secured/views"
)

// SigninResponseBody is the type of the "secured" service "signin" endpoint
// HTTP response body.
type SigninResponseBody struct {
	// New Jwt token
	Authorization string `form:"Authorization" json:"Authorization" xml:"Authorization"`
}

// SigninUnauthorizedResponseBody is the type of the "secured" service "signin"
// endpoint HTTP response body for the "unauthorized" error.
type SigninUnauthorizedResponseBody string

// NewSigninResponseBody builds the HTTP response body from the result of the
// "signin" endpoint of the "secured" service.
func NewSigninResponseBody(res *securedviews.GoaJWTView) *SigninResponseBody {
	body := &SigninResponseBody{
		Authorization: *res.Authorization,
	}
	return body
}

// NewSigninUnauthorizedResponseBody builds the HTTP response body from the
// result of the "signin" endpoint of the "secured" service.
func NewSigninUnauthorizedResponseBody(res secured.Unauthorized) SigninUnauthorizedResponseBody {
	body := SigninUnauthorizedResponseBody(res)
	return body
}

// NewSigninPayload builds a secured service signin endpoint payload.
func NewSigninPayload() *secured.SigninPayload {
	return &secured.SigninPayload{}
}
