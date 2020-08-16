package spotifygo

// APIResponse represents a response from the Spotify REST API,
// where JSONBody is the returned JSON.
type APIResponse struct {
	StatusCode int
	JSONBody   string
}
