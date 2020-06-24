package app

import (
	"fmt"
	"gitbackup/gitpuller"
	"gitbackup/util"
	"time"
)

func (app *App) pullRepo(remoteURL string) (string, error) {
	outLog := fmt.Sprintf("# Log for pulling %s\n", remoteURL)

	numOfRetries := 3
	cloneLog, err := app.pullRepoRetry(remoteURL, numOfRetries)
	outLog += cloneLog
	if err != nil {
		outLog += fmt.Sprintf("[ERROR] Cannot pull '%s': %v", remoteURL, err)
		return outLog, fmt.Errorf("Cannot pull '%s': %v", remoteURL, err)
	}
	outLog += "\n"
	return outLog, nil
}

func (app *App) pullRepoRetry(remoteURL string, numOfRetries int) (string, error) {
	out := "\n"
	var lastError error = nil
	for i := 0; i < numOfRetries; i++ {
		out, err := app.doPullRepo(remoteURL)
		if err != nil {
			out += fmt.Sprintf("[PULL FAILED](%s) %v\n", remoteURL, err)
			out += fmt.Sprintf("[PULL FAILED](%s) Retrying %d time in 15s\n", remoteURL, i)
			lastError = err
			time.Sleep(15 * time.Second)
		} else {
			return out, nil
		}
	}
	return out, lastError
}

func (app *App) doPullRepo(remoteURL string) (string, error) {
	repoName, err := util.GetRepoNameFromRemoteURL(remoteURL)
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("%s/%s.git", app.Config.RepositoriesDir, repoName)
	cloneLog, err := gitpuller.CloneOrPullRepo(path, remoteURL, app.Auth)
	if err != nil {
		return "", err
	}

	cloneLogIndented := util.IndentMultiline(cloneLog, 2)
	return cloneLogIndented, nil
}
