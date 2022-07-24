package stagehand

import (
	"errors"
	"io"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/kardianos/service"
)

const (
	name          = "dodo-stage-proxy"
	qualifiedName = "com.wabenet.dodo.stage.proxy"
	description   = "dodo proxy"

	configDir = "/etc/dodo"
	binary    = "/usr/local/bin/dodo-stagehand"
)

type ProxyConfig struct {
	Address    string
	CA         []byte
	ServerCert []byte
	ServerKey  []byte
}

type program struct{}

func (p *program) Start(_ service.Service) error {
	return nil
}

func (p *program) Stop(_ service.Service) error {
	return nil
}

func InstallProxyService(config *ProxyConfig) error {
	if err := installSelf(); err != nil {
		return err
	}

	caPath := filepath.Join(configDir, "ca.pem")
	certPath := filepath.Join(configDir, "server.pem")
	keyPath := filepath.Join(configDir, "server-key.pem")

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	if err := os.WriteFile(caPath, config.CA, 0600); err != nil {
		return err
	}
	if err := os.WriteFile(certPath, config.ServerCert, 0600); err != nil {
		return err
	}
	if err := os.WriteFile(keyPath, config.ServerKey, 0600); err != nil {
		return err
	}

	scfg := &service.Config{
		Name:        qualifiedName,
		DisplayName: name,
		Description: description,
		Option:      map[string]interface{}{},
		Executable:  binary,
		Arguments: []string{
			"proxyserver",
			"--address", config.Address,
			"--tls-ca-file", caPath,
			"--tls-cert-file", certPath,
			"--tls-key-file", keyPath,
		},
	}

	if u, err := user.Current(); err == nil && u.Uid != "0" {
		scfg.UserName = u.Username
	}

	svc, err := service.New(&program{}, scfg)
	if err != nil {
		return err
	}

	// TODO: implement an update-and-reload on service

	if err := svc.Uninstall(); err != nil {
		return err
	}

	if err := svc.Install(); err != nil {
		return err
	}

	if err := svc.Start(); err != nil {
		return err
	}

	return nil
}

func installSelf() error {
	self, err := os.Executable()
	if err != nil {
		return err
	}

	binaryIn, err := os.Open(self)
	if err != nil {
		return err
	}

	defer binaryIn.Close()

	binaryOut, err := os.Create(binary)
	if err != nil {
		return err
	}

	defer binaryOut.Close()

	if _, err := io.Copy(binaryOut, binaryIn); err != nil {
		return err
	}

	if err := binaryOut.Chmod(0o755); err != nil {
		return err
	}

	return nil
}

func CheckProxy() error {
	for attempts := 0; attempts < 60; attempts++ {
		if conn, err := net.Dial("tcp", "127.0.0.1:20257"); err == nil {
			conn.Close()
			return nil
		}
		time.Sleep(5 * time.Second)
	}

	return errors.New("proxy did not start successfully")
}
