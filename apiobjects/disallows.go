package apiobjects

import "github.com/taiypeo/spotifygo/apierrors"

// Disallows represents a disallows object
// in the Spotify API Object model.
type Disallows struct {
	InterruptingPlayback  bool `json:"interrupting_playback"`
	Pausing               bool `json:"pausing"`
	Resuming              bool `json:"resuming"`
	Seeking               bool `json:"seeking"`
	SkippingNext          bool `json:"skipping_next"`
	SkippingPrev          bool `json:"skipping_prev"`
	TogglingRepeatContext bool `json:"toggling_repeat_context"`
	TogglingShuffle       bool `json:"toggling_shuffle"`
	TogglingRepeatTrack   bool `json:"toggling_repeat_track"`
	TransferringPlayback  bool `json:"transferring_playback"`
}

// Validate returns a TypedError if a Disallows struct is incorrect.
func (disallows Disallows) Validate() apierrors.TypedError {
	return nil
}
