package gitpuller

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

func formatRepoStatus(repo *git.Repository, name string) string {
	out := fmt.Sprintf("Status of '%s':'n", name)

	out += formatRemotes(repo)
	out += "\n"
	out += formatHead(repo)
	out += "\n"

	return out
}

func formatRemotes(repo *git.Repository) string {
	out := "  Remotes:\n"
	remotes, err := repo.Remotes()
	if err != nil {
		out += fmt.Sprintf("Cannot get remotes: %v\n", err)
		return out
	}
	for _, remote := range remotes {
		out += fmt.Sprintf("  - %s\n", remote.String())
	}

	return out
}

func formatHead(repo *git.Repository) string {
	out := "  Head:\n"

	headRef, err := repo.Head()
	if err != nil {
		out += fmt.Sprintf("Cannot get HEAD ref: %v\n", err)
		return out
	}
	commit, err := repo.CommitObject(headRef.Hash())
	if err != nil {
		out += fmt.Sprintf("Cannot get HEAD commit by hash: %v\n", err)
		return out
	}

	out += fmt.Sprintf("  %s", commit.String())
	out += "\n"

	return out
}
