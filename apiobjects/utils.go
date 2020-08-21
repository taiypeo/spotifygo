package apiobjects

import "strings"

func stringInSliceCaseIndependent(str string, slice []string) bool {
	str = strings.ToLower(str)
	for _, sliceItem := range slice {
		if str == strings.ToLower(sliceItem) {
			return true
		}
	}

	return false
}
