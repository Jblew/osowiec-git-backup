package gitpuller

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/jblew/osowiec-git-backup/util"
)

func cloneRepo(path string, remoteURL string, auth transport.AuthMethod) (PullResult, error) {
	log.Printf("Cloning '%s' into '%s'\n", remoteURL, path)

	isBare := true
	repo, err := git.PlainClone(path, isBare, &git.CloneOptions{
		Progress: os.Stdout,
		URL:      remoteURL,
		Auth:     auth,
	})
	if err != nil {
		return PullResult{Type: "error"}, fmt.Errorf("Cannot clone: %v", err)
	}

	log.Printf(util.IndentMultiline(describeRepo(repo, remoteURL), 2))
	return PullResult{
		Type:          "clone",
		CommitCount:   countCommits(repo, remoteURL),
		BranchesCount: countBranches(repo, remoteURL),
	}, nil
}
