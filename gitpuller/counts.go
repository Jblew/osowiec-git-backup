package gitpuller

import (
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func countCommits(repo *git.Repository, name string) int {
	iter, err := repo.CommitObjects()
	if err != nil {
		log.Printf("Cannot count commits in %s: %+v", name, err)
		return 0
	}
	count := 0
	iter.ForEach(func(commit *object.Commit) error {
		count++
		return nil
	})
	return count
}

func countBranches(repo *git.Repository, name string) int {
	iter, err := repo.Branches()
	if err != nil {
		log.Printf("Cannot count commits in %s: %+v", name, err)
		return 0
	}
	count := 0
	iter.ForEach(func(ref *plumbing.Reference) error {
		count++
		return nil
	})
	return count
}
