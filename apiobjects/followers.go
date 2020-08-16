package apiobjects

import "errors"

// Followers represents an followers object
// in the Spotify API Object model.
type Followers struct {
	Href  string `json:"href"`
	Total int64  `json:"total"`
}

// Valid returns an error if an Followers struct is incorrect.
func (followers Followers) Valid() error {
	if followers.Href != "" {
		return errors.New("Href is not empty in Followers")
	}
	if followers.Total < 0 {
		return errors.New("Total is less than 0 in Followers")
	}

	return nil
}
