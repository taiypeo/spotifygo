package apiobjects

import (
	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
)

// Copyright represents a copyright object
// in the Spotify API Object model.
type Copyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

// Validate returns a TypedError if a Copyright struct is incorrect.
func (copyright Copyright) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	if copyright.Type != "" && copyright.Type != "C" && copyright.Type != "P" {
		return apierrors.NewBasicErrorFromString("Unknown Type in Copyright")
	}

	return nil
}
