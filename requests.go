package spotifygo

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var client = &http.Client{}

// APIResponse represents a response from the Spotify REST API,
// where JSONBody is the returned JSON
type APIResponse struct {
	StatusCode int
	JSONBody   string
}

func stringInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if str == s {
			return true
		}
	}

	return false
}

func getFullAPIURL(subURL string) (string, error) {
	const baseAPIURL = "https://api.spotify.com/v1/"
	parsedBaseURL, err := url.Parse(baseAPIURL)
	if err != nil {
		return "", err
	}

	parsedSubURL, err := url.Parse(subURL)
	if err != nil {
		return "", err
	}

	resolvedURL := parsedBaseURL.ResolveReference(parsedSubURL)
	if resolvedURL == nil {
		return "", errors.New("resolvedURL is nil in getFullURL")
	}

	return resolvedURL.String(), nil
}

func makeBasicRequest(
	httpMethod,
	url string,
	headers map[string]string,
	payloadJSON string,
) (APIResponse, error) {
	if !stringInSlice(
		httpMethod,
		[]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	) {
		return APIResponse{}, errors.New("Unsupported HTTP method")
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if httpMethod != http.MethodGet {
		request, err = http.NewRequest(httpMethod, url, strings.NewReader(payloadJSON))
	}
	if err != nil {
		return APIResponse{}, err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		return APIResponse{}, err
	}
	defer response.Body.Close()

	// We can safely use ReadAll here because all the responses
	// will be from the Spotify API, and therefore guaranteed
	// to not be too big
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return APIResponse{}, err
	}

	return APIResponse{StatusCode: response.StatusCode, JSONBody: string(body)}, nil
}

func makeAPIRequest(
	httpMethod,
	subURL string,
	headers map[string]string,
	payloadJSON string,
) (APIResponse, error) {
	url, err := getFullAPIURL(subURL)
	if err != nil {
		return APIResponse{}, err
	}

	updatedHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	for key, value := range headers {
		updatedHeaders[key] = value
	}

	return makeBasicRequest(httpMethod, url, updatedHeaders, payloadJSON)
}

// GetAPI performs an HTTP GET request to a given Spotify API URL
// identified by subURL (part after .../v1/) with
// given headers and returns an APIResponse struct
func GetAPI(subURL string, headers map[string]string) (APIResponse, error) {
	return makeAPIRequest(http.MethodGet, subURL, headers, "")
}

// PostAPI performs an HTTP POST request to a given Spotify API URL
// identified by subURL (part after .../v1/) with
// given headers and a JSON payload and returns an APIResponse struct
func PostAPI(subURL string, headers map[string]string, payloadJSON string) (APIResponse, error) {
	return makeAPIRequest(http.MethodPost, subURL, headers, payloadJSON)
}

// PutAPI performs an HTTP PUT request to a given Spotify API URL
// identified by subURL (part after .../v1/) with
// given headers and a JSON payload and returns an APIResponse struct
func PutAPI(subURL string, headers map[string]string, payloadJSON string) (APIResponse, error) {
	return makeAPIRequest(http.MethodPut, subURL, headers, payloadJSON)
}

// DeleteAPI performs an HTTP DELETE request to a given Spotify API URL
// identified by subURL (part after .../v1/) with
// given headers and returns an APIResponse struct
func DeleteAPI(subURL string, headers map[string]string) (APIResponse, error) {
	return makeAPIRequest(http.MethodDelete, subURL, headers, "")
}
