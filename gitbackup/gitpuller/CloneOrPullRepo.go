package gitpuller

import (
	"gitbackup/util"
)

// CloneOrPullRepo clones or pulls git repo returning an operation log
func CloneOrPullRepo(path string, remoteURL string) (string, error) {
	exists, err := util.FileExists(path)
	if err != nil {
		return "", err
	}
	if exists {
		return pullRepo(path, remoteURL)
	}
	return cloneRepo(path, remoteURL)
}
