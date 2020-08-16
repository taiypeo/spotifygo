package apiobjects

import "github.com/taiypeo/spotifygo/apierrors"

// Image represents an image object
// in the Spotify API Object model.
type Image struct {
	Height int64  `json:"height"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
}

// Validate returns a TypedError if an Image struct is incorrect.
func (image Image) Validate() apierrors.TypedError {
	if image.Height < 0 {
		return apierrors.NewBasicErrorFromString("Height is less than 0 in Image")
	}
	if image.Width < 0 {
		return apierrors.NewBasicErrorFromString("Width is less than 0 in Image")
	}

	return nil
}
