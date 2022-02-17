package util

import (
	"fmt"
	"io/ioutil"
	"net"

	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	xSSH "golang.org/x/crypto/ssh"
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

	var publicKeyAuth *ssh.PublicKeys
	sshKey, _ := ioutil.ReadFile(filePath)
	publicKeyAuth, err = ssh.NewPublicKeys("git", []byte(sshKey), "")
	if err != nil {
		return nil, err
	}
	publicKeyAuth.HostKeyCallback = makeEmptyHostkeyCallback()
	return publicKeyAuth, err
}

func makeEmptyHostkeyCallback() xSSH.HostKeyCallback {
	// allows all known hosts
	return func(hostname string, remote net.Addr, key xSSH.PublicKey) error {
		return nil
	}
}
