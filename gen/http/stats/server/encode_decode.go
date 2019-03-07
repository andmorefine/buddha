// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// stats HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/tonouchi510/goa2-sample/design

package server

import (
	"context"
	"net/http"

	statsviews "github.com/tonouchi510/goa2-sample/gen/stats/views"
	goahttp "goa.design/goa/http"
)

// EncodeUserNumberResponse returns an encoder for responses returned by the
// stats user_number endpoint.
func EncodeUserNumberResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*statsviews.Statsuser)
		enc := encoder(ctx, w)
		body := NewUserNumberResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// marshalStatsGuideTypeViewToStatsGuideTypeResponseBody builds a value of type
// *StatsGuideTypeResponseBody from a value of type
// *statsviews.StatsGuideTypeView.
func marshalStatsGuideTypeViewToStatsGuideTypeResponseBody(v *statsviews.StatsGuideTypeView) *StatsGuideTypeResponseBody {
	res := &StatsGuideTypeResponseBody{}
	if v.X != nil {
		res.X = marshalStatsLabelTypeViewToStatsLabelTypeResponseBody(v.X)
	}
	if v.Y != nil {
		res.Y = marshalStatsLabelTypeViewToStatsLabelTypeResponseBody(v.Y)
	}

	return res
}

// marshalStatsLabelTypeViewToStatsLabelTypeResponseBody builds a value of type
// *StatsLabelTypeResponseBody from a value of type
// *statsviews.StatsLabelTypeView.
func marshalStatsLabelTypeViewToStatsLabelTypeResponseBody(v *statsviews.StatsLabelTypeView) *StatsLabelTypeResponseBody {
	if v == nil {
		return nil
	}
	res := &StatsLabelTypeResponseBody{
		Label: *v.Label,
	}

	return res
}
