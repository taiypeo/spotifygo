package apierrors

// ErrorType is the type of a TypedError.
type ErrorType int

const (
	// BasicErrorType is the type for BasicError.
	BasicErrorType = iota
	// AuthenticationErrorType is the type for AuthenticationError.
	AuthenticationErrorType
	// RestAPIErrorType is the type for RestAPIError.
	RestAPIErrorType
)

func (errorType ErrorType) String() (string, TypedError) {
	typeString, ok := map[ErrorType]string{
		BasicErrorType:          "BasicError",
		AuthenticationErrorType: "AuthenticationError",
		RestAPIErrorType:        "RestAPIError",
	}[errorType]
	if !ok {
		return "", NewBasicErrorFromString("Unknown error type")
	}

	return typeString, nil
}
