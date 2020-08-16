package apiobjects

// PrivateUser represents a private user object
// in the Spotify API Object model.
type PrivateUser struct {
	Country string `json:"country"`
	Email   string `json:"email"`
	Product string `json:"product"`
	PublicUser
}

// Valid returns an error if a PrivateUser struct is incorrect.
func (user PrivateUser) Valid() error {
	return user.PublicUser.Valid()
}
