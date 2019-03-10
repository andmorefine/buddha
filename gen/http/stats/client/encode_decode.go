// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// stats HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/andmorefine/buddha/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	stats "github.com/andmorefine/buddha/gen/stats"
	statsviews "github.com/andmorefine/buddha/gen/stats/views"
	goahttp "goa.design/goa/http"
)

// BuildUserNumberRequest instantiates a HTTP request object with method and
// path set to call the "stats" service "user_number" endpoint
func (c *Client) BuildUserNumberRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UserNumberStatsPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("stats", "user_number", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeUserNumberResponse returns a decoder for responses returned by the
// stats user_number endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeUserNumberResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UserNumberResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("stats", "user_number", err)
			}
			p := NewUserNumberStatsuserOK(&body)
			view := "default"
			vres := &statsviews.Statsuser{p, view}
			if err = statsviews.ValidateStatsuser(vres); err != nil {
				return nil, goahttp.ErrValidationError("stats", "user_number", err)
			}
			res := stats.NewStatsuser(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("stats", "user_number", resp.StatusCode, string(body))
		}
	}
}

// unmarshalStatsGuideTypeResponseBodyToStatsGuideTypeView builds a value of
// type *statsviews.StatsGuideTypeView from a value of type
// *StatsGuideTypeResponseBody.
func unmarshalStatsGuideTypeResponseBodyToStatsGuideTypeView(v *StatsGuideTypeResponseBody) *statsviews.StatsGuideTypeView {
	res := &statsviews.StatsGuideTypeView{}
	if v.X != nil {
		res.X = unmarshalStatsLabelTypeResponseBodyToStatsLabelTypeView(v.X)
	}
	if v.Y != nil {
		res.Y = unmarshalStatsLabelTypeResponseBodyToStatsLabelTypeView(v.Y)
	}

	return res
}

// unmarshalStatsLabelTypeResponseBodyToStatsLabelTypeView builds a value of
// type *statsviews.StatsLabelTypeView from a value of type
// *StatsLabelTypeResponseBody.
func unmarshalStatsLabelTypeResponseBodyToStatsLabelTypeView(v *StatsLabelTypeResponseBody) *statsviews.StatsLabelTypeView {
	if v == nil {
		return nil
	}
	res := &statsviews.StatsLabelTypeView{
		Label: v.Label,
	}

	return res
}
