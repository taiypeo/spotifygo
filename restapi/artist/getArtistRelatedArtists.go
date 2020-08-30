package artist

import (
	"encoding/json"

	"github.com/taiypeo/spotifygo/apierrors"
	"github.com/taiypeo/spotifygo/apiobjects"
	"github.com/taiypeo/spotifygo/requests"
	"github.com/taiypeo/spotifygo/tokenauth"
)

// GetArtistRelatedArtists performs a GET request to /artists/{artist_id}/related-artists
// to receive a slice of (up to 20) full artist objects.
func GetArtistRelatedArtists(
	token tokenauth.Token,
	artistID string,
) ([]apiobjects.FullArtist, apierrors.TypedError) {
	response, typedErr := requests.GetRestAPI(
		"artists/"+artistID+"/related-artists",
		map[string]string{"Authorization": token.GetToken()},
		[]int{200},
	)
	if typedErr != nil {
		return nil, typedErr
	}

	var artistResponse struct {
		Artists []apiobjects.FullArtist `json:"artists"`
	}
	if basicErr := json.Unmarshal([]byte(response.JSONBody), &artistResponse); basicErr != nil {
		return nil, apierrors.NewBasicErrorFromError(basicErr)
	}

	for _, artist := range artistResponse.Artists {
		if typedErr := artist.Validate(); typedErr != nil {
			return artistResponse.Artists, typedErr
		}
	}

	return artistResponse.Artists, nil
}
