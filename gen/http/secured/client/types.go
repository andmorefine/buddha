// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// secured HTTP client types
//
// Command:
// $ goa gen github.com/tonouchi510/goa2-sample/design

package client

import (
	secured "github.com/tonouchi510/goa2-sample/gen/secured"
	securedviews "github.com/tonouchi510/goa2-sample/gen/secured/views"
)

// SigninResponseBody is the type of the "secured" service "signin" endpoint
// HTTP response body.
type SigninResponseBody struct {
	// New Jwt token
	Authorization *string `form:"Authorization,omitempty" json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

// SigninUnauthorizedResponseBody is the type of the "secured" service "signin"
// endpoint HTTP response body for the "unauthorized" error.
type SigninUnauthorizedResponseBody string

// NewSigninGoaJWTOK builds a "secured" service "signin" endpoint result from a
// HTTP "OK" response.
func NewSigninGoaJWTOK(body *SigninResponseBody) *securedviews.GoaJWTView {
	v := &securedviews.GoaJWTView{
		Authorization: body.Authorization,
	}
	return v
}

// NewSigninUnauthorized builds a secured service signin endpoint unauthorized
// error.
func NewSigninUnauthorized(body SigninUnauthorizedResponseBody) secured.Unauthorized {
	v := secured.Unauthorized(body)
	return v
}
