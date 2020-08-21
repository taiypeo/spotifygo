package apiobjects

import (
	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
)

// SimplifiedAlbum represents a simplified album object
// in the Spotify API Object model.
type SimplifiedAlbum struct {
	AlbumGroup           string             `json:"album_group"`
	AlbumType            string             `json:"album_type"`
	Artists              []SimplifiedArtist `json:"artists"`
	AvailableMarkets     []string           `json:"available_markets"`
	ExternalURLs         ExternalURL        `json:"external_urls"`
	Href                 string             `json:"href"`
	ID                   string             `json:"id"`
	Images               []Image            `json:"images"`
	Name                 string             `json:"name"`
	ReleaseDate          string             `json:"release_date"`
	ReleaseDatePrecision string             `json:"release_date_precision"`
	Restrictions         Restrictions       `json:"restrictions"`
	Type                 string             `json:"type"`
	URI                  string             `json:"uri"`
}

// Validate returns a TypedError if a SimplifiedTrack struct is incorrect.
func (album SimplifiedAlbum) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	if !stringInSliceCaseIndependent(
		album.AlbumGroup,
		[]string{"", "album", "single", "compilation", "appears_on"},
	) {
		return apierrors.NewBasicErrorFromString(
			"Unknown AlbumGroup " + album.AlbumGroup + " in SimplifiedAlbum",
		)
	}

	if !stringInSliceCaseIndependent(
		album.AlbumType,
		[]string{"", "album", "single", "compilation"},
	) {
		return apierrors.NewBasicErrorFromString(
			"Unknown AlbumType " + album.AlbumType + " in SimplifiedAlbum",
		)
	}

	for _, artist := range album.Artists {
		if err := artist.Validate(); err != nil {
			return err
		}
	}

	if err := album.ExternalURLs.Validate(); err != nil {
		return err
	}

	for _, image := range album.Images {
		if err := image.Validate(); err != nil {
			return err
		}
	}

	if err := album.Restrictions.Validate(); err != nil {
		return err
	}

	if album.Type != "" && album.Type != "album" {
		return apierrors.NewBasicErrorFromString("Type is not 'album' in SimplifiedAlbum")
	}

	return album.ExternalURLs.Validate()
}
