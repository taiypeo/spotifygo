package profile

import (
	"encoding/json"

	"github.com/taiypeo/spotifygo/apierrors"
	"github.com/taiypeo/spotifygo/apiobjects"
	"github.com/taiypeo/spotifygo/requests"
	"github.com/taiypeo/spotifygo/tokenauth"
)

// GetUserProfile performs a GET request to /users/{user_id} to receive
// a public user profile.
func GetUserProfile(
	token tokenauth.Token,
	userID string,
) (apiobjects.PublicUser, apierrors.TypedError) {
	response, typedErr := requests.GetRestAPI(
		"users/"+userID,
		map[string]string{"Authorization": token.GetToken()},
		[]int{200},
	)
	if typedErr != nil {
		return apiobjects.PublicUser{}, typedErr
	}

	var user apiobjects.PublicUser
	if basicErr := json.Unmarshal([]byte(response.JSONBody), &user); basicErr != nil {
		return apiobjects.PublicUser{}, apierrors.NewBasicErrorFromError(basicErr)
	}

	if typedErr := user.Validate(); typedErr != nil {
		return user, typedErr
	}

	return user, nil
}
