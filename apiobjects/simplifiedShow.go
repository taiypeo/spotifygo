package apiobjects

import (
	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
)

// SimplifiedShow represents a simplified show object
// in the Spotify API Object model.
type SimplifiedShow struct {
	AvailableMarkets   []string    `json:"available_markets"`
	Copyrights         []Copyright `json:"copyrights"`
	Description        string      `json:"description"`
	Explicit           bool        `json:"explicit"`
	ExternalURLs       ExternalURL `json:"external_urls"`
	Href               string      `json:"href"`
	ID                 string      `json:"id"`
	Images             []Image     `json:"images"`
	IsExternallyHosted bool        `json:"is_externally_hosted"`
	Languages          []string    `json:"languages"`
	MediaType          string      `json:"media_type"`
	Name               string      `json:"name"`
	Publisher          string      `json:"publisher"`
	Type               string      `json:"type"`
	URI                string      `json:"uri"`
}

// Validate returns a TypedError if a SimplifiedShow struct is incorrect.
func (show SimplifiedShow) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	for _, copyright := range show.Copyrights {
		if err := copyright.Validate(); err != nil {
			return err
		}
	}

	if err := show.ExternalURLs.Validate(); err != nil {
		return err
	}

	for _, image := range show.Images {
		if err := image.Validate(); err != nil {
			return err
		}
	}

	if show.Type != "" && show.Type != "show" {
		return apierrors.NewBasicErrorFromString("Unknown Type in SimplifiedShow")
	}

	return nil
}
