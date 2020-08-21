package apiobjects

import "github.com/taiypeo/spotifygo/apierrors"

// Cursor represents a cursor object
// in the Spotify API Object model.
type Cursor struct {
	After string `json:"after"`
}

// Validate returns a TypedError if a Cursor struct is incorrect.
func (cursor Cursor) Validate() apierrors.TypedError {
	return nil
}
