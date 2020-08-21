package apiobjects

import (
	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
)

// Context represents a context object
// in the Spotify API Object model.
type Context struct {
	Type         string      `json:"type"`
	Href         string      `json:"href"`
	ExternalURLs ExternalURL `json:"external_urls"`
	URI          string      `json:"uri"`
}

// Validate returns a TypedError if a Context struct is incorrect.
func (context Context) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	if !stringInSliceCaseIndependent(context.Type, []string{"", "artist", "playlist", "album"}) {
		return apierrors.NewBasicErrorFromString("Unknown type in Context")
	}

	if err := context.ExternalURLs.Validate(); err != nil {
		return err
	}

	return nil
}
