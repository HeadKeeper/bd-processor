package util

import (
	"strings"
	"time"
)

const _URL_SLASH_CONCATENATOR = "/"
const _TAG_CONCATENATOR = ";"

func JoinURL(parts ...string) string {
	return strings.Join(parts, _URL_SLASH_CONCATENATOR)
}

func ConcatTags(tags ...string) string {
	return strings.Join(tags, _TAG_CONCATENATOR)
}

func GetCurrentTimeInMillis() int64 {
	return time.Now().UnixNano() / 1000000
}

func ParseStackOverflowPattern() (string, error) {
	return "", nil
}
