package app

import (
	"gitbackup/util"
	"io/ioutil"
	"strings"
)

func (app *App) loadRepositoryList() error {
	path := app.Config.RepositoriesListFile
	list, err := doLoadList(path)
	if err != nil {
		return err
	}
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
