package tokenauth

// Token represents any type that can return a
// Spotify authorization token.
type Token interface {
	GetToken() string
}
