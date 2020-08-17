package tokenauth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
	"github.com/taiypeo/spotifygo/requests"
)

// AuthToken is a struct that represents a simple token
// generated by client credentials flow.
type AuthToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	CreationTime time.Time
}

// Validate returns a TypedError if an AuthToken struct is incorrect.
func (auth *AuthToken) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	if auth.TokenType != "Bearer" {
		return apierrors.NewBasicErrorFromString("TokenType is not Bearer in AuthToken")
	}

	return nil
}

// GetToken returns a token that is used in Spotify REST API to authorize
// user actions.
func (auth *AuthToken) GetToken() string {
	return "Bearer " + auth.AccessToken
}

// NewAuthToken creates a new AuthToken.
// clientId is the Spotify application client id;
// clientSecret is the Spotify application client secret.
func NewAuthToken(
	clientID,
	clientSecret string,
) (AuthToken, apierrors.TypedError) {
	encodedAuthorizationHeader := "Basic " + base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s:%s", clientID, clientSecret)),
	)

	response, err := requests.PostAuthorization(
		map[string]string{"Authorization": encodedAuthorizationHeader},
		"grant_type=client_credentials",
	)
	if err != nil {
		return AuthToken{}, err
	}

	var createdAuthToken AuthToken
	if err := json.Unmarshal([]byte(response.JSONBody), &createdAuthToken); err != nil {
		return AuthToken{}, apierrors.NewBasicErrorFromError(err)
	}
	if err := (&createdAuthToken).Validate(); err != nil {
		return AuthToken{}, err
	}

	createdAuthToken.CreationTime = time.Now()
	return createdAuthToken, nil
}
