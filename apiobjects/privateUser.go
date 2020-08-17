package apiobjects

import (
	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
)

// PrivateUser represents a private user object
// in the Spotify API Object model.
type PrivateUser struct {
	Country string `json:"country"`
	Email   string `json:"email"`
	Product string `json:"product"`
	PublicUser
}

// Validate returns a TypedError if a PrivateUser struct is incorrect.
func (user PrivateUser) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	return user.PublicUser.Validate()
}
