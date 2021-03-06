package apiobjects

import (
	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
)

// TrackKeyType represents a key the track is in. It is equivalent to int64.
type TrackKeyType int64

// Every constant in this enum block represents a TrackKeyType encoded in its corresponding
// pitch class.
const (
	CKeyType TrackKeyType = iota
	CSharpDFlatKeyType
	DKeyType
	DSharpEFlatKeyType
	EKeyType
	FKeyType
	FSharpGFlatKeyType
	GKeyType
	GSharpAFlatKeyType
	AKeyType
	ASharpBFlatKeyType
	BKeyType
)

func (keyType TrackKeyType) String() (string, apierrors.TypedError) {
	strKeyType, ok := map[TrackKeyType]string{
		CKeyType:           "C",
		CSharpDFlatKeyType: "C♯/D♭",
		DKeyType:           "D",
		DSharpEFlatKeyType: "D♯/E♭",
		EKeyType:           "E",
		FKeyType:           "F",
		FSharpGFlatKeyType: "F♯/G♭",
		GKeyType:           "G",
		GSharpAFlatKeyType: "G♯/A♭",
		AKeyType:           "A",
		ASharpBFlatKeyType: "A♯/B♭",
		BKeyType:           "B",
	}[keyType]

	if !ok {
		return "", apierrors.NewBasicErrorFromString("Unknown TrackKeyType")
	}

	return strKeyType, nil
}

// TrackModeType represents the modality of the track. It is equivalent to int64.
type TrackModeType int64

const (
	// MinorModeType is the minor modality.
	MinorModeType TrackModeType = iota
	// MajorModeType is the major modality.
	MajorModeType
)

func (keyType TrackModeType) String() (string, apierrors.TypedError) {
	if keyType == MinorModeType {
		return "Minor", nil
	} else if keyType == MajorModeType {
		return "Major", nil
	}

	return "", apierrors.NewBasicErrorFromString("Unknown TrackModeType")
}

// AudioFeatures represents an audio features object
// in the Spotify API Object model.
type AudioFeatures struct {
	Acousticness     float64       `json:"acousticness"`
	AnalysisURL      string        `json:"analysis_url"`
	Danceability     float64       `json:"danceability"`
	DurationMS       int64         `json:"duration_ms"`
	Energy           float64       `json:"energy"`
	ID               string        `json:"id"`
	Instrumentalness float64       `json:"instrumentalness"`
	Key              TrackKeyType  `json:"key"`
	Liveness         float64       `json:"liveness"`
	Loudness         float64       `json:"loudness"`
	Mode             TrackModeType `json:"mode"`
	Speechiness      float64       `json:"speechiness"`
	Tempo            float64       `json:"tempo"`
	TimeSignature    int64         `json:"time_signature"`
	TrackHref        string        `json:"track_href"`
	Type             string        `json:"type"`
	URI              string        `json:"URI"`
	Valence          float64       `json:"valence"`
}

// Validate returns a TypedError if a AudioFeatures struct is incorrect.
func (features AudioFeatures) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	if features.Type != "" && features.Type != "audio_features" {
		return apierrors.NewBasicErrorFromString("Unknown Type in AudioFeatures")
	}

	if features.Acousticness < 0 || features.Acousticness > 1 {
		return apierrors.NewBasicErrorFromString("Acousticness is out of bounds in AudioFeatures")
	}

	if features.Danceability < 0 || features.Danceability > 1 {
		return apierrors.NewBasicErrorFromString("Danceability is out of bounds in AudioFeatures")
	}

	if features.DurationMS < 0 {
		return apierrors.NewBasicErrorFromString("DurationMS is less than 0 in AudioFeatures")
	}

	if features.Energy < 0 || features.Energy > 1 {
		return apierrors.NewBasicErrorFromString("Energy is out of bounds in AudioFeatures")
	}

	if features.Instrumentalness < 0 || features.Instrumentalness > 1 {
		return apierrors.NewBasicErrorFromString(
			"Instrumentalness is out of bounds in AudioFeatures",
		)
	}

	if features.Key < CKeyType || features.Key > BKeyType {
		return apierrors.NewBasicErrorFromString("Key is invalid in AudioFeatures")
	}

	if features.Mode != MinorModeType && features.Mode != MajorModeType {
		return apierrors.NewBasicErrorFromString("Mode is invalid in AudioFeatures")
	}

	if features.Liveness < 0 || features.Liveness > 1 {
		return apierrors.NewBasicErrorFromString("Liveness is out of bounds in AudioFeatures")
	}

	if features.Speechiness < 0 || features.Speechiness > 1 {
		return apierrors.NewBasicErrorFromString("Speechiness is out of bounds in AudioFeatures")
	}

	if features.Tempo < 0 {
		return apierrors.NewBasicErrorFromString("Tempo is less than 0 in AudioFeatures")
	}

	if features.Valence < 0 || features.Valence > 1 {
		return apierrors.NewBasicErrorFromString("Valence is out of bounds in AudioFeatures")
	}

	return nil
}
