package app

import (
	"fmt"
	"gitbackup/gitpuller"
	"gitbackup/util"
)

func (app *App) pullRepo(remoteURL string) (string, error) {
	outLog := fmt.Sprintf("# Log for pulling %s\n", remoteURL)
	cloneLog, err := app.doPullRepo(remoteURL)
	outLog += cloneLog
	if err != nil {
		outLog += fmt.Sprintf("[ERROR] Cannot pull '%s': %v", remoteURL, err)
		return outLog, fmt.Errorf("Cannot pull '%s': %v", remoteURL, err)
	}
	outLog += "\n"
	return outLog, nil
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

	return cloneLog, nil
}
