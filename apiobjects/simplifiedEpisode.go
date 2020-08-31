package apiobjects

import (
	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
)

// SimplifiedEpisode represents a simplified episode object
// in the Spotify API Object model.
// The field "language" is deliberately removed from SimplifiedEpisode, as it is considered
// deprecated in the API docs.
type SimplifiedEpisode struct {
	AudioPreviewURL      string      `json:"audio_preview_url"`
	Description          string      `json:"description"`
	DurationMS           int64       `json:"duration_ms"`
	Explicit             bool        `json:"explicit"`
	ExternalURLs         ExternalURL `json:"external_urls"`
	Href                 string      `json:"href"`
	ID                   string      `json:"id"`
	Images               []Image     `json:"images"`
	IsExternallyHosted   bool        `json:"is_externally_hosted"`
	IsPlayable           bool        `json:"is_playable"`
	Languages            []string    `json:"languages"`
	Name                 string      `json:"name"`
	ReleaseDate          string      `json:"release_date"`
	ReleaseDatePrecision string      `json:"release_date_precision"`
	ResumePoint          ResumePoint `json:"resume_point"`
	Type                 string      `json:"type"`
	URI                  string      `json:"uri"`
}

// Validate returns a TypedError if a SimplifiedEpisode struct is incorrect.
func (episode SimplifiedEpisode) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	if episode.DurationMS < 0 {
		return apierrors.NewBasicErrorFromString("DurationMS is less than 0 in SimplifiedEpisode")
	}

	if err := episode.ExternalURLs.Validate(); err != nil {
		return err
	}

	for _, image := range episode.Images {
		if err := image.Validate(); err != nil {
			return err
		}
	}

	if err := episode.ResumePoint.Validate(); err != nil {
		return err
	}

	if episode.Type != "" && episode.Type != "episode" {
		return apierrors.NewBasicErrorFromString("Unknown Type in SimplifiedEpisode")
	}

	return nil
}
