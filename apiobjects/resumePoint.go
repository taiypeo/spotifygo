package apiobjects

import (
	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
)

// ResumePoint represents a resume point object
// in the Spotify API Object model.
type ResumePoint struct {
	FullyPlayed      bool  `json:"fully_played"`
	ResumePositionMS int64 `json:"resume_position_ms"`
}

// Validate returns a TypedError if a ResumePoint struct is incorrect.
func (point ResumePoint) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	if point.ResumePositionMS < 0 {
		return apierrors.NewBasicErrorFromString("ResumePositionMS is less than 0 in ResumePoint")
	}

	return nil
}
