package util

import "strings"

const _URL_SLASH_CONCATENATOR = "/"

func joinURL(parts ...string) string {
	return strings.Join(parts, _URL_SLASH_CONCATENATOR)
}

func parseStackOverflowPattern() (string, error) {

}
