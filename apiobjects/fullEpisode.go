package apiobjects

import (
	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
)

// FullEpisode represents a full episode object
// in the Spotify API Object model.
type FullEpisode struct {
	Show SimplifiedShow `json:"show"`
	SimplifiedEpisode
}

// Validate returns a TypedError if a FullEpisode struct is incorrect.
func (episode FullEpisode) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	if err := episode.Show.Validate(); err != nil {
		return err
	}

	return episode.SimplifiedEpisode.Validate()
}
