package apiobjects

// ExternalURL represents an external URL object
// in the Spotify API Object model. As it is just
// an associative key-value store, it is represented
// as a map[string]string in spotifygo.
type ExternalURL map[string]string

// Validate returns an error if an ExternalURL struct is incorrect.
func (url ExternalURL) Validate() error {
	return nil
}
