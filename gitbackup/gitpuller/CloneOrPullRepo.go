package gitpuller

import "log"

// CloneOrPullRepo clones or pulls git repo returning an operation log
func CloneOrPullRepo(path string, remoteURL string) (string, error) {
	log.Printf("Pulling '%s' to '%s'\n", remoteURL, path)
	return "", nil
}
