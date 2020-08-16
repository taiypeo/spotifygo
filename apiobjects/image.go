package apiobjects

import "errors"

// Image represents an image object
// in the Spotify API Object model.
type Image struct {
	Height int64  `json:"height"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
}

// Validate returns an error if an Image struct is incorrect.
func (image Image) Validate() error {
	if image.Height < 0 {
		return errors.New("Height is less than 0 in Image")
	}
	if image.Width < 0 {
		return errors.New("Width is less than 0 in Image")
	}

	return nil
}
