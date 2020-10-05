package util

import "testing"

func TestGetRepoNameFromRemoteURL(t *testing.T) {
	got, err := GetRepoNameFromRemoteURL("git@github.com:Jblew/hipine.git")
	if got != "hipine" {
		t.Errorf("Wrong repo name: %s", got)
	}
	if err != nil {
		t.Error(err)
	}

	got, err = GetRepoNameFromRemoteURL("git@github.com:Jblew/hi.pine.git")
	if got != "hi.pine" {
		t.Errorf("Wrong repo name: %s", got)
	}
	if err != nil {
		t.Error(err)
	}

	got, err = GetRepoNameFromRemoteURL("git@github.com:Jblew/BME280_driver.git")
	if got != "BME280_driver" {
		t.Errorf("Wrong repo name: %s", got)
	}
	if err != nil {
		t.Error(err)
	}

	got, err = GetRepoNameFromRemoteURL("git@github.com:Jblew/hi-pine.git")
	if got != "hi-pine" {
		t.Errorf("Wrong repo name: %s", got)
	}
	if err != nil {
		t.Error(err)
	}

	got, err = GetRepoNameFromRemoteURL("git@github.com:Jblew/hi.st-pine.git")
	if got != "hi.st-pine" {
		t.Errorf("Wrong repo name: %s", got)
	}
	if err != nil {
		t.Error(err)
	}

	got, err = GetRepoNameFromRemoteURL("git@github.com:Jblew/hi.st-pi.ne.git")
	if got != "hi.st-pi.ne" {
		t.Errorf("Wrong repo name: %s", got)
	}
	if err != nil {
		t.Error(err)
	}

	got, err = GetRepoNameFromRemoteURL("git@github.com:Jblew/hi.st-pi.ne")
	if got != "hi.st-pi.ne" {
		t.Errorf("Wrong repo name: %s", got)
	}
	if err != nil {
		t.Error(err)
	}
}
