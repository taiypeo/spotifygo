package apiobjects

import (
	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
)

// FullAlbum represents a full album object
// in the Spotify API Object model.
// Notice that a full album JSON does not have the 'album_group' field,
// so it should always be empty in a FullAlbum.
type FullAlbum struct {
	Copyrights  []Copyright           `json:"copyrights"`
	ExternalIDs ExternalID            `json:"external_ids"`
	Genres      []string              `json:"genres"`
	Label       string                `json:"label"`
	Popularity  int64                 `json:"popularity"`
	Tracks      SimplifiedTrackPaging `json:"tracks"`
	SimplifiedAlbum
}

// Validate returns a TypedError if a FullAlbum struct is incorrect.
func (album FullAlbum) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	if album.AlbumGroup != "" {
		return apierrors.NewBasicErrorFromString("AlbumGroup is not empty in FullAlbum")
	}

	for _, copyright := range album.Copyrights {
		if err := copyright.Validate(); err != nil {
			return err
		}
	}

	if err := album.ExternalIDs.Validate(); err != nil {
		return err
	}

	if album.Popularity < 0 || album.Popularity > 100 {
		return apierrors.NewBasicErrorFromString("Popularity is out of bounds in FullAlbum")
	}

	if err := album.Tracks.Validate(); err != nil {
		return err
	}

	return album.SimplifiedAlbum.Validate()
}
