// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// users HTTP server types
//
// Command:
// $ goa gen github.com/andmorefine/buddha/design

package server

import (
	"unicode/utf8"

	users "github.com/andmorefine/buddha/gen/users"
	usersviews "github.com/andmorefine/buddha/gen/users/views"
	goa "goa.design/goa"
)

// AddRequestBody is the type of the "users" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// Name of user
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
}

// UpdateRequestBody is the type of the "users" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	// Name of user
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
}

// StoredUserResponseCollection is the type of the "users" service "list"
// endpoint HTTP response body.
type StoredUserResponseCollection []*StoredUserResponse

// ShowResponseBody is the type of the "users" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID is the unique id of the user.
	ID int64 `form:"id" json:"id" xml:"id"`
	// Name of user
	Name string `form:"name" json:"name" xml:"name"`
	// Email of user
	Email string `form:"email" json:"email" xml:"email"`
}

// UpdateResponseBody is the type of the "users" service "update" endpoint HTTP
// response body.
type UpdateResponseBody struct {
	// ID is the unique id of the user.
	ID int64 `form:"id" json:"id" xml:"id"`
	// Name of user
	Name string `form:"name" json:"name" xml:"name"`
	// Email of user
	Email string `form:"email" json:"email" xml:"email"`
}

// ShowNotFoundResponseBody is the type of the "users" service "show" endpoint
// HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// ID of missing user
	ID int64 `form:"id" json:"id" xml:"id"`
}

// StoredUserResponse is used to define fields on response body types.
type StoredUserResponse struct {
	// ID is the unique id of the user.
	ID int64 `form:"id" json:"id" xml:"id"`
	// Name of user
	Name string `form:"name" json:"name" xml:"name"`
	// Email of user
	Email string `form:"email" json:"email" xml:"email"`
}

// NewStoredUserResponseCollection builds the HTTP response body from the
// result of the "list" endpoint of the "users" service.
func NewStoredUserResponseCollection(res usersviews.StoredUserCollectionView) StoredUserResponseCollection {
	body := make([]*StoredUserResponse, len(res))
	for i, val := range res {
		body[i] = &StoredUserResponse{
			ID:    *val.ID,
			Name:  *val.Name,
			Email: *val.Email,
		}
	}
	return body
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "users" service.
func NewShowResponseBody(res *usersviews.StoredUserView) *ShowResponseBody {
	body := &ShowResponseBody{
		ID:    *res.ID,
		Name:  *res.Name,
		Email: *res.Email,
	}
	return body
}

// NewUpdateResponseBody builds the HTTP response body from the result of the
// "update" endpoint of the "users" service.
func NewUpdateResponseBody(res *usersviews.StoredUserView) *UpdateResponseBody {
	body := &UpdateResponseBody{
		ID:    *res.ID,
		Name:  *res.Name,
		Email: *res.Email,
	}
	return body
}

// NewShowNotFoundResponseBody builds the HTTP response body from the result of
// the "show" endpoint of the "users" service.
func NewShowNotFoundResponseBody(res *users.NotFound) *ShowNotFoundResponseBody {
	body := &ShowNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewShowPayload builds a users service show endpoint payload.
func NewShowPayload(id int64) *users.ShowPayload {
	return &users.ShowPayload{
		ID: id,
	}
}

// NewAddUser builds a users service add endpoint payload.
func NewAddUser(body *AddRequestBody) *users.User {
	v := &users.User{
		Name:  *body.Name,
		Email: *body.Email,
	}
	return v
}

// NewUpdatePayload builds a users service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, id int64) *users.UpdatePayload {
	v := &users.UpdatePayload{
		Name:  body.Name,
		Email: body.Email,
	}
	v.ID = id
	return v
}

// NewRemovePayload builds a users service remove endpoint payload.
func NewRemovePayload(id int64) *users.RemovePayload {
	return &users.RemovePayload{
		ID: id,
	}
}

// ValidateAddRequestBody runs the validations defined on AddRequestBody
func ValidateAddRequestBody(body *AddRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 30 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 30, false))
		}
	}
	return
}

// ValidateStoredUserResponse runs the validations defined on StoredUserResponse
func ValidateStoredUserResponse(body *StoredUserResponse) (err error) {
	if utf8.RuneCountInString(body.Name) > 30 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", body.Name, utf8.RuneCountInString(body.Name), 30, false))
	}
	return
}
