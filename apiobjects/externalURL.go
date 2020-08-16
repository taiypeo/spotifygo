package apiobjects

// ExternalURL represents an external URL object
// in the Spotify API Object model. As it is just
// an associative key-value store, it is represented
// as a map[string]string in spotifygo.
type ExternalURL map[string]string

// Valid returns an error if an ExternalURL struct is incorrect.
func (url ExternalURL) Valid() error {
	return nil
}
