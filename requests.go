package spotifygo

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const baseURL = "https://api.spotify.com/v1/"

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

func getFullURL(subURL string) (string, error) {
	parsedBaseURL, err := url.Parse(baseURL)
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

func makeRequest(
	httpMethod,
	subURL string,
	headers map[string]string,
	payloadJSON string,
) (APIResponse, error) {
	if !stringInSlice(
		httpMethod,
		[]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	) {
		return APIResponse{}, errors.New("Unsupported HTTP method")
	}

	url, err := getFullURL(subURL)
	if err != nil {
		return APIResponse{}, err
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if httpMethod != http.MethodGet {
		request, err = http.NewRequest(httpMethod, url, strings.NewReader(payloadJSON))
	}
	if err != nil {
		return APIResponse{}, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		return APIResponse{}, err
	}
	defer response.Body.Close()

	// We can safely use ReadAll here because all the requests
	// will be to/from the Spotify API, and therefore guaranteed
	// to not be too big
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return APIResponse{}, err
	}

	return APIResponse{StatusCode: response.StatusCode, JSONBody: string(body)}, nil
}

// Get performs an HTTP GET request to a given Spotify API URL
// identified by subURL (part after .../v1/) with
// given headers and returns an APIResponse struct
func Get(subURL string, headers map[string]string) (APIResponse, error) {
	return makeRequest(http.MethodGet, subURL, headers, "")
}

// Post performs an HTTP POST request to a given Spotify API URL
// identified by subURL (part after .../v1/) with
// given headers and a JSON payload and returns an APIResponse struct
func Post(subURL string, headers map[string]string, payloadJSON string) (APIResponse, error) {
	return makeRequest(http.MethodPost, subURL, headers, payloadJSON)
}

// Put performs an HTTP PUT request to a given Spotify API URL
// identified by subURL (part after .../v1/) with
// given headers and a JSON payload and returns an APIResponse struct
func Put(subURL string, headers map[string]string, payloadJSON string) (APIResponse, error) {
	return makeRequest(http.MethodPut, subURL, headers, payloadJSON)
}

// Delete performs an HTTP DELETE request to a given Spotify API URL
// identified by subURL (part after .../v1/) with
// given headers and returns an APIResponse struct
func Delete(subURL string, headers map[string]string) (APIResponse, error) {
	return makeRequest(http.MethodDelete, subURL, headers, "")
}
