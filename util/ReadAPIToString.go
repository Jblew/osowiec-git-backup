package util

import (
	"io/ioutil"
	"net/http"
)

// ReadAPIToString reads external HTTP.GET api to string
func ReadAPIToString(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(responseData), nil
}
