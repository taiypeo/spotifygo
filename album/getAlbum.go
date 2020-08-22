package album

import (
	"encoding/json"

	"github.com/taiypeo/spotifygo/apierrors"
	"github.com/taiypeo/spotifygo/apiobjects"
	"github.com/taiypeo/spotifygo/requests"
	"github.com/taiypeo/spotifygo/tokenauth"
	"github.com/taiypeo/spotifygo/urltools"
)

// GetAlbum performs a GET request to /albums/{album_id}?market={market} to receive
// a full album object.
// 'market' (same as 'country') has a default value of ''.
func GetAlbum(
	token tokenauth.Token,
	albumID,
	market string,
) (apiobjects.FullAlbum, apierrors.TypedError) {
	url, basicErr := urltools.GetURLWithQueryParameters("albums/"+albumID, map[string]string{
		"market": market,
	})
	if basicErr != nil {
		return apiobjects.FullAlbum{}, apierrors.NewBasicErrorFromError(basicErr)
	}

	response, typedErr := requests.GetRestAPI(
		url,
		map[string]string{"Authorization": token.GetToken()},
		[]int{200},
	)
	if typedErr != nil {
		return apiobjects.FullAlbum{}, typedErr
	}

	var album apiobjects.FullAlbum
	if basicErr := json.Unmarshal([]byte(response.JSONBody), &album); basicErr != nil {
		return apiobjects.FullAlbum{}, apierrors.NewBasicErrorFromError(basicErr)
	}

	if typedErr := album.Validate(); typedErr != nil {
		return album, typedErr
	}

	return album, nil
}
