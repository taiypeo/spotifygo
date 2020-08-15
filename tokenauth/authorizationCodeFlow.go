package tokenauth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/taiypeo/spotifygo/requests"
)

type authorizationResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// NewAuthorizationCodeFlow creates a new RefreshableAuthToken.
// authCode is the authorization code returned from the request to the /authorize endpoint;
// redirectURI is the same redirect_uri that was supplied when requesting the authorization code;
// clientId is the Spotify application client id;
// clientSecret is the Spotify application client secret.
func NewAuthorizationCodeFlow(
	authCode,
	redirectURI,
	clientID,
	clientSecret string,
) (RefreshableAuthToken, requests.APIResponse, error) {
	payload := fmt.Sprintf(
		"grant_type=authorization_code&code=%s&redirect_uri=%s",
		authCode,
		redirectURI,
	)
	encodedAuthorizationHeader := "Basic " + base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s:%s", clientID, clientSecret)),
	)

	response, err := requests.PostAuthorization(
		map[string]string{"Authorization": encodedAuthorizationHeader},
		payload,
	)
	if err != nil {
		return RefreshableAuthToken{}, response, err
	}

	var decodedResponse authorizationResponse
	if err := json.Unmarshal([]byte(response.JSONBody), &decodedResponse); err != nil {
		return RefreshableAuthToken{}, response, err
	}

	if decodedResponse.TokenType != "Bearer" {
		return RefreshableAuthToken{}, response, errors.New("token_type is not Bearer")
	}

	var createdRefreshableAuthToken RefreshableAuthToken
	createdRefreshableAuthToken.CreationTime = time.Now()
	createdRefreshableAuthToken.AccessToken = decodedResponse.AccessToken
	createdRefreshableAuthToken.ExpiresIn = decodedResponse.ExpiresIn
	createdRefreshableAuthToken.RefreshToken = decodedResponse.RefreshToken
	createdRefreshableAuthToken.Scope = strings.Split(decodedResponse.Scope, " ")
	return createdRefreshableAuthToken, response, nil
}

// Refresh refreshes the access token using the refresh token.
// clientId is the Spotify application client id;
// clientSecret is the Spotify application client secret.
func (auth *RefreshableAuthToken) Refresh(
	clientID,
	clientSecret string,
) (requests.APIResponse, error) {
	payload := fmt.Sprintf("grant_type=refresh_token&refresh_token=%s", auth.RefreshToken)
	encodedAuthorizationHeader := "Basic " + base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s:%s", clientID, clientSecret)),
	)

	response, err := requests.PostAuthorization(
		map[string]string{"Authorization": encodedAuthorizationHeader},
		payload,
	)
	if err != nil {
		return response, err
	}

	var decodedResponse authorizationResponse
	if err := json.Unmarshal([]byte(response.JSONBody), &decodedResponse); err != nil {
		return response, err
	}

	if decodedResponse.TokenType != "Bearer" {
		return response, errors.New("token_type is not Bearer")
	}

	auth.CreationTime = time.Now()
	auth.AccessToken = decodedResponse.AccessToken
	auth.ExpiresIn = decodedResponse.ExpiresIn
	auth.Scope = strings.Split(decodedResponse.Scope, " ")

	return response, nil
}
