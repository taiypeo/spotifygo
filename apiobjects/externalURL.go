package apiobjects

import "github.com/taiypeo/spotifygo/apierrors"

// ExternalURL represents an external URL object
// in the Spotify API Object model. As it is just
// an associative key-value store, it is represented
// as a map[string]string in spotifygo.
type ExternalURL map[string]string

// Validate returns a TypedError if an ExternalURL struct is incorrect.
func (url ExternalURL) Validate() apierrors.TypedError {
	return nil
}
