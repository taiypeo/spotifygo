package apiobjects

import "github.com/taiypeo/spotifygo/apierrors"

// SimplifiedArtist represents a simplified artist object
// in the Spotify API Object model.
type SimplifiedArtist struct {
	ExternalURLs ExternalURL `json:"external_urls"`
	Href         string      `json:"href"`
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	URI          string      `json:"uri"`
}

// Validate returns a TypedError if a SimplifiedArtist struct is incorrect.
func (artist SimplifiedArtist) Validate() apierrors.TypedError {
	if artist.Type != "" && artist.Type != "artist" {
		return apierrors.NewBasicErrorFromString("Type is not 'artist' is SimplifiedArtist")
	}

	return artist.ExternalURLs.Validate()
}
