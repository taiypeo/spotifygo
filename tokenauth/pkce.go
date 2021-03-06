package tokenauth

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
	"github.com/taiypeo/spotifygo/requests"
)

// PKCERefreshableAuthToken is a struct that represents a token that
// can refresh after expiring. It can be generated by
// authorization code flow with proof key for code exchange (PKCE).
// While identical to RefreshableAuthToken, a new struct is necessary
// because the refreshing behavior is different.
type PKCERefreshableAuthToken struct {
	RefreshToken string `json:"refresh_token"`
	ScopedAuthToken
}

// Validate returns a TypedError if an PKCERefreshableAuthToken struct is incorrect.
func (auth *PKCERefreshableAuthToken) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	return auth.ScopedAuthToken.Validate()
}

// NewPKCERefreshableAuthToken creates a new PKCERefreshableAuthToken.
// clientID is the Spotify application client id;
// code is the authorization code obtained after the user is redirected to redirectURI;
// redirectURI is the same redirect_uri that was supplied when requesting the authorization code;
// codeVerifier is the same code verifier that you generate in the beginning of this method
func NewPKCERefreshableAuthToken(
	authCode,
	redirectURI,
	clientID,
	codeVerifier string,
) (PKCERefreshableAuthToken, apierrors.TypedError) {
	payload := fmt.Sprintf(
		"client_id=%s&grant_type=authorization_code&code=%s&redirect_uri=%s&code_verifier=%s",
		clientID,
		authCode,
		redirectURI,
		codeVerifier,
	)

	response, err := requests.PostAuthorization(map[string]string{}, payload)
	if err != nil {
		return PKCERefreshableAuthToken{}, err
	}

	var createdPKCERefreshableAuthToken PKCERefreshableAuthToken
	if err := json.Unmarshal(
		[]byte(response.JSONBody),
		&createdPKCERefreshableAuthToken,
	); err != nil {
		return PKCERefreshableAuthToken{}, apierrors.NewBasicErrorFromError(err)
	}
	if err := (&createdPKCERefreshableAuthToken).Validate(); err != nil {
		return PKCERefreshableAuthToken{}, err
	}

	createdPKCERefreshableAuthToken.CreationTime = time.Now()
	createdPKCERefreshableAuthToken.Scope = strings.Split(
		createdPKCERefreshableAuthToken.ScopeString,
		" ",
	)
	return createdPKCERefreshableAuthToken, nil
}

// Refresh refreshes the access token using the refresh token.
// clientId is the Spotify application client id.
func (auth *PKCERefreshableAuthToken) Refresh(clientID string) apierrors.TypedError {
	payload := fmt.Sprintf(
		"grant_type=refresh_token&refresh_token=%s&client_id=%s",
		auth.RefreshToken,
		clientID,
	)

	response, err := requests.PostAuthorization(map[string]string{}, payload)
	if err != nil {
		return err
	}

	var refreshedToken PKCERefreshableAuthToken
	if err := json.Unmarshal([]byte(response.JSONBody), &refreshedToken); err != nil {
		return apierrors.NewBasicErrorFromError(err)
	}
	if err := (&refreshedToken).Validate(); err != nil {
		return err
	}

	auth.CreationTime = time.Now()
	auth.AccessToken = refreshedToken.AccessToken
	auth.ExpiresIn = refreshedToken.ExpiresIn
	auth.ScopeString = refreshedToken.ScopeString
	auth.Scope = strings.Split(auth.ScopeString, " ")
	auth.RefreshToken = refreshedToken.RefreshToken

	return nil
}
