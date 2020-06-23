package util

import (
	"fmt"
	"regexp"
)

// GetRepoNameFromRemoteURL parses standard bare repo git url and returns repo name
func GetRepoNameFromRemoteURL(remoteURL string) (string, error) {
	var re = regexp.MustCompile(`(?m)^.*\/([a-zA-Z0-9-]*).*$`)

	matches := re.FindStringSubmatch(remoteURL)
	if len(matches) != 2 {
		return "", fmt.Errorf("Malformed repo remote URL '%s'", remoteURL)
	}

	// [0] contains ful match
	return matches[1], nil
}
