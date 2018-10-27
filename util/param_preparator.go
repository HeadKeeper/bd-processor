package util

import "strings"

const _URL_SLASH_CONCATENATOR = "/"

func JoinURL(parts ...string) string {
	return strings.Join(parts, _URL_SLASH_CONCATENATOR)
}

func ParseStackOverflowPattern() (string, error) {
	return "", nil
}
