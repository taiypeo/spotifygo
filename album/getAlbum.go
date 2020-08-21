package album

import (
	"encoding/json"

	"github.com/taiypeo/spotifygo/apierrors"
	"github.com/taiypeo/spotifygo/apiobjects"
	"github.com/taiypeo/spotifygo/requests"
	"github.com/taiypeo/spotifygo/tokenauth"
)

// GetAlbum performs a GET request to /albums/{album_id}?market={market} to receive
// a full album object.
// 'market' (same as 'country') has a default value of ''.
func GetAlbum(
	token tokenauth.Token,
	albumID,
	market string,
) (apiobjects.FullAlbum, apierrors.TypedError) {
	query := ""
	if market != "" {
		query = "?market=" + market
	}

	response, err := requests.GetRestAPI(
		"albums/"+albumID+query,
		map[string]string{"Authorization": token.GetToken()},
		[]int{200},
	)
	if err != nil {
		return apiobjects.FullAlbum{}, err
	}

	var album apiobjects.FullAlbum
	if err := json.Unmarshal([]byte(response.JSONBody), &album); err != nil {
		return apiobjects.FullAlbum{}, apierrors.NewBasicErrorFromError(err)
	}

	if err := album.Validate(); err != nil {
		return album, err
	}

	return album, nil
}
