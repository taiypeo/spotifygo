package requests

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var client = &http.Client{}

// APIResponse represents a response from the Spotify REST API,
// where JSONBody is the returned JSON.
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

func acceptedStatusCode(statusCode int, acceptedCodes []int) bool {
	if len(acceptedCodes) == 0 {
		return true
	}

	for _, code := range acceptedCodes {
		if statusCode == code {
			return true
		}
	}

	return false
}

func getFullRestAPIURL(subURL string) (string, error) {
	const baseRestAPIURL = "https://api.spotify.com/v1/"
	parsedBaseURL, err := url.Parse(baseRestAPIURL)
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
	payload string,
	acceptedStatusCodes []int,
) (APIResponse, error) {
	if !stringInSlice(
		httpMethod,
		[]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	) {
		return APIResponse{}, errors.New("Unsupported HTTP method")
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if httpMethod != http.MethodGet {
		request, err = http.NewRequest(httpMethod, url, strings.NewReader(payload))
	}
	if err != nil {
		return APIResponse{}, err
	}

	request.Header.Set("Accept", "application/json")
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		return APIResponse{}, err
	}
	defer response.Body.Close()

	// We can safely use ReadAll here because all the responses
	// will be from the Spotify API, and are therefore guaranteed
	// to not be too big.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return APIResponse{StatusCode: response.StatusCode}, err
	}

	apiResponse := APIResponse{StatusCode: response.StatusCode, JSONBody: string(body)}

	if !acceptedStatusCode(apiResponse.StatusCode, acceptedStatusCodes) {
		errorMessage := fmt.Sprintf(
			"Got an unsupported status code in a request: %d",
			apiResponse.StatusCode,
		)
		return apiResponse, errors.New(errorMessage)
	}

	return apiResponse, nil
}

func makeRestAPIRequest(
	httpMethod,
	subURL string,
	headers map[string]string,
	payloadJSON string,
	acceptedStatusCodes []int,
) (APIResponse, error) {
	url, err := getFullRestAPIURL(subURL)
	if err != nil {
		return APIResponse{}, err
	}

	updatedHeaders := map[string]string{"Content-Type": "application/json"}
	for key, value := range headers {
		updatedHeaders[key] = value
	}

	return makeBasicRequest(httpMethod, url, updatedHeaders, payloadJSON, acceptedStatusCodes)
}

// GetRestAPI performs an HTTP GET request to a given Spotify REST API URL
// (identified by subURL (part after .../v1/)) with the given headers.
func GetRestAPI(
	subURL string,
	headers map[string]string,
	acceptedStatusCodes []int,
) (APIResponse, error) {
	return makeRestAPIRequest(http.MethodGet, subURL, headers, "", acceptedStatusCodes)
}

// PostRestAPI performs an HTTP POST request to a given Spotify REST API URL
// (identified by subURL (part after .../v1/)) with the given headers and JSON payload.
func PostRestAPI(
	subURL string,
	headers map[string]string,
	payloadJSON string,
	acceptedStatusCodes []int,
) (APIResponse, error) {
	return makeRestAPIRequest(http.MethodPost, subURL, headers, payloadJSON, acceptedStatusCodes)
}

// PutRestAPI performs an HTTP PUT request to a given Spotify REST API URL
// (identified by subURL (part after .../v1/)) with the given headers and JSON payload.
func PutRestAPI(
	subURL string,
	headers map[string]string,
	payloadJSON string,
	acceptedStatusCodes []int,
) (APIResponse, error) {
	return makeRestAPIRequest(http.MethodPut, subURL, headers, payloadJSON, acceptedStatusCodes)
}

// DeleteRestAPI performs an HTTP DELETE request to a given Spotify REST API URL
// (identified by subURL (part after .../v1/)) with the given headers.
func DeleteRestAPI(
	subURL string,
	headers map[string]string,
	acceptedStatusCodes []int,
) (APIResponse, error) {
	return makeRestAPIRequest(http.MethodDelete, subURL, headers, "", acceptedStatusCodes)
}

// PostAuthorization performs an HTTP POST request to the Spotify token API URL
// to retrieve an authorization token. The used authorization flow is specified
// by the headers and the x-www-form-urlencoded payload.
func PostAuthorization(
	headers map[string]string,
	payloadFormURLEncoded string,
) (APIResponse, error) {
	const tokenAPIURL = "https://accounts.spotify.com/api/token"

	updatedHeaders := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	for key, value := range headers {
		updatedHeaders[key] = value
	}

	response, err := makeBasicRequest(
		http.MethodPost,
		tokenAPIURL,
		updatedHeaders,
		payloadFormURLEncoded,
		[]int{200},
	)
	if err != nil {
		return response, err
	}

	return response, nil
}
