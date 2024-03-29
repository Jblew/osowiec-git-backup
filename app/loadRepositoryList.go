package app

import (
	"io/ioutil"
	"strings"

	"github.com/jblew/osowiec-git-backup/util"
)

func (app *App) loadRepositoryList() error {
	path := app.Config.RepositoriesListFile
	list, err := doLoadList(path)
	if err != nil {
		return err
	}
	app.setRepositoriesCountMetric(len(list))
	app.Repositories = list
	return nil
}

func doLoadList(path string) ([]string, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return []string{}, err
	}

	rawList := strings.Split(string(contents), "\n")
	return util.DeleteEmptyFromStringSlice(rawList), nil
}
