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
	response, typedErr := requests.GetRestAPI(
		"me/",
		map[string]string{"Authorization": token.GetToken()},
		[]int{200},
	)
	if typedErr != nil {
		return apiobjects.PrivateUser{}, typedErr
	}

	var user apiobjects.PrivateUser
	if basicErr := json.Unmarshal([]byte(response.JSONBody), &user); basicErr != nil {
		return apiobjects.PrivateUser{}, apierrors.NewBasicErrorFromError(basicErr)
	}

	if typedErr := user.Validate(); typedErr != nil {
		return user, typedErr
	}

	return user, nil
}
