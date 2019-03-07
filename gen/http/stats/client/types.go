// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// stats HTTP client types
//
// Command:
// $ goa gen github.com/tonouchi510/goa2-sample/design

package client

import (
	statsviews "github.com/tonouchi510/goa2-sample/gen/stats/views"
	goa "goa.design/goa"
)

// UserNumberResponseBody is the type of the "stats" service "user_number"
// endpoint HTTP response body.
type UserNumberResponseBody struct {
	// グラフデータ
	Data []*DataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// X軸に使用するkey
	X *string `form:"x,omitempty" json:"x,omitempty" xml:"x,omitempty"`
	// Y軸に使用するkey
	Y *string `form:"y,omitempty" json:"y,omitempty" xml:"y,omitempty"`
	// ドットの大きさに使用するkey
	Size *string `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	// ドットの色分けに使用するkey
	Color *string                     `form:"color,omitempty" json:"color,omitempty" xml:"color,omitempty"`
	Guide *StatsGuideTypeResponseBody `form:"guide,omitempty" json:"guide,omitempty" xml:"guide,omitempty"`
}

// DataResponseBody is used to define fields on response body types.
type DataResponseBody struct {
	Key   *string      `form:"key,omitempty" json:"key,omitempty" xml:"key,omitempty"`
	Value *interface{} `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// StatsGuideTypeResponseBody is used to define fields on response body types.
type StatsGuideTypeResponseBody struct {
	X *StatsLabelTypeResponseBody `form:"x,omitempty" json:"x,omitempty" xml:"x,omitempty"`
	Y *StatsLabelTypeResponseBody `form:"y,omitempty" json:"y,omitempty" xml:"y,omitempty"`
}

// StatsLabelTypeResponseBody is used to define fields on response body types.
type StatsLabelTypeResponseBody struct {
	Label *string `form:"label,omitempty" json:"label,omitempty" xml:"label,omitempty"`
}

// NewUserNumberStatsuserOK builds a "stats" service "user_number" endpoint
// result from a HTTP "OK" response.
func NewUserNumberStatsuserOK(body *UserNumberResponseBody) *statsviews.StatsuserView {
	v := &statsviews.StatsuserView{
		X:     body.X,
		Y:     body.Y,
		Size:  body.Size,
		Color: body.Color,
	}
	v.Data = make([]*statsviews.DataView, len(body.Data))
	for i, val := range body.Data {
		v.Data[i] = &statsviews.DataView{
			Key:   val.Key,
			Value: val.Value,
		}
	}
	v.Guide = unmarshalStatsGuideTypeResponseBodyToStatsGuideTypeView(body.Guide)
	return v
}

// ValidateStatsGuideTypeResponseBody runs the validations defined on
// StatsGuideTypeResponseBody
func ValidateStatsGuideTypeResponseBody(body *StatsGuideTypeResponseBody) (err error) {
	if body.X != nil {
		if err2 := ValidateStatsLabelTypeResponseBody(body.X); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Y != nil {
		if err2 := ValidateStatsLabelTypeResponseBody(body.Y); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStatsLabelTypeResponseBody runs the validations defined on
// StatsLabelTypeResponseBody
func ValidateStatsLabelTypeResponseBody(body *StatsLabelTypeResponseBody) (err error) {
	if body.Label == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("label", "body"))
	}
	return
}
