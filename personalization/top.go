package personalization

import (
	"encoding/json"
	"fmt"

	"github.com/taiypeo/spotifygo/apierrors"
	"github.com/taiypeo/spotifygo/apiobjects"
	"github.com/taiypeo/spotifygo/requests"
	"github.com/taiypeo/spotifygo/tokenauth"
)

// TimeRange represents the time range for GetUserTopArtistsAndTracks.
type TimeRange int

const (
	// ShortTerm is the 'short_term' time range.
	ShortTerm TimeRange = iota
	// MediumTerm is the 'medium_term' time range.
	MediumTerm
	// LongTerm is the 'long_term' time range.
	LongTerm
)

// GetUserTopArtists performs a GET request to /me/{type} to receive
// the current user's top artists.
// The default value for limit is 0, this will set limit to 20, with accordance to the docs.
func GetUserTopArtists(
	token tokenauth.Token,
	limit int64,
	offset int64,
	timeRange TimeRange,
) (apiobjects.FullArtistPaging, apierrors.TypedError) {
	if limit == 0 {
		limit = 20 // default limit value, according to the docs
	} else if limit < 1 || limit > 50 {
		return apiobjects.FullArtistPaging{},
			apierrors.NewBasicErrorFromString("Limit has to be between 1 and 50")
	}

	if offset < 0 {
		return apiobjects.FullArtistPaging{},
			apierrors.NewBasicErrorFromString("Offset cannot be negative")
	}

	timeRangeString, ok := map[TimeRange]string{
		ShortTerm:  "short_term",
		MediumTerm: "medium_term",
		LongTerm:   "long_term",
	}[timeRange]
	if !ok {
		return apiobjects.FullArtistPaging{},
			apierrors.NewBasicErrorFromString("Unknown time range")
	}

	url := fmt.Sprintf(
		"me/top/artists?limit=%d&offset=%d&time_range=%s",
		limit,
		offset,
		timeRangeString,
	)
	response, err := requests.GetRestAPI(
		url,
		map[string]string{"Authorization": token.GetToken()},
		[]int{200},
	)
	if err != nil {
		return apiobjects.FullArtistPaging{}, err
	}

	var paging apiobjects.FullArtistPaging
	if err := json.Unmarshal([]byte(response.JSONBody), &paging); err != nil {
		return apiobjects.FullArtistPaging{}, apierrors.NewBasicErrorFromError(err)
	}

	if err := paging.Validate(); err != nil {
		return paging, err
	}

	return paging, nil
}

// GetUserTopTracks performs a GET request to /me/{type} to receive
// the current user's top tracks.
// The default value for limit is 0, this will set limit to 20, with accordance to the docs.
func GetUserTopTracks(
	token tokenauth.Token,
	limit int64,
	offset int64,
	timeRange TimeRange,
) (apiobjects.FullTrackPaging, apierrors.TypedError) {
	if limit == 0 {
		limit = 20 // default limit value, according to the docs
	} else if limit < 1 || limit > 50 {
		return apiobjects.FullTrackPaging{},
			apierrors.NewBasicErrorFromString("Limit has to be between 1 and 50")
	}

	if offset < 0 {
		return apiobjects.FullTrackPaging{},
			apierrors.NewBasicErrorFromString("Offset cannot be negative")
	}

	timeRangeString, ok := map[TimeRange]string{
		ShortTerm:  "short_term",
		MediumTerm: "medium_term",
		LongTerm:   "long_term",
	}[timeRange]
	if !ok {
		return apiobjects.FullTrackPaging{},
			apierrors.NewBasicErrorFromString("Unknown time range")
	}

	url := fmt.Sprintf(
		"me/top/tracks?limit=%d&offset=%d&time_range=%s",
		limit,
		offset,
		timeRangeString,
	)
	response, err := requests.GetRestAPI(
		url,
		map[string]string{"Authorization": token.GetToken()},
		[]int{200},
	)
	if err != nil {
		return apiobjects.FullTrackPaging{}, err
	}

	var paging apiobjects.FullTrackPaging
	if err := json.Unmarshal([]byte(response.JSONBody), &paging); err != nil {
		return apiobjects.FullTrackPaging{}, apierrors.NewBasicErrorFromError(err)
	}

	if err := paging.Validate(); err != nil {
		return paging, err
	}

	return paging, nil
}
