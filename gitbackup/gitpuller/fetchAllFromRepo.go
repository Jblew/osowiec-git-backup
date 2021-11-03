package gitpuller

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

func fetchAllFromRepo(path string, remoteURL string, auth transport.AuthMethod) (string, error) {
	out := fmt.Sprintf("Fetching into '%s'\n", path)

	repo, err := git.PlainOpen(path)
	if err != nil {
		return out, fmt.Errorf("Cannot open repo: %v", err)
	}

	refs := []config.RefSpec{"refs/*:refs/*", "HEAD:refs/heads/HEAD"}
	err = repo.Fetch(&git.FetchOptions{
		RemoteName: "origin",
		Auth:       auth,
		RefSpecs:   refs,
		Force:      true,
	})
	if err != nil {
		if err == git.NoErrAlreadyUpToDate {
			out += "Repository already up to date \n"
		} else {
			return out, fmt.Errorf("Cannot fetch all from repo: %v", err)
		}
	}

	out += formatRepoStatus(repo, remoteURL)
	return out, nil
}
