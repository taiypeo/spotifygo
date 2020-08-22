package personalization

import (
	"encoding/json"
	"strconv"

	"github.com/taiypeo/spotifygo"
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

func (timeRange TimeRange) String() (string, apierrors.TypedError) {
	timeRangeString, ok := map[TimeRange]string{
		ShortTerm:  "short_term",
		MediumTerm: "medium_term",
		LongTerm:   "long_term",
	}[timeRange]
	if !ok {
		return "", apierrors.NewBasicErrorFromString("Unknown time range")
	}

	return timeRangeString, nil
}

func sendRequest(
	token tokenauth.Token,
	limit int64,
	offset int64,
	timeRange TimeRange,
	personalizationType string,
) (spotifygo.APIResponse, apierrors.TypedError) {
	if limit == 0 {
		limit = 20 // default limit value, according to the docs
	} else if limit < 1 || limit > 50 {
		return spotifygo.APIResponse{},
			apierrors.NewBasicErrorFromString("Limit has to be between 1 and 50")
	}

	if offset < 0 {
		return spotifygo.APIResponse{},
			apierrors.NewBasicErrorFromString("Offset cannot be negative")
	}

	timeRangeString, typedErr := timeRange.String()
	if typedErr != nil {
		return spotifygo.APIResponse{}, typedErr
	}

	url, basicErr := spotifygo.GetURLWithQueryParameters(
		"me/top/"+personalizationType,
		map[string]string{
			"limit":      strconv.FormatInt(limit, 10),
			"offset":     strconv.FormatInt(offset, 10),
			"time_range": timeRangeString,
		},
	)
	if basicErr != nil {
		return spotifygo.APIResponse{}, apierrors.NewBasicErrorFromError(basicErr)
	}

	response, typedErr := requests.GetRestAPI(
		url,
		map[string]string{"Authorization": token.GetToken()},
		[]int{200},
	)
	if typedErr != nil {
		return spotifygo.APIResponse{}, typedErr
	}

	return response, nil
}

// GetUserTopArtists performs a GET request to /me/{type} to receive
// the current user's top artists.
// The default value for limit is 0, this will set limit to 20, with accordance to the docs.
func GetUserTopArtists(
	token tokenauth.Token,
	limit int64,
	offset int64,
	timeRange TimeRange,
) (apiobjects.FullArtistPaging, apierrors.TypedError) {
	response, typedErr := sendRequest(token, limit, offset, timeRange, "artists")
	if typedErr != nil {
		return apiobjects.FullArtistPaging{}, typedErr
	}

	var paging apiobjects.FullArtistPaging
	if basicErr := json.Unmarshal([]byte(response.JSONBody), &paging); basicErr != nil {
		return apiobjects.FullArtistPaging{}, apierrors.NewBasicErrorFromError(basicErr)
	}

	if typedErr := paging.Validate(); typedErr != nil {
		return paging, typedErr
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
	response, typedErr := sendRequest(token, limit, offset, timeRange, "tracks")
	if typedErr != nil {
		return apiobjects.FullTrackPaging{}, typedErr
	}

	var paging apiobjects.FullTrackPaging
	if basicErr := json.Unmarshal([]byte(response.JSONBody), &paging); basicErr != nil {
		return apiobjects.FullTrackPaging{}, apierrors.NewBasicErrorFromError(basicErr)
	}

	if typedErr := paging.Validate(); typedErr != nil {
		return paging, typedErr
	}

	return paging, nil
}
