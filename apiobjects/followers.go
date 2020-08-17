package apiobjects

import (
	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
)

// Followers represents an followers object
// in the Spotify API Object model.
type Followers struct {
	Href  string `json:"href"`
	Total int64  `json:"total"`
}

// Validate returns a TypedError if an Followers struct is incorrect.
func (followers Followers) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	if followers.Href != "" {
		return apierrors.NewBasicErrorFromString("Href is not empty in Followers")
	}
	if followers.Total < 0 {
		return apierrors.NewBasicErrorFromString("Total is less than 0 in Followers")
	}

	return nil
}
