package app

import (
	"fmt"
	"gitbackup/gitpuller"
	"gitbackup/util"
)

func (app *App) pullRepo(remoteURL string) (bool, string) {
	outLog := fmt.Sprintf("# Log for pulling %s\n", remoteURL)
	cloneLog, err := app.doPullRepo(remoteURL)
	outLog += cloneLog
	if err != nil {
		outLog += fmt.Sprintf("[ERROR] Cannot pull '%s': %v", remoteURL, err)
		return false, outLog
	}
	outLog += "\n"
	return true, outLog
}

func (app *App) doPullRepo(remoteURL string) (string, error) {
	repoName, err := util.GetRepoNameFromRemoteURL(remoteURL)
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("%s/%s", app.Config.RepositoriesDir, repoName)
	cloneLog, err := gitpuller.CloneOrPullRepo(path, remoteURL)
	if err != nil {
		return "", err
	}

	return cloneLog, nil
}
