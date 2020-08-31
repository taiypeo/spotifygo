package apiobjects

import (
	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
)

// FullShow represents a full show object
// in the Spotify API Object model.
type FullShow struct {
	Episodes SimplifiedEpisodePaging `json:"episodes"`
	SimplifiedShow
}

// Validate returns a TypedError if a FullShow struct is incorrect.
func (show FullShow) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	if err := show.Episodes.Validate(); err != nil {
		return err
	}

	return show.SimplifiedShow.Validate()
}
