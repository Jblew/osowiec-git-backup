package gitpuller

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

func cloneRepo(path string, remoteURL string, auth transport.AuthMethod) (string, error) {
	out := fmt.Sprintf("Directory '%s' doesnt exist, cloning '%s'\n", path, remoteURL)

	isBare := true
	repo, err := git.PlainClone(path, isBare, &git.CloneOptions{
		Progress: os.Stdout,
		URL:      remoteURL,
		Auth:     auth,
	})
	if err != nil {
		return out, fmt.Errorf("Cannot clone: %v", err)
	}

	out += formatRepoStatus(repo, remoteURL)
	return out, nil
}
