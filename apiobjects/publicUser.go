package apiobjects

import "errors"

// PublicUser represents a public user object
// in the Spotify API Object model.
type PublicUser struct {
	DisplayName  string      `json:"display_name"`
	ExternalURLs ExternalURL `json:"external_urls"`
	Followers    Followers   `json:"followers"`
	Href         string      `json:"href"`
	ID           string      `json:"id"`
	Images       []Image     `json:"images"`
	Type         string      `json:"type"`
	URI          string      `json:"uri"`
}

// Validate returns an error if a PublicUser struct is incorrect.
func (user PublicUser) Validate() error {
	if err := user.ExternalURLs.Validate(); err != nil {
		return err
	}
	if err := user.Followers.Validate(); err != nil {
		return err
	}
	for _, image := range user.Images {
		if err := image.Validate(); err != nil {
			return err
		}
	}
	if user.Type != "user" {
		return errors.New("Type is not 'user' in User")
	}

	return nil
}
