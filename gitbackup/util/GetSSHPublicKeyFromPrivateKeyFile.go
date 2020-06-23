package util

import (
	"fmt"
	"io/ioutil"

	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

// GetSSHPublicKeyFromPrivateKeyFile retrives ssh public key from file
func GetSSHPublicKeyFromPrivateKeyFile(filePath string) (*ssh.PublicKeys, error) {
	exists, err := FileExists(filePath)
	if err != nil {
		return nil, fmt.Errorf("Cannot check if key file '%s' exists: %v", filePath, err)
	}
	if exists != true {
		return nil, fmt.Errorf("Key file '%s' does not exist", filePath)
	}

	var publicKey *ssh.PublicKeys
	sshKey, _ := ioutil.ReadFile(filePath)
	publicKey, err = ssh.NewPublicKeys("git", []byte(sshKey), "")
	if err != nil {
		return nil, err
	}
	return publicKey, err
}
