package util

import (
	"strings"
)

const _URL_SLASH_CONCATENATOR = "/"
const _TAG_CONCATENATOR = ";"

func JoinURL(parts ...string) string {
	len := len(parts)
	firstPart := strings.Join(parts[0:len-1], _URL_SLASH_CONCATENATOR)
	urlParts := []string{firstPart, parts[len-1]}
	return strings.Join(urlParts, "?")
}

func ConcatTags(tags ...string) string {
	return strings.Join(tags, _TAG_CONCATENATOR)
}

func GetCurrentTimeInMillis() int64 {
	return 1540944000
}

func ParseStackOverflowPattern() (string, error) {
	return "", nil
}
