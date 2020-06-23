package gitpuller

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

func pullRepo(path string, remoteURL string) (string, error) {
	out := fmt.Sprintf("Directory '%s' doesnt exist. Cloning '%s'", path, remoteURL)

	repo, err := git.PlainOpen(path)
	if err != nil {
		return out, fmt.Errorf("Cannot open repo: %v", err)
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return out, fmt.Errorf("Cannot get repo worktree: %v", err)
	}

	err = worktree.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		return out, fmt.Errorf("Cannot pull worktree: %v", err)
	}

	out += formatRepoStatus(repo, remoteURL)
	return out, nil
}
