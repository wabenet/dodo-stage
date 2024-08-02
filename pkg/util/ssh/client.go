package ssh

import (
	"encoding/pem"
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	stage "github.com/wabenet/dodo-stage/api/stage/v1alpha4"
	"golang.org/x/crypto/ssh"
)

var (
	ErrNoPEMData    = errors.New("no PEM data in key file")
	ErrEncryptedKey = errors.New("private key is encrypted")
)

func getClient(opts *stage.SSHOptions) (*ssh.Client, error) {
	signer, err := parseIdentityFile(opts.PrivateKeyFile)
	if err != nil {
		return nil, err
	}

	client, err := ssh.Dial(
		"tcp",
		net.JoinHostPort(opts.Hostname, strconv.Itoa(int(opts.Port))),
		&ssh.ClientConfig{
			User: opts.Username,
			Auth: []ssh.AuthMethod{ssh.PublicKeys(signer)},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil // TODO handle host key callback
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("could not connect to ssh: %w", err)
	}

	return client, nil
}

func parseIdentityFile(path string) (ssh.Signer, error) {
	buffer, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unreadable file: %w", err)
	}

	pemData, _ := pem.Decode(buffer)
	if pemData == nil {
		return nil, ErrNoPEMData
	}

	// TODO: handle password entry
	if strings.Contains(pemData.Headers["Proc-Type"], "ENCRYPTED") {
		return nil, ErrEncryptedKey
	}

	signer, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil, fmt.Errorf("can not read key: %w", err)
	}

	return signer, nil
}
