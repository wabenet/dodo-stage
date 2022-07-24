package stagehand

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/oclaussen/go-gimme/ssl"
	"github.com/wabenet/dodo-stage/pkg/stagehand/docker"
)

func Provision(config *Config) (*ProvisionResult, error) {
	log.Printf("replace insecure SSH key")
	if err := ConfigureSSHKeys(config); err != nil {
		return nil, err
	}

	log.Printf("configure host network...")
	ip, err := ConfigureNetwork(Network{Device: "eth1"})
	if err != nil {
		return nil, err
	}

	log.Printf("set hostname...")
	if err := ConfigureHostname(config.Hostname); err != nil {
		return nil, err
	}

	log.Printf("running provision script...")
	for _, script := range config.Script {
		if _, err := exec.Command("/bin/sh", "-c", script).CombinedOutput(); err != nil {
			return nil, err
		}
	}

	certs, _, err := ssl.GimmeCertificates(&ssl.Options{
		Org:   fmt.Sprintf("dodo.%s", config.Hostname),
		Hosts: []string{ip, "localhost"},
	})
	if err != nil {
		return nil, err
	}

	if err = docker.Provision(&docker.Config{
		CA:          certs.CA,
		ServerCert:  certs.ServerCert,
		ServerKey:   certs.ServerKey,
		Environment: config.Environment,
		Arguments:   config.DockerArgs,
		User:        config.DefaultUser,
	}); err != nil {
		return nil, err
	}

	log.Printf("install proxy service")
	if err := InstallProxyService(&ProxyConfig{
		Address:    "tcp://0.0.0.0:20257",
		CA:         certs.CA,
		ServerCert: certs.ServerCert,
		ServerKey:  certs.ServerKey,
	}); err != nil {
		return nil, err
	}

	result := &ProvisionResult{
		IPAddress:  ip,
		CA:         string(certs.CA),
		ClientCert: string(certs.ClientCert),
		ClientKey:  string(certs.ClientKey),
	}

	if err := docker.CheckRunning(); err != nil {
		return nil, err
	}

	if err := CheckProxy(); err != nil {
		return nil, err
	}

	return result, nil
}
