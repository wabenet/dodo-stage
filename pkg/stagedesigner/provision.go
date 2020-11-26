package stagedesigner

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os/exec"
	"time"

	"github.com/oclaussen/go-gimme/ssl"
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

	log.Printf("installing docker...")
	if err := InstallDocker(); err != nil {
		return nil, err
	}

	certs, _, err := ssl.GimmeCertificates(&ssl.Options{
		Org:   fmt.Sprintf("dodo.%s", config.Hostname),
		Hosts: []string{ip, "localhost"},
	})
	if err != nil {
		return nil, err
	}

	log.Printf("configuring docker...")
	if err := ConfigureDocker(&DockerConfig{
		CA:          certs.CA,
		ServerCert:  certs.ServerCert,
		ServerKey:   certs.ServerKey,
		Environment: config.Environment,
		Arguments:   config.DockerArgs,
	}); err != nil {
		return nil, err
	}

	if err := AddDockerUser(config.DefaultUser); err != nil {
		return nil, err
	}

	log.Printf("starting docker...")
	if err := RestartDocker(); err != nil {
		return nil, err
	}

	result := &ProvisionResult{
		IPAddress:  ip,
		CA:         string(certs.CA),
		ClientCert: string(certs.ClientCert),
		ClientKey:  string(certs.ClientKey),
	}

	for attempts := 0; attempts < 60; attempts++ {
		if conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", dockerPort)); err == nil {
			conn.Close()
			return result, nil
		}
		time.Sleep(5 * time.Second)
	}

	return nil, errors.New("docker did not start successfully")
}
