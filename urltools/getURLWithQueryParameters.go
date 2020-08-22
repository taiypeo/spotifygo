package urltools

import "net/url"

// GetURLWithQueryParameters adds query parameters to a URL string and returns the resulting URL.
func GetURLWithQueryParameters(baseURL string, params map[string]string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
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
