// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// users service
//
// Command:
// $ goa gen github.com/tonouchi510/goa2-sample/design

package users

import (
	"context"

	usersviews "github.com/tonouchi510/goa2-sample/gen/users/views"
)

// users serves user relative information.
type Service interface {
	// List all stored users
	List(context.Context) (res StoredUserCollection, err error)
	// Show user by ID
	Show(context.Context, *ShowPayload) (res *StoredUser, err error)
	// Add new user and return its ID.
	Add(context.Context, *User) (res int64, err error)
	// Update user item.
	Update(context.Context, *UpdatePayload) (res *StoredUser, err error)
	// Remove user from storage
	Remove(context.Context, *RemovePayload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "users"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"list", "show", "add", "update", "remove"}

// StoredUserCollection is the result type of the users service list method.
type StoredUserCollection []*StoredUser

// ShowPayload is the payload type of the users service show method.
type ShowPayload struct {
	// ID of user to show
	ID int64
}

// StoredUser is the result type of the users service show method.
type StoredUser struct {
	// ID is the unique id of the user.
	ID int64
	// Name of user
	Name string
	// Email of user
	Email string
}

// User is the payload type of the users service add method.
type User struct {
	// Name of user
	Name string
	// Email of user
	Email string
}

// UpdatePayload is the payload type of the users service update method.
type UpdatePayload struct {
	// ID of user to show
	ID int64
	// Name of user
	Name *string
	// Email of user
	Email *string
}

// RemovePayload is the payload type of the users service remove method.
type RemovePayload struct {
	// ID of user to remove
	ID int64
}

// NotFound is the type returned when attempting to show or delete a user that
// does not exist.
type NotFound struct {
	// Message of error
	Message string
	// ID of missing user
	ID int64
}

// Error returns an error description.
func (e *NotFound) Error() string {
	return "NotFound is the type returned when attempting to show or delete a user that does not exist."
}

// ErrorName returns "NotFound".
func (e *NotFound) ErrorName() string {
	return e.Message
}

// NewStoredUserCollection initializes result type StoredUserCollection from
// viewed result type StoredUserCollection.
func NewStoredUserCollection(vres usersviews.StoredUserCollection) StoredUserCollection {
	var res StoredUserCollection
	switch vres.View {
	case "default", "":
		res = newStoredUserCollection(vres.Projected)
	}
	return res
}

// NewViewedStoredUserCollection initializes viewed result type
// StoredUserCollection from result type StoredUserCollection using the given
// view.
func NewViewedStoredUserCollection(res StoredUserCollection, view string) usersviews.StoredUserCollection {
	var vres usersviews.StoredUserCollection
	switch view {
	case "default", "":
		p := newStoredUserCollectionView(res)
		vres = usersviews.StoredUserCollection{p, "default"}
	}
	return vres
}

// NewStoredUser initializes result type StoredUser from viewed result type
// StoredUser.
func NewStoredUser(vres *usersviews.StoredUser) *StoredUser {
	var res *StoredUser
	switch vres.View {
	case "default", "":
		res = newStoredUser(vres.Projected)
	}
	return res
}

// NewViewedStoredUser initializes viewed result type StoredUser from result
// type StoredUser using the given view.
func NewViewedStoredUser(res *StoredUser, view string) *usersviews.StoredUser {
	var vres *usersviews.StoredUser
	switch view {
	case "default", "":
		p := newStoredUserView(res)
		vres = &usersviews.StoredUser{p, "default"}
	}
	return vres
}

// newStoredUserCollection converts projected type StoredUserCollection to
// service type StoredUserCollection.
func newStoredUserCollection(vres usersviews.StoredUserCollectionView) StoredUserCollection {
	res := make(StoredUserCollection, len(vres))
	for i, n := range vres {
		res[i] = newStoredUser(n)
	}
	return res
}

// newStoredUserCollectionView projects result type StoredUserCollection into
// projected type StoredUserCollectionView using the "default" view.
func newStoredUserCollectionView(res StoredUserCollection) usersviews.StoredUserCollectionView {
	vres := make(usersviews.StoredUserCollectionView, len(res))
	for i, n := range res {
		vres[i] = newStoredUserView(n)
	}
	return vres
}

// newStoredUser converts projected type StoredUser to service type StoredUser.
func newStoredUser(vres *usersviews.StoredUserView) *StoredUser {
	res := &StoredUser{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Email != nil {
		res.Email = *vres.Email
	}
	return res
}

// newStoredUserView projects result type StoredUser into projected type
// StoredUserView using the "default" view.
func newStoredUserView(res *StoredUser) *usersviews.StoredUserView {
	vres := &usersviews.StoredUserView{
		ID:    &res.ID,
		Name:  &res.Name,
		Email: &res.Email,
	}
	return vres
}
