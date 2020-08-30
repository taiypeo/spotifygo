package artist

import (
	"encoding/json"

	"github.com/taiypeo/spotifygo/apierrors"
	"github.com/taiypeo/spotifygo/apiobjects"
	"github.com/taiypeo/spotifygo/requests"
	"github.com/taiypeo/spotifygo/tokenauth"
	"github.com/taiypeo/spotifygo/urltools"
)

// GetArtistTopTracks performs a GET request to /artists/{artist_id}/top-tracks to receive
// a slice of (up to 10) full track objects.
// Country is a mandatory parameter.
func GetArtistTopTracks(
	token tokenauth.Token,
	artistID string,
	country string,
) ([]apiobjects.FullTrack, apierrors.TypedError) {
	url, typedErr := urltools.GetURLWithQueryParameters(
		"artists/"+artistID+"/top-tracks",
		map[string]string{
			"country": country,
		},
	)
	if typedErr != nil {
		return nil, typedErr
	}

	response, typedErr := requests.GetRestAPI(
		url,
		map[string]string{"Authorization": token.GetToken()},
		[]int{200},
	)
	if typedErr != nil {
		return nil, typedErr
	}

	var trackResponse struct {
		Tracks []apiobjects.FullTrack `json:"tracks"`
	}
	if basicErr := json.Unmarshal([]byte(response.JSONBody), &trackResponse); basicErr != nil {
		return nil, apierrors.NewBasicErrorFromError(basicErr)
	}

	for _, track := range trackResponse.Tracks {
		if typedErr := track.Validate(); typedErr != nil {
			return trackResponse.Tracks, typedErr
		}
	}

	return trackResponse.Tracks, nil
}
