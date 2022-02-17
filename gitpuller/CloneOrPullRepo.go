package gitpuller

import (
	"github.com/jblew/osowiec-git-backup/util"

	"github.com/go-git/go-git/v5/plumbing/transport"
)

// CloneOrPullRepo clones or pulls git repo returning an operation log
func CloneOrPullRepo(path string, remoteURL string, auth transport.AuthMethod) (string, error) {
	exists, err := util.FileExists(path)
	if err != nil {
		return "", err
	}
	if exists {
		return fetchAllFromRepo(path, remoteURL, auth)
	}
	return cloneRepo(path, remoteURL, auth)
}
