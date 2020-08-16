package apierrors

import (
	"encoding/json"

	"github.com/taiypeo/spotifygo"
)

// RestAPIError represents a REST API error object
// as per the documentation (regular error object in the docs).
type RestAPIError struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

// NewRestAPIError creates a new RestAPIError
// from the given APIResponse.
// If json.Unmarshal failed, will return a BasicError, so
// check for the type of the returned value using GetType.
func NewRestAPIError(response spotifygo.APIResponse) TypedError {
	var restAPIError struct {
		Error RestAPIError `json:"error"`
	}

	if err := json.Unmarshal([]byte(response.JSONBody), &restAPIError); err != nil {
		return &BasicError{err}
	}

	return &(restAPIError.Error)
}

func (restAPIError *RestAPIError) Error() string {
	return restAPIError.Message
}

// GetType returns the type of RestAPIError, so RestAPIError implements TypedError.
func (restAPIError *RestAPIError) GetType() ErrorType {
	return RestAPIErrorType
}
