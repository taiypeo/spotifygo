package profile

import (
	"encoding/json"

	"github.com/taiypeo/spotifygo/apiobjects"
	"github.com/taiypeo/spotifygo/requests"
	"github.com/taiypeo/spotifygo/tokenauth"
)

// GetUserProfile performs a GET request to /users/{user_id} to receive
// a public user profile.
func GetUserProfile(
	token tokenauth.Token,
	userID string,
) (apiobjects.PublicUser, requests.APIResponse, error) {
	response, err := requests.GetRestAPI(
		"users/"+userID,
		map[string]string{"Authorization": token.GetToken()},
		[]int{200},
	)
	if err != nil {
		return apiobjects.PublicUser{}, response, err
	}

	var user apiobjects.PublicUser
	if err := json.Unmarshal([]byte(response.JSONBody), &user); err != nil {
		return apiobjects.PublicUser{}, response, err
	}

	if err := user.Validate(); err != nil {
		return user, response, err
	}

	return user, response, nil
}
