package artist

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/taiypeo/spotifygo/apierrors"
	"github.com/taiypeo/spotifygo/apiobjects"
	"github.com/taiypeo/spotifygo/requests"
	"github.com/taiypeo/spotifygo/tokenauth"
	"github.com/taiypeo/spotifygo/urltools"
)

// IncludeGroupType represents the include group type for GetArtistAlbums.
type IncludeGroupType int64

// Constants in this enum represent the base include groups (combine them with logical or).
const (
	AlbumIncludeGroup       IncludeGroupType = 1 << iota
	SingleIncludeGroup                       = 1 << iota
	AppearsOnIncludeGroup                    = 1 << iota
	CompilationIncludeGroup                  = 1 << iota
)

func (includeGroup IncludeGroupType) String() (string, apierrors.TypedError) {
	groupToStringMap := map[IncludeGroupType]string{
		AlbumIncludeGroup:       "album",
		SingleIncludeGroup:      "single",
		AppearsOnIncludeGroup:   "appears_on",
		CompilationIncludeGroup: "compilation",
	}

	strGroups := make([]string, 0)
	for group := AlbumIncludeGroup; group <= CompilationIncludeGroup; group <<= 1 {
		if includeGroup&group != 0 {
			strGroup, ok := groupToStringMap[group]
			if !ok {
				return "", apierrors.NewBasicErrorFromString("Unknown IncludeGroupType")
			}

			strGroups = append(strGroups, strGroup)
		}
	}

	return strings.Join(strGroups, ","), nil
}

// GetArtistAlbums performs a GET request to /artists/{artist_id}/albums to receive
// a simplified album paging object that contains the artist's albums.
// You can combine include groups (also defined in getArtistAlbums.go) with logical or.
// The default value for limit is 0, this will set limit to 20, with accordance to the docs.
func GetArtistAlbums(
	token tokenauth.Token,
	artistID string,
	includeGroups IncludeGroupType,
	country string,
	limit,
	offset int64,
) (apiobjects.SimplifiedAlbumPaging, apierrors.TypedError) {
	if limit == 0 {
		limit = 20 // default limit value, according to the docs
	} else if limit < 1 || limit > 50 {
		return apiobjects.SimplifiedAlbumPaging{},
			apierrors.NewBasicErrorFromString("Limit has to be between 1 and 50")
	}

	if offset < 0 {
		return apiobjects.SimplifiedAlbumPaging{},
			apierrors.NewBasicErrorFromString("Offset cannot be negative")
	}

	includeGroupStr, err := includeGroups.String()
	if err != nil {
		return apiobjects.SimplifiedAlbumPaging{}, err
	}

	params := map[string]string{
		"include_groups": includeGroupStr,
		"country":        country,
		"limit":          strconv.FormatInt(limit, 10),
		"offset":         strconv.FormatInt(offset, 10),
	}

	url, typedErr := urltools.GetURLWithQueryParameters("artists/"+artistID+"/albums", params)
	if typedErr != nil {
		return apiobjects.SimplifiedAlbumPaging{}, typedErr
	}

	response, typedErr := requests.GetRestAPI(
		url,
		map[string]string{"Authorization": token.GetToken()},
		[]int{200},
	)
	if typedErr != nil {
		return apiobjects.SimplifiedAlbumPaging{}, typedErr
	}

	var paging apiobjects.SimplifiedAlbumPaging
	if basicErr := json.Unmarshal([]byte(response.JSONBody), &paging); basicErr != nil {
		return apiobjects.SimplifiedAlbumPaging{}, apierrors.NewBasicErrorFromError(basicErr)
	}

	if typedErr := paging.Validate(); typedErr != nil {
		return paging, typedErr
	}

	return paging, nil
}
