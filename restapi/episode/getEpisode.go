package episode

import (
	"encoding/json"

	"github.com/taiypeo/spotifygo/apierrors"
	"github.com/taiypeo/spotifygo/apiobjects"
	"github.com/taiypeo/spotifygo/requests"
	"github.com/taiypeo/spotifygo/tokenauth"
	"github.com/taiypeo/spotifygo/urltools"
)

// GetEpisode performs a GET request to /episodes/{episode_id} to receive
// a full episode object.
func GetEpisode(
	token tokenauth.Token,
	episodeID,
	market string,
) (apiobjects.FullEpisode, apierrors.TypedError) {
	url, typedErr := urltools.GetURLWithQueryParameters(
		"episodes/"+episodeID,
		map[string]string{
			"market": market,
		},
	)
	if typedErr != nil {
		return apiobjects.FullEpisode{}, typedErr
	}

	response, typedErr := requests.GetRestAPI(
		url,
		map[string]string{"Authorization": token.GetToken()},
		[]int{200},
	)
	if typedErr != nil {
		return apiobjects.FullEpisode{}, typedErr
	}

	var episode apiobjects.FullEpisode
	if basicErr := json.Unmarshal([]byte(response.JSONBody), &episode); basicErr != nil {
		return apiobjects.FullEpisode{}, apierrors.NewBasicErrorFromError(basicErr)
	}

	if typedErr := episode.Validate(); typedErr != nil {
		return episode, typedErr
	}

	return episode, nil
}
