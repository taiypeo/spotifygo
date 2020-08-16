package profileapi

import (
	"encoding/json"

	"github.com/taiypeo/spotifygo/apiobjects"
	"github.com/taiypeo/spotifygo/requests"
	"github.com/taiypeo/spotifygo/tokenauth"
)

// GetCurrentUserProfile performs a GET request to /me to receive
// the current user's private user profile.
func GetCurrentUserProfile(
	token tokenauth.Token,
) (apiobjects.PrivateUser, requests.APIResponse, error) {
	response, err := requests.GetRestAPI(
		"me/",
		map[string]string{"Authorization": token.GetToken()},
		[]int{200},
	)
	if err != nil {
		return apiobjects.PrivateUser{}, response, err
	}

	var user apiobjects.PrivateUser
	if err := json.Unmarshal([]byte(response.JSONBody), &user); err != nil {
		return apiobjects.PrivateUser{}, response, err
	}

	if err := user.Validate(); err != nil {
		return user, response, err
	}

	return user, response, nil
}
