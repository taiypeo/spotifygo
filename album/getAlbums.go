package album

import (
	"encoding/json"
	"strings"

	"github.com/taiypeo/spotifygo/apierrors"
	"github.com/taiypeo/spotifygo/apiobjects"
	"github.com/taiypeo/spotifygo/requests"
	"github.com/taiypeo/spotifygo/tokenauth"
	"github.com/taiypeo/spotifygo/urltools"
)

// GetAlbums performs a GET request to /albums?ids={ids}&market={market} to receive
// several full album objects (ids in the URL are comma-separated).
// 'market' (same as 'country') has a default value of ''.
func GetAlbums(
	token tokenauth.Token,
	IDs []string,
	market string,
) ([]apiobjects.FullAlbum, apierrors.TypedError) {
	if len(IDs) > 20 {
		return nil, apierrors.NewBasicErrorFromString("IDs cannot be longer than 20")
	}

	params := map[string]string{
		"ids":    strings.Join(IDs, ","),
		"market": market,
	}

	url, basicErr := urltools.GetURLWithQueryParameters("albums", params)
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

	var responseAlbums struct {
		Albums []apiobjects.FullAlbum `json:"albums"`
	}

	if basicErr := json.Unmarshal([]byte(response.JSONBody), &responseAlbums); basicErr != nil {
		return nil, apierrors.NewBasicErrorFromError(basicErr)
	}

	for _, album := range responseAlbums.Albums {
		if typedErr := album.Validate(); typedErr != nil {
			return nil, typedErr
		}
	}

	return responseAlbums.Albums, nil
}
