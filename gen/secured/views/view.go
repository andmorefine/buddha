// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// secured views
//
// Command:
// $ goa gen github.com/andmorefine/goa2-sample/design

package views

import (
	goa "goa.design/goa"
)

// GoaJWT is the viewed result type that is projected based on a view.
type GoaJWT struct {
	// Type to project
	Projected *GoaJWTView
	// View to render
	View string
}

// GoaJWTView is a type that runs validations on a projected type.
type GoaJWTView struct {
	// New Jwt token
	Authorization *string
}

// ValidateGoaJWT runs the validations defined on the viewed result type GoaJWT.
func ValidateGoaJWT(result *GoaJWT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateGoaJWTView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateGoaJWTView runs the validations defined on GoaJWTView using the
// "default" view.
func ValidateGoaJWTView(result *GoaJWTView) (err error) {
	if result.Authorization == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "result"))
	}
	return
}
