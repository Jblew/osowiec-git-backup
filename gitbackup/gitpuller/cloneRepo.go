package gitpuller

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

func cloneRepo(path string, remoteURL string) (string, error) {
	out := fmt.Sprintf("Directory '%s' doesnt exist. Cloning '%s'", path, remoteURL)

	isBare := true
	repo, err := git.PlainClone(path, isBare, &git.CloneOptions{
		URL: remoteURL,
	})
	if err != nil {
		return out, fmt.Errorf("Cannot clone: %v", err)
	}

	out += formatRepoStatus(repo, remoteURL)
	return out, nil
}
