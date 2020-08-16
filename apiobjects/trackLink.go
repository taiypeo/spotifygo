package apiobjects

import (
	"github.com/taiypeo/spotifygo/apierrors"
)

// TrackLink represents a linked track object
// in the Spotify API Object model.
type TrackLink struct {
	ExternalURLs ExternalURL `json:"external_urls"`
	Href         string      `json:"href"`
	ID           string      `json:"id"`
	Type         string      `json:"type"`
	URI          string      `json:"uri"`
}

// Validate returns a TypedError if a TrackLink struct is incorrect.
func (track TrackLink) Validate() apierrors.TypedError {
	if track.Type != "" && track.Type != "track" {
		return apierrors.NewBasicErrorFromString("Type is not 'track' in TrackLink")
	}

	return track.ExternalURLs.Validate()
}
