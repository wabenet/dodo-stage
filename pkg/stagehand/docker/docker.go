package docker

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
	"time"
)

const (
	Type = "install-docker"

	Port = 2376

	configDir = "/etc/docker"

	systemdUnitPath     = "/etc/systemd/system/docker.service"
	systemdUnitTemplate = `[Service]
LimitNOFILE=infinity
LimitNPROC=infinity
LimitCORE=infinity
ExecStart={{ .DockerdBinary }} -H tcp://0.0.0.0:{{ .DockerPort }} -H unix:///var/run/docker.sock --storage-driver {{ .StorageDriver }} --tlsverify --tlscacert {{ .CACert }} --tlscert {{ .ServerCert }} --tlskey {{ .ServerKey }} {{ range .DockerArgs }}--{{.}} {{ end }}
Environment={{range .Environment }}{{ printf "%q" . }} {{end}}
`

	genericOptionsPath     = "/etc/docker/profile"
	genericOptionsTemplate = `
DOCKER_OPTS='-H tcp://0.0.0.0:{{ .DockerPort }} -H unix:///var/run/docker.sock --storage-driver {{ .StorageDriver }} --tlsverify --tlscacert {{ .CACert }} --tlscert {{ .ServerCert }} --tlskey {{ .ServerKey }}{{ range .DockerArgs}}--{{.}} {{ end }}'
{{range .Environment }}export \"{{ printf "%q" . }}\"
{{end}}
`
)

type Action struct {
	CA          string   `mapstructure:"ca"`
	ServerCert  string   `mapstructure:"cert"`
	ServerKey   string   `mapstructure:"key"`
	Environment []string `mapstructure:"environment"`
	Arguments   []string `mapstructure::"arguments"`
	User        string   `mapstructure:"user"`
}

type optionsContext struct {
	DockerdBinary string
	DockerPort    int
	StorageDriver string
	CACert        string
	ServerCert    string
	ServerKey     string
	Environment   []string
	DockerArgs    []string
}

func (a *Action) Type() string {
	return Type
}

func (a *Action) Execute() error {
	log.Printf("installing docker...")
	if err := install(); err != nil {
		return err
	}

	log.Printf("configuring docker...")
	if err := configure(a); err != nil {
		return err
	}

	if err := addUser(a.User); err != nil {
		return err
	}

	log.Printf("starting docker...")
	if err := restart(); err != nil {
		return err
	}

	if err := validate(); err != nil {
		return err
	}

	return nil
}

func install() error {
	if pacman, err := exec.LookPath("pacman"); err == nil {
		return exec.Command(pacman, "-Sy", "--noconfirm", "--noprogressbar", "docker").Run()
	} else if zypper, err := exec.LookPath("zypper"); err == nil {
		return exec.Command(zypper, "-n", "in", "docker").Run()
	} else if yum, err := exec.LookPath("yum"); err == nil {
		return exec.Command(yum, "install", "-y", "docker").Run()
	} else if aptget, err := exec.LookPath("apt-get"); err == nil {
		if err := exec.Command(aptget, "update").Run(); err != nil {
			return err
		}
		aptcache, err := exec.LookPath("apt-cache")
		if err != nil {
			return err
		}
		for _, pkg := range []string{"docker-ce", "docker.io", "docker-engine", "docker"} {
			out, err := exec.Command(aptcache, "show", "-q", pkg).Output()
			if err == nil && len(out) > 0 {
				cmd := exec.Command(aptget, "install", "-y", pkg)
				cmd.Env = append(os.Environ(), "DEBIAN_FRONTEND=noninteractive")
				return cmd.Run()
			}
		}
	}
	log.Printf("no valid docker installation method found, assuming it is already installed")
	return nil
}

func restart() error {
	if systemctl, err := exec.LookPath("systemctl"); err == nil {
		if err := exec.Command(systemctl, "daemon-reload").Run(); err != nil {
			return err
		}
		if err := exec.Command(systemctl, "-f", "restart", "docker").Run(); err != nil {
			return err
		}
		if err := exec.Command(systemctl, "-f", "enable", "docker").Run(); err != nil {
			return err
		}
		return nil
	} else if service, err := exec.LookPath("service"); err == nil {
		return exec.Command(service, "docker", "restart").Run()
	}
	log.Printf("could not start docker daemon")
	return nil
}

func addUser(user string) error {
	if usermod, err := exec.LookPath("usermod"); err == nil {
		return exec.Command(usermod, "-a", "-G", "docker", user).Run()
	}
	log.Printf("could not modify user")
	return nil
}

func configure(config *Action) error {
	caPath := filepath.Join(configDir, "ca.pem")
	certPath := filepath.Join(configDir, "server.pem")
	keyPath := filepath.Join(configDir, "server-key.pem")

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	if err := ioutil.WriteFile(caPath, []byte(config.CA), 0644); err != nil {
		return err
	}
	if err := ioutil.WriteFile(certPath, []byte(config.ServerCert), 0644); err != nil {
		return err
	}
	if err := ioutil.WriteFile(keyPath, []byte(config.ServerKey), 0644); err != nil {
		return err
	}

	dockerd, err := exec.LookPath("dockerd")
	if err != nil {
		return err
	}

	context := optionsContext{
		DockerdBinary: dockerd,
		DockerPort:    Port,
		StorageDriver: "overlay2",
		CACert:        caPath,
		ServerCert:    certPath,
		ServerKey:     keyPath,
		Environment:   config.Environment,
		DockerArgs:    config.Arguments,
	}

	if _, err := exec.LookPath("systemctl"); err == nil {
		tmpl, err := template.New("systemd").Parse(systemdUnitTemplate)
		if err != nil {
			return err
		}

		var content bytes.Buffer
		tmpl.Execute(&content, context)
		if err := ioutil.WriteFile(systemdUnitPath, content.Bytes(), 0644); err != nil {
			return err
		}
	} else {
		tmpl, err := template.New("dockerOptions").Parse(genericOptionsTemplate)
		if err != nil {
			return err
		}

		var content bytes.Buffer
		tmpl.Execute(&content, context)
		if err := ioutil.WriteFile(genericOptionsPath, content.Bytes(), 0644); err != nil {
			return err
		}
	}

	return nil
}

func validate() error {
	for attempts := 0; attempts < 60; attempts++ {
		if conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", Port)); err == nil {
			conn.Close()
			return nil
		}
		time.Sleep(5 * time.Second)
	}

	return errors.New("docker did not start successfully")
}
