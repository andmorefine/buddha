// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// HTTP request path constructors for the users service.
//
// Command:
// $ goa gen github.com/tonouchi510/goa2-sample/design

package server

import (
	"fmt"
)

// ListUsersPath returns the URL path to the users service list HTTP endpoint.
func ListUsersPath() string {
	return "/api/v1/users"
}

// ShowUsersPath returns the URL path to the users service show HTTP endpoint.
func ShowUsersPath(id int64) string {
	return fmt.Sprintf("/api/v1/users/%v", id)
}

// AddUsersPath returns the URL path to the users service add HTTP endpoint.
func AddUsersPath() string {
	return "/api/v1/users"
}

// UpdateUsersPath returns the URL path to the users service update HTTP endpoint.
func UpdateUsersPath(id int64) string {
	return fmt.Sprintf("/api/v1/users/%v", id)
}

// RemoveUsersPath returns the URL path to the users service remove HTTP endpoint.
func RemoveUsersPath(id int64) string {
	return fmt.Sprintf("/api/v1/users/%v", id)
}
