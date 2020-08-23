package artist

import (
	"encoding/json"

	"github.com/taiypeo/spotifygo/apierrors"
	"github.com/taiypeo/spotifygo/apiobjects"
	"github.com/taiypeo/spotifygo/requests"
	"github.com/taiypeo/spotifygo/tokenauth"
)

// GetArtist performs a GET request to /artists/{artist_id} to receive
// a full artist object.
func GetArtist(
	token tokenauth.Token,
	artistID string,
) (apiobjects.FullArtist, apierrors.TypedError) {
	response, typedErr := requests.GetRestAPI(
		"artists/"+artistID,
		map[string]string{"Authorization": token.GetToken()},
		[]int{200},
	)
	if typedErr != nil {
		return apiobjects.FullArtist{}, typedErr
	}

	var artist apiobjects.FullArtist
	if basicErr := json.Unmarshal([]byte(response.JSONBody), &artist); basicErr != nil {
		return apiobjects.FullArtist{}, apierrors.NewBasicErrorFromError(basicErr)
	}

	if typedErr := artist.Validate(); typedErr != nil {
		return artist, typedErr
	}

	return artist, nil
}
