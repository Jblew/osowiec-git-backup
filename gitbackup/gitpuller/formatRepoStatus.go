package gitpuller

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func formatRepoStatus(repo *git.Repository, name string) string {
	out := fmt.Sprintf("Status of '%s':'n", name)

	out += formatRemotes(repo)
	out += "\n"
	out += formatHead(repo)
	out += "\n"
	out += formatBranches(repo)
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

func formatBranches(repo *git.Repository) string {
	out := "  Branches:\n"
	branchesIter, err := repo.Branches()
	if err != nil {
		out += fmt.Sprintf("Cannot get branches: %v\n", err)
		return out
	}
	branchesIter.ForEach(func(branch *plumbing.Reference) error {
		out += formatRef(repo, branch)
		return nil
	})

	return out
}

func formatHead(repo *git.Repository) string {
	out := "  Head:\n"

	headRef, err := repo.Head()
	if err != nil {
		out += fmt.Sprintf("Cannot get HEAD ref: %v\n", err)
		return out
	}

	out += formatRef(repo, headRef)
	out += "\n"

	return out
}

func formatRef(repo *git.Repository, ref *plumbing.Reference) string {
	out := fmt.Sprintf("[%s]", ref.Name().String())

	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		out += fmt.Sprintf("Cannot get commit by hash: %v\n", err)
		return out
	}

	out += fmt.Sprintf("  %s", commit.String())

	return out
}
