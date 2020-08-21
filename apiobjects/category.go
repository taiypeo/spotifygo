package apiobjects

import (
	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
)

// Category represents a category object
// in the Spotify API Object model.
type Category struct {
	Href  string  `json:"href"`
	Icons []Image `json:"icons"`
	ID    string  `json:"id"`
	Name  string  `json:"name"`
}

// Validate returns a TypedError if a Category struct is incorrect.
func (category Category) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	for _, icon := range category.Icons {
		if err := icon.Validate(); err != nil {
			return err
		}
	}

	return nil
}
