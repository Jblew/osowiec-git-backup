package util

import "testing"

func TestGetRepoNameFromRemoteURL(t *testing.T) {
	expected := "osowiec-git-backup"
	got, err := GetRepoNameFromRemoteURL("git@github.com:Jblew/osowiec-git-backup.git")
	if err != nil {
		t.Errorf("TestGetRepoNameFromRemoteURL error %v", err)
	}

	if got != expected {
		t.Errorf("Expected: '%s', got '%s'", expected, got)
	}
}
