package apierrors

import (
	"errors"
)

// BasicError represents a TypedError version of the normal error.
type BasicError struct {
	error
}

// NewBasicErrorFromString creates a new BasicError with a given message.
// Notice that this function returns TypedError and not BasicError.
func NewBasicErrorFromString(message string) TypedError {
	return &BasicError{errors.New(message)}
}

// NewBasicErrorFromError creates a new BasicError from an already existing error.
// Notice that this function returns TypedError and not BasicError.
func NewBasicErrorFromError(err error) TypedError {
	return &BasicError{err}
}

// GetType returns the type of BasicError, so BasicError implements TypedError.
func (err *BasicError) GetType() string {
	return "Basic"
}
