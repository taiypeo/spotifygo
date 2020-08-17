package apiobjects

import (
	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
)

// FullTrack represents a full track object
// in the Spotify API Object model.
type FullTrack struct {
	Album       SimplifiedAlbum `json:"album"`
	ExternalIDs ExternalID      `json:"external_ids"`
	Popularity  int64           `json:"popularity"`
	SimplifiedTrack
}

// Validate returns a TypedError if a FullTrack struct is incorrect.
func (track FullTrack) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	if err := track.Album.Validate(); err != nil {
		return err
	}
	if err := track.ExternalIDs.Validate(); err != nil {
		return err
	}
	if track.Popularity < 0 || track.Popularity > 100 {
		return apierrors.NewBasicErrorFromString("Popularity is out of bounds in FullTrack")
	}

	return track.SimplifiedTrack.Validate()
}
