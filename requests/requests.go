package requests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/taiypeo/spotifygo"
	"github.com/taiypeo/spotifygo/apierrors"
)

var client = &http.Client{}

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

func getFullRestAPIURL(subURL string) (string, apierrors.TypedError) {
	const baseRestAPIURL = "https://api.spotify.com/v1/"
	parsedBaseURL, err := url.Parse(baseRestAPIURL)
	if err != nil {
		return "", apierrors.NewBasicErrorFromError(err)
	}

	parsedSubURL, err := url.Parse(subURL)
	if err != nil {
		return "", apierrors.NewBasicErrorFromError(err)
	}

	resolvedURL := parsedBaseURL.ResolveReference(parsedSubURL)
	if resolvedURL == nil {
		return "", apierrors.NewBasicErrorFromString("resolvedURL is nil in getFullURL")
	}

	return resolvedURL.String(), nil
}

func makeBasicRequest(
	httpMethod,
	url string,
	headers map[string]string,
	payload string,
	acceptedStatusCodes []int,
	createStatusCodeError func(spotifygo.APIResponse) apierrors.TypedError,
) (spotifygo.APIResponse, apierrors.TypedError) {
	if !stringInSlice(
		httpMethod,
		[]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	) {
		return spotifygo.APIResponse{},
			apierrors.NewBasicErrorFromString("Unsupported HTTP method")
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if httpMethod != http.MethodGet {
		request, err = http.NewRequest(httpMethod, url, strings.NewReader(payload))
	}
	if err != nil {
		return spotifygo.APIResponse{}, apierrors.NewBasicErrorFromError(err)
	}

	request.Header.Set("Accept", "application/json")
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		return spotifygo.APIResponse{}, apierrors.NewBasicErrorFromError(err)
	}
	defer response.Body.Close()

	// We can safely use ReadAll here because all the responses
	// will be from the Spotify API, and are therefore guaranteed
	// to not be too big.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return spotifygo.APIResponse{StatusCode: response.StatusCode},
			apierrors.NewBasicErrorFromError(err)
	}

	apiResponse := spotifygo.APIResponse{StatusCode: response.StatusCode, JSONBody: string(body)}

	if !acceptedStatusCode(apiResponse.StatusCode, acceptedStatusCodes) {
		if createStatusCodeError == nil {
			errorMessage := fmt.Sprintf(
				"Got an unsupported status code in a request: %d",
				apiResponse.StatusCode,
			)
			return apiResponse, apierrors.NewBasicErrorFromString(errorMessage)
		}

		return apiResponse, createStatusCodeError(apiResponse)
	}

	return apiResponse, nil
}

func makeRestAPIRequest(
	httpMethod,
	subURL string,
	headers map[string]string,
	payloadJSON string,
	acceptedStatusCodes []int,
) (spotifygo.APIResponse, apierrors.TypedError) {
	url, err := getFullRestAPIURL(subURL)
	if err != nil {
		return spotifygo.APIResponse{}, err
	}

	updatedHeaders := map[string]string{"Content-Type": "application/json"}
	for key, value := range headers {
		updatedHeaders[key] = value
	}

	return makeBasicRequest(
		httpMethod,
		url,
		updatedHeaders,
		payloadJSON,
		acceptedStatusCodes,
		apierrors.NewRestAPIError,
	)
}

// GetRestAPI performs an HTTP GET request to a given Spotify REST API URL
// (identified by subURL (part after .../v1/)) with the given headers.
func GetRestAPI(
	subURL string,
	headers map[string]string,
	acceptedStatusCodes []int,
) (spotifygo.APIResponse, apierrors.TypedError) {
	return makeRestAPIRequest(http.MethodGet, subURL, headers, "", acceptedStatusCodes)
}

// PostRestAPI performs an HTTP POST request to a given Spotify REST API URL
// (identified by subURL (part after .../v1/)) with the given headers and JSON payload.
func PostRestAPI(
	subURL string,
	headers map[string]string,
	payloadJSON string,
	acceptedStatusCodes []int,
) (spotifygo.APIResponse, apierrors.TypedError) {
	return makeRestAPIRequest(http.MethodPost, subURL, headers, payloadJSON, acceptedStatusCodes)
}

// PutRestAPI performs an HTTP PUT request to a given Spotify REST API URL
// (identified by subURL (part after .../v1/)) with the given headers and JSON payload.
func PutRestAPI(
	subURL string,
	headers map[string]string,
	payloadJSON string,
	acceptedStatusCodes []int,
) (spotifygo.APIResponse, apierrors.TypedError) {
	return makeRestAPIRequest(http.MethodPut, subURL, headers, payloadJSON, acceptedStatusCodes)
}

// DeleteRestAPI performs an HTTP DELETE request to a given Spotify REST API URL
// (identified by subURL (part after .../v1/)) with the given headers.
func DeleteRestAPI(
	subURL string,
	headers map[string]string,
	acceptedStatusCodes []int,
) (spotifygo.APIResponse, apierrors.TypedError) {
	return makeRestAPIRequest(http.MethodDelete, subURL, headers, "", acceptedStatusCodes)
}

// PostAuthorization performs an HTTP POST request to the Spotify token API URL
// to retrieve an authorization token. The used authorization flow is specified
// by the headers and the x-www-form-urlencoded payload.
func PostAuthorization(
	headers map[string]string,
	payloadFormURLEncoded string,
) (spotifygo.APIResponse, apierrors.TypedError) {
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
		apierrors.NewAuthenticationError,
	)
	if err != nil {
		return response, err
	}

	return response, nil
}
