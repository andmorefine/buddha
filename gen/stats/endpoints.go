// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// stats endpoints
//
// Command:
// $ goa gen github.com/tonouchi510/goa2-sample/design

package stats

import (
	"context"

	goa "goa.design/goa"
)

// Endpoints wraps the "stats" service endpoints.
type Endpoints struct {
	UserNumber goa.Endpoint
}

// NewEndpoints wraps the methods of the "stats" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		UserNumber: NewUserNumberEndpoint(s),
	}
}

// Use applies the given middleware to all the "stats" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.UserNumber = m(e.UserNumber)
}

// NewUserNumberEndpoint returns an endpoint function that calls the method
// "user_number" of service "stats".
func NewUserNumberEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		res, err := s.UserNumber(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStatsuser(res, "default")
		return vres, nil
	}
}