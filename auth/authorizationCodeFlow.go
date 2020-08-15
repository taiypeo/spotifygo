package auth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/taiypeo/spotifygo/requests"
)

type authorizationResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// AuthorizationCodeFlow contains all the data that describes a successful
// authorization by this method
type AuthorizationCodeFlow struct {
	AccessToken  string
	Scope        []string
	ExpiresIn    int64
	RefreshToken string
}

// NewAuthorizationCodeFlow creates a new AuthorizationCodeFlow.
// authCode is the authorization code returned from the request to the /authorize endpoint
// redirectURI is the same redirect_uri that was supplied when requesting the authorization code
// clientId is the Spotify application client id
// clientSecret is the Spotify application client secret
func NewAuthorizationCodeFlow(
	authCode,
	redirectURI,
	clientID,
	clientSecret string,
) (AuthorizationCodeFlow, requests.APIResponse, error) {
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
		return AuthorizationCodeFlow{}, response, err
	}

	var decodedResponse authorizationResponse
	if err := json.Unmarshal([]byte(response.JSONBody), &decodedResponse); err != nil {
		return AuthorizationCodeFlow{}, response, err
	}

	if decodedResponse.TokenType != "Bearer" {
		return AuthorizationCodeFlow{}, response, errors.New("token_type is not Bearer")
	}

	var createdAuthorizationCodeFlow AuthorizationCodeFlow
	createdAuthorizationCodeFlow.AccessToken = decodedResponse.AccessToken
	createdAuthorizationCodeFlow.ExpiresIn = decodedResponse.ExpiresIn
	createdAuthorizationCodeFlow.RefreshToken = decodedResponse.RefreshToken
	createdAuthorizationCodeFlow.Scope = strings.Split(decodedResponse.Scope, " ")
	return createdAuthorizationCodeFlow, response, nil
}

// GetToken returns a token that is used in Spotify REST API to authorize
// user actions
func (auth *AuthorizationCodeFlow) GetToken() string {
	return "Bearer " + auth.AccessToken
}

// Refresh refreshes the access token using the refresh token
// clientId is the Spotify application client id
// clientSecret is the Spotify application client secret
func (auth *AuthorizationCodeFlow) Refresh(
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

	auth.AccessToken = decodedResponse.AccessToken
	auth.ExpiresIn = decodedResponse.ExpiresIn
	auth.Scope = strings.Split(decodedResponse.Scope, " ")

	return response, nil
}
