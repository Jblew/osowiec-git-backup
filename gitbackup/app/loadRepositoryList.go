package app

import (
	"gitbackup/util"
	"strings"
)

func (app *App) loadRepositoryList() error {
	url := app.Config.RepositoriesListEndpoint
	list, err := doLoadList(url)
	if err != nil {
		return err
	}
	app.Repositories = list
	return nil
}

func doLoadList(url string) ([]string, error) {
	contents, err := util.ReadAPIToString(url)
	if err != nil {
		return []string{}, err
	}

	rawList := strings.Split(contents, "\n")
	return util.DeleteEmptyFromStringSlice(rawList), nil
}
