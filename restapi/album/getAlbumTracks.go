package album

import (
	"encoding/json"
	"strconv"

	"github.com/taiypeo/spotifygo/apierrors"
	"github.com/taiypeo/spotifygo/apiobjects"
	"github.com/taiypeo/spotifygo/requests"
	"github.com/taiypeo/spotifygo/tokenauth"
	"github.com/taiypeo/spotifygo/urltools"
)

// GetAlbumTracks performs a GET request to /albums/{album_id}/tracks with the given
// query parameters to receive a paging object of simplified tracks of the album specified
// by the given ID.
// 'market' (same as 'country') has a default value of ''.
func GetAlbumTracks(
	token tokenauth.Token,
	albumID string,
	limit,
	offset int64,
	market string,
) (apiobjects.SimplifiedTrackPaging, apierrors.TypedError) {
	if limit == 0 {
		limit = 20 // default limit value, according to the docs
	} else if limit < 1 || limit > 50 {
		return apiobjects.SimplifiedTrackPaging{},
			apierrors.NewBasicErrorFromString("Limit has to be between 1 and 50")
	}

	if offset < 0 {
		return apiobjects.SimplifiedTrackPaging{},
			apierrors.NewBasicErrorFromString("Offset cannot be negative")
	}

	url, typedErr := urltools.GetURLWithQueryParameters(
		"albums/"+albumID+"/tracks",
		map[string]string{
			"limit":  strconv.FormatInt(limit, 10),
			"offset": strconv.FormatInt(offset, 10),
			"market": market,
		},
	)
	if typedErr != nil {
		return apiobjects.SimplifiedTrackPaging{}, typedErr
	}

	response, typedErr := requests.GetRestAPI(
		url,
		map[string]string{"Authorization": token.GetToken()},
		[]int{200},
	)
	if typedErr != nil {
		return apiobjects.SimplifiedTrackPaging{}, typedErr
	}

	var paging apiobjects.SimplifiedTrackPaging
	if basicErr := json.Unmarshal([]byte(response.JSONBody), &paging); basicErr != nil {
		return apiobjects.SimplifiedTrackPaging{}, apierrors.NewBasicErrorFromError(basicErr)
	}

	if typedErr := paging.Validate(); typedErr != nil {
		return paging, typedErr
	}

	return paging, nil
}
