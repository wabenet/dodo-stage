package stagehand

import (
	"errors"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/wabenet/dodo-stage/pkg/stagehand/service"
)

const (
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

	svc := &service.Service{
		Name:   "dodo-stage-proxy",
		Binary: binary,
		Arguments: []string{
			"proxyserver",
			"--address", config.Address,
			"--tls-ca-file", caPath,
			"--tls-cert-file", certPath,
			"--tls-key-file", keyPath,
		},
		Environment: map[string]string{
			"DODO_LOG_FILE": "-",
		},
	}

	log.Printf("install proxy...")
	if err := svc.Install(); err != nil {
		return err
	}

	log.Printf("stop proxy...")
	if err := svc.Stop(); err != nil {
		return err
	}

	log.Printf("copy binary...")
	if err := installSelf(); err != nil {
		return err
	}

	log.Printf("run proxy...")
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
