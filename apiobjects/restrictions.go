package apiobjects

import (
	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
)

// Restrictions represents a restrictions object
// in the Spotify API Object model.
// Reason should always be "market" (if restrictions are applied)
// or "" (otherwise).
type Restrictions struct {
	Reason string `json:"reason"`
}

// Validate returns a TypedError if a Restrictions struct is incorrect.
func (restrictions Restrictions) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	if restrictions.Reason != "" && restrictions.Reason != "market" {
		return apierrors.NewBasicErrorFromString("Invalid Reason in Restrictions")
	}

	return nil
}
