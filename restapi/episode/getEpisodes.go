package episode

import (
	"encoding/json"
	"strings"

	"github.com/taiypeo/spotifygo/apierrors"
	"github.com/taiypeo/spotifygo/apiobjects"
	"github.com/taiypeo/spotifygo/requests"
	"github.com/taiypeo/spotifygo/tokenauth"
	"github.com/taiypeo/spotifygo/urltools"
)

// GetEpisodes performs a GET request to /episodes to receive
// a slice of full episode objects (up to 50).
func GetEpisodes(
	token tokenauth.Token,
	episodeIDs []string,
	market string,
) ([]apiobjects.FullEpisode, apierrors.TypedError) {
	if len(episodeIDs) > 50 {
		return nil, apierrors.NewBasicErrorFromString(
			"episodeIDs cannot be longer than 50 elements",
		)
	}

	url, typedErr := urltools.GetURLWithQueryParameters(
		"episodes/",
		map[string]string{
			"ids":    strings.Join(episodeIDs, ","),
			"market": market,
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

	var episodesResponse struct {
		Episodes []apiobjects.FullEpisode `json:"episodes"`
	}
	if basicErr := json.Unmarshal([]byte(response.JSONBody), &episodesResponse); basicErr != nil {
		return nil, apierrors.NewBasicErrorFromError(basicErr)
	}

	for _, episode := range episodesResponse.Episodes {
		if typedErr := episode.Validate(); typedErr != nil {
			return episodesResponse.Episodes, typedErr
		}
	}

	return episodesResponse.Episodes, nil
}
