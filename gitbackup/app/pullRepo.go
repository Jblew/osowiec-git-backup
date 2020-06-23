package app

import (
	"fmt"
	"gitbackup/gitpuller"
	"gitbackup/util"
	"log"
)

func (app *App) pullRepo(remoteURL string) error {
	repoName, err := util.GetRepoNameFromRemoteURL(remoteURL)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s", app.Config.RepositoriesDir, repoName)
	cloneLog, err := gitpuller.CloneOrPullRepo(path, remoteURL)
	if err != nil {
		return err
	}

	log.Printf("# Log for pulling %s\n", remoteURL)
	log.Println(cloneLog)
	log.Println()
	return nil
}
