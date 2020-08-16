package profile

import (
	"encoding/json"

	"github.com/taiypeo/spotifygo/apierrors"
	"github.com/taiypeo/spotifygo/apiobjects"
	"github.com/taiypeo/spotifygo/requests"
	"github.com/taiypeo/spotifygo/tokenauth"
)

// GetCurrentUserProfile performs a GET request to /me to receive
// the current user's private user profile.
func GetCurrentUserProfile(
	token tokenauth.Token,
) (apiobjects.PrivateUser, apierrors.TypedError) {
	response, err := requests.GetRestAPI(
		"me/",
		map[string]string{"Authorization": token.GetToken()},
		[]int{200},
	)
	if err != nil {
		return apiobjects.PrivateUser{}, err
	}

	var user apiobjects.PrivateUser
	if err := json.Unmarshal([]byte(response.JSONBody), &user); err != nil {
		return apiobjects.PrivateUser{}, apierrors.NewBasicErrorFromError(err)
	}

	if err := user.Validate(); err != nil {
		return user, err
	}

	return user, nil
}
