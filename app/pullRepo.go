package app

import (
	"fmt"
	"log"
	"time"

	"github.com/jblew/osowiec-git-backup/gitpuller"
	"github.com/jblew/osowiec-git-backup/util"
)

func (app *App) pullRepo(remoteURL string) error {
	log.Printf("# Syncing %s\n", remoteURL)

	numOfRetries := 3
	err := app.pullRepoRetry(remoteURL, numOfRetries)
	if err != nil {
		return fmt.Errorf("Cannot pull '%s': %v", remoteURL, err)
	}
	return nil
}

func (app *App) pullRepoRetry(remoteURL string, numOfRetries int) error {
	var lastError error = nil
	for i := 0; i < numOfRetries; i++ {
		err := app.measurePullTimeMetric(func() error { return app.doPullRepo(remoteURL) })
		if err == nil {
			return nil
		}
		log.Printf("[PULL FAILED](%s) %v\n", remoteURL, err)
		log.Printf("[PULL FAILED](%s) Retrying %d time in 15s\n", remoteURL, i)
		app.incRetriesMetric()
		lastError = err
		time.Sleep(15 * time.Second)
	}
	return lastError
}

func (app *App) doPullRepo(remoteURL string) error {
	repoName, err := util.GetRepoNameFromRemoteURL(remoteURL)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s.git", app.Config.RepositoriesDir, repoName)
	result, err := gitpuller.CloneOrPullRepo(path, remoteURL, app.Auth)
	if err != nil {
		app.incPullsMetricFailure()
		return err
	}
	app.incPullsMetricSuccess(result.Type)
	app.incBranchesMetric(result.BranchesCount)
	app.incCommitsMetric(result.CommitCount)
	return nil
}
