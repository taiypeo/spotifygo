package apierrors

import (
	"encoding/json"

	"github.com/taiypeo/spotifygo"
)

// AuthenticationError represents an authentication error object
// as per the documentation.
type AuthenticationError struct {
	StatusCode       int
	ErrorHighLevel   string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// NewAuthenticationError creates a new AuthenticationError
// from the given APIResponse.
// If json.Unmarshal failed, will return a BasicError, so
// check for the type of the returned value using GetType.
func NewAuthenticationError(response spotifygo.APIResponse) TypedError {
	var authError AuthenticationError
	authError.StatusCode = response.StatusCode

	if err := json.Unmarshal([]byte(response.JSONBody), &authError); err != nil {
		return &BasicError{err}
	}

	return &authError
}

func (authError *AuthenticationError) Error() string {
	return authError.ErrorHighLevel
}

// GetType returns the type of AuthenticationError, so AuthenticationError implements TypedError.
func (authError *AuthenticationError) GetType() ErrorType {
	return AuthenticationErrorType
}
