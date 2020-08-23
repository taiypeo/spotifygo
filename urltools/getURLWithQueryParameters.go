package urltools

import (
	"net/url"

	"github.com/taiypeo/spotifygo/apierrors"
)

// GetURLWithQueryParameters adds query parameters to a URL string and returns the resulting URL.
func GetURLWithQueryParameters(
	baseURL string,
	params map[string]string,
) (string, apierrors.TypedError) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", apierrors.NewBasicErrorFromError(err)
	}

	query := u.Query()

	for key, value := range params {
		if key != "" && value != "" {
			query.Set(key, value)
		}
	}

	u.RawQuery = query.Encode()
	return u.String(), nil
}
