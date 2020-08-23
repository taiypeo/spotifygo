package artist

import (
	"encoding/json"
	"strings"

	"github.com/taiypeo/spotifygo/apierrors"
	"github.com/taiypeo/spotifygo/apiobjects"
	"github.com/taiypeo/spotifygo/requests"
	"github.com/taiypeo/spotifygo/tokenauth"
	"github.com/taiypeo/spotifygo/urltools"
)

// GetArtists performs a GET request to /artists?ids={artist_ids} to receive
// several full artist objects (ids in the URL are comma-separated).
func GetArtists(
	token tokenauth.Token,
	artistIDs []string,
) ([]apiobjects.FullArtist, apierrors.TypedError) {
	if len(artistIDs) > 50 {
		return nil, apierrors.NewBasicErrorFromString("artistIDs cannot be longer than 50")
	}

	params := map[string]string{
		"ids": strings.Join(artistIDs, ","),
	}

	url, basicErr := urltools.GetURLWithQueryParameters("artists", params)
	if basicErr != nil {
		return nil, apierrors.NewBasicErrorFromError(basicErr)
	}

	response, typedErr := requests.GetRestAPI(
		url,
		map[string]string{"Authorization": token.GetToken()},
		[]int{200},
	)
	if typedErr != nil {
		return nil, typedErr
	}

	var responseArtists struct {
		Artists []apiobjects.FullArtist `json:"artists"`
	}
	if basicErr := json.Unmarshal([]byte(response.JSONBody), &responseArtists); basicErr != nil {
		return nil, apierrors.NewBasicErrorFromError(basicErr)
	}

	for _, artist := range responseArtists.Artists {
		if typedErr := artist.Validate(); typedErr != nil {
			return responseArtists.Artists, typedErr
		}
	}

	return responseArtists.Artists, nil
}
