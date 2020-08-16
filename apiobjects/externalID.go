package apiobjects

import "github.com/taiypeo/spotifygo/apierrors"

// ExternalID represents an external ID object
// in the Spotify API Object model. As it is just
// an associative key-value store, it is represented
// as a map[string]string in spotifygo.
type ExternalID map[string]string

// Validate returns a TypedError if an ExternalID struct is incorrect.
func (id ExternalID) Validate() apierrors.TypedError {
	return nil
}
