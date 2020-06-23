package gitpuller

import (
	"gitbackup/util"

	"github.com/go-git/go-git/v5/plumbing/transport"
)

// CloneOrPullRepo clones or pulls git repo returning an operation log
func CloneOrPullRepo(path string, remoteURL string, auth transport.AuthMethod) (string, error) {
	exists, err := util.FileExists(path)
	if err != nil {
		return "", err
	}
	if exists {
		return pullRepo(path, remoteURL, auth)
	}
	return cloneRepo(path, remoteURL, auth)
}
