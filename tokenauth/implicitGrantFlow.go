package tokenauth

import (
	"strings"
	"time"

	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
)

// ScopedAuthToken is a struct that represents a token with
// scope that was generated by implicit grant flow.
// NOTE: implicit grant flow is meant to be used in the resource owner's browser
// and not on the backend, therefore the use of this method is discouraged.
type ScopedAuthToken struct {
	Scope       []string `json:"-"`
	ScopeString string   `json:"scope"`
	AuthToken
}

// Validate returns a TypedError if an AuthToken struct is incorrect.
func (auth *ScopedAuthToken) Validate() apierrors.TypedError {
	if !spotifygo.Debug {
		return nil
	}

	return auth.AuthToken.Validate()
}

// NewScopedAuthToken creates a new ScopedAuthToken.
// As the bulk of the work is done in the browser, this function only needs to
// fetch that processed data and save it to a struct.
// accessToken is the access token returned from the /authorize endpoint;
// expiresIn is the time (in seconds) until accessToken expires;
// scope is the slice of scopes provided to /authorize.
func NewScopedAuthToken(accessToken string, expiresIn int64, scope []string) ScopedAuthToken {
	var createdScopedAuthToken ScopedAuthToken
	createdScopedAuthToken.CreationTime = time.Now()
	createdScopedAuthToken.AccessToken = accessToken
	createdScopedAuthToken.ExpiresIn = expiresIn
	createdScopedAuthToken.Scope = scope
	createdScopedAuthToken.ScopeString = strings.Join(scope, " ")
	createdScopedAuthToken.TokenType = "Bearer"
	return createdScopedAuthToken
}
