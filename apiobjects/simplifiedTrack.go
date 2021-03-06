package apiobjects

import (
	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
)

// SimplifiedTrack represents a simplified track object
// in the Spotify API Object model.
type SimplifiedTrack struct {
	Artists          []SimplifiedArtist `json:"artists"`
	AvailableMarkets []string           `json:"available_markets"`
	DiscNumber       int64              `json:"disc_number"`
	DurationMS       int64              `json:"duration_ms"`
	Explicit         bool               `json:"explicit"`
	ExternalURLs     ExternalURL        `json:"external_urls"`
	Href             string             `json:"href"`
	ID               string             `json:"id"`
	IsPlayable       bool               `json:"is_playable"`
	LinkedFrom       TrackLink          `json:"linked_from"`
	Restrictions     Restrictions       `json:"restrictions"`
	Name             string             `json:"name"`
	PreviewURL       string             `json:"preview_url"`
	TrackNumber      int64              `json:"track_number"`
	Type             string             `json:"type"`
	URI              string             `json:"uri"`
	IsLocal          bool               `json:"is_local"`
}

// Validate returns a TypedError if a SimplifiedTrack struct is incorrect.
func (track SimplifiedTrack) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	for _, artist := range track.Artists {
		if err := artist.Validate(); err != nil {
			return err
		}
	}

	if track.DiscNumber < 0 {
		return apierrors.NewBasicErrorFromString("DiscNumber is less than 0 in SimplifiedTrack")
	}

	if track.DurationMS < 0 {
		return apierrors.NewBasicErrorFromString("DurationMS is less than 0 in SimplifiedTrack")
	}

	if err := track.ExternalURLs.Validate(); err != nil {
		return err
	}

	if err := track.LinkedFrom.Validate(); err != nil {
		return err
	}

	if err := track.Restrictions.Validate(); err != nil {
		return err
	}

	if track.TrackNumber < 0 {
		return apierrors.NewBasicErrorFromString("TrackNumber is less than 0 in SimplifiedTrack")
	}

	if track.Type != "" && track.Type != "track" {
		return apierrors.NewBasicErrorFromString("Type is not 'track' in SimplifiedTrack")
	}

	return track.ExternalURLs.Validate()
}
