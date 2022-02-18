package gitpuller

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/jblew/osowiec-git-backup/util"
)

func fetchAllFromRepo(path string, remoteURL string, auth transport.AuthMethod) (PullResult, error) {
	log.Printf("  Fetching into '%s'\n", path)
	resultType := "empty"

	repo, err := git.PlainOpen(path)
	if err != nil {
		return PullResult{Type: "error"}, fmt.Errorf("Cannot open repo: %v", err)
	}

	err = repo.Fetch(&git.FetchOptions{
		RemoteName: "origin",
		Auth:       auth,
		RefSpecs:   []config.RefSpec{"refs/*:refs/*", "HEAD:refs/heads/HEAD"},
		Force:      true,
	})
	if err == git.NoErrAlreadyUpToDate {
		resultType = "uptodate"
		err = nil
	} else {
		resultType = "changed"
	}
	if err != nil {
		return PullResult{Type: "error"}, fmt.Errorf("Cannot fetch all from repo: %v", err)
	}

	log.Printf(util.IndentMultiline(describeRepo(repo, remoteURL), 2))

	return PullResult{
		Type:          resultType,
		CommitCount:   countCommits(repo, remoteURL),
		BranchesCount: countBranches(repo, remoteURL),
	}, nil
}
