package ssh

import (
	"errors"
	"net"
	"os"
	"strconv"

	stage "github.com/wabenet/dodo-stage/api/stage/v1alpha3"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

const (
	maxAuthTries = 6
)

func getClient(opts *stage.SSHOptions) (*ssh.Client, error) {
	if opts.Hostname == "" {
		return nil, errors.New("no target host specified")
	}

	if conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		auth := ssh.PublicKeysCallback(agent.NewClient(conn).Signers)
		if client, err := tryConnect(auth, opts); err == nil {
			return client, nil
		}
	}

	signers := []ssh.Signer{}

	signer, err := parseIdentityFile(file)
	if err != nil {
		continue
	}
	signers = append(signers, signer)

	auth := ssh.PublicKeys(signers...)
	if client, err := tryConnect(auth, opts); err == nil {
		return client, nil
	}

	return nil, errors.New("could not connect, all auth methods tried")
}

func tryConnect(auth ssh.AuthMethod, opts *stage.SSHOptions) (*ssh.Client, error) {
	return ssh.Dial(
		"tcp",
		net.JoinHostPort(opts.Hostname, strconv.Itoa(int(opts.Port))),
		&ssh.ClientConfig{
			User: opts.Username,
			Auth: []ssh.AuthMethod{auth},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil // TODO handle host key callback
			},
		},
	)
}

func parseIdentityFile(path string, interactive bool) (ssh.Signer, error) {
	buffer, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "unreadable file")
	}

	pemData, _ := pem.Decode(buffer)
	if pemData == nil {
		return nil, errors.New("not a pem file")
	}

	// TODO: sort encrypted files to the end
	if strings.Contains(pemData.Headers["Proc-Type"], "ENCRYPTED") {
		if !interactive {
			return nil, errors.Wrap(err, "can not decrypt key")
		}
		fmt.Printf("Passphrase for %s: ", path)
		passphrase, err := terminal.ReadPassword(syscall.Stdin)
		fmt.Printf("\n")
		if err != nil {
			return nil, err // User skipped?
		}
		signer, err := ssh.ParsePrivateKeyWithPassphrase(buffer, passphrase)
		if err != nil {
			// TODO: might be wrong passphrase - retry a few times
			return nil, errors.Wrap(err, "could not read key")
		}
		return signer, nil
	}

	signer, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil, errors.Wrap(err, "can not read key")
	}
	return signer, nil
}
