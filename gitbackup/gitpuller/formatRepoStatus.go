package gitpuller

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func formatRepoStatus(repo *git.Repository, name string) string {
	out := ""
	out += formatRemotes(repo)
	out += formatHead(repo)
	out += formatBranches(repo)

	return out
}

func formatRemotes(repo *git.Repository) string {
	out := "Remotes:\n"
	remotes, err := repo.Remotes()
	if err != nil {
		out += fmt.Sprintf("Cannot get remotes: %v\n", err)
		return out
	}
	for _, remote := range remotes {
		remoteDesc := truncate(strings.ReplaceAll(remote.String(), "\n", " / "), 230)
		out += fmt.Sprintf("   - %s\n", remoteDesc)
	}

	return out
}

func formatBranches(repo *git.Repository) string {
	out := "Branches:\n"
	branchesIter, err := repo.Branches()
	if err != nil {
		out += fmt.Sprintf("Cannot get branches: %v\n", err)
		return out
	}
	branchesIter.ForEach(func(branch *plumbing.Reference) error {
		out += "   - "
		out += formatRef(repo, branch)
		out += "\n"
		return nil
	})

	return out
}

func formatHead(repo *git.Repository) string {
	out := "Head: "

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
	out := fmt.Sprintf("[%s] ", ref.Name().String())

	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		out += fmt.Sprintf("Cannot get commit by hash: %v\n", err)
		return out
	}

	commitDesc := strings.ReplaceAll(commit.Message+commit.String(), "\n", " / ")
	out += commitDesc

	return truncate(out, 230)
}

func truncate(str string, max int) string {
	runes := []rune(str)
	if len(runes) <= max {
		return str
	}
	return fmt.Sprintf("%s...", string(runes[:max]))
}
