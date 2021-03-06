package apiobjects

import (
	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
)

// BasicPaging represents a paging object
// in the Spotify API Object model without the
// items field.
type BasicPaging struct {
	Href     string `json:"href"`
	Limit    int64  `json:"limit"`
	Next     string `json:"next"`
	Offset   int64  `json:"offset"`
	Previous string `json:"previous"`
	Total    int64  `json:"total"`
}

// Validate returns a TypedError if a BasicPaging struct is incorrect.
func (paging BasicPaging) Validate() apierrors.TypedError {
	return nil
}

// FullArtistPaging represents a full artist paging object
// in the Spotify API Object model.
type FullArtistPaging struct {
	Items []FullArtist `json:"items"`
	BasicPaging
}

// Validate returns a TypedError if a FullArtistPaging struct is incorrect.
func (paging FullArtistPaging) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	for _, item := range paging.Items {
		if err := item.Validate(); err != nil {
			return err
		}
	}

	return paging.BasicPaging.Validate()
}

// FullTrackPaging represents a full track paging object
// in the Spotify API Object model.
type FullTrackPaging struct {
	Items []FullTrack `json:"items"`
	BasicPaging
}

// Validate returns a TypedError if a FullTrackPaging struct is incorrect.
func (paging FullTrackPaging) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	for _, item := range paging.Items {
		if err := item.Validate(); err != nil {
			return err
		}
	}

	return paging.BasicPaging.Validate()
}

// SimplifiedTrackPaging represents a simplified track paging object
// in the Spotify API Object model.
type SimplifiedTrackPaging struct {
	Items []SimplifiedTrack `json:"items"`
	BasicPaging
}

// Validate returns a TypedError if a SimplifiedTrackPaging struct is incorrect.
func (paging SimplifiedTrackPaging) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	for _, item := range paging.Items {
		if err := item.Validate(); err != nil {
			return err
		}
	}

	return paging.BasicPaging.Validate()
}

// SimplifiedAlbumPaging represents a simplified album paging object
// in the Spotify API Object model.
type SimplifiedAlbumPaging struct {
	Items []SimplifiedAlbum `json:"items"`
	BasicPaging
}

// Validate returns a TypedError if a SimplifiedAlbumPaging struct is incorrect.
func (paging SimplifiedAlbumPaging) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	for _, item := range paging.Items {
		if err := item.Validate(); err != nil {
			return err
		}
	}

	return paging.BasicPaging.Validate()
}

// SimplifiedEpisodePaging represents a simplified episode paging object
// in the Spotify API Object model.
type SimplifiedEpisodePaging struct {
	Items []SimplifiedEpisode `json:"items"`
	BasicPaging
}

// Validate returns a TypedError if a SimplifiedEpisodePaging struct is incorrect.
func (paging SimplifiedEpisodePaging) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	for _, item := range paging.Items {
		if err := item.Validate(); err != nil {
			return err
		}
	}

	return paging.BasicPaging.Validate()
}
