package apierrors

// TypedError is the main error interface used throughout spotifygo.
// It is a "super"-interface of the normal error, as it only adds GetType().
type TypedError interface {
	GetType() ErrorType
	Error() string
}
