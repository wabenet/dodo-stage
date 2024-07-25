package hostname

import (
	"io/ioutil"
	"log"
	"os/exec"
)

const (
	Type = "set-hostname"
)

type Action struct {
	Hostname string `mapstructure:"hostname"`
}

func (a *Action) Type() string {
	return Type
}

func (a *Action) Execute() error {
	log.Printf("set hostname...")

	if err := ioutil.WriteFile("/etc/hostname", []byte(a.Hostname), 0644); err != nil {
		return err
	}

	if hostnamectl, err := exec.LookPath("hostnamectl"); err == nil {
		return exec.Command(hostnamectl, "set-hostname", a.Hostname).Run()
	} else if hostname, err := exec.LookPath("hostname"); err == nil {
		return exec.Command(hostname, a.Hostname).Run()
	} else {
		// TODO what to do?
		return nil
	}
}
