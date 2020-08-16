package apiobjects

import "github.com/taiypeo/spotifygo/apierrors"

// FullArtist represents a full artist object
// in the Spotify API Object model.
type FullArtist struct {
	Followers  Followers `json:"followers"`
	Genres     []string  `json:"genres"`
	Images     []Image   `json:"images"`
	Popularity int64     `json:"popularity"`
	SimplifiedArtist
}

// Validate returns a TypedError if a FullArtist struct is incorrect.
func (artist FullArtist) Validate() apierrors.TypedError {
	if err := artist.Followers.Validate(); err != nil {
		return err
	}
	for _, image := range artist.Images {
		if err := image.Validate(); err != nil {
			return err
		}
	}
	return artist.SimplifiedArtist.Validate()
}
