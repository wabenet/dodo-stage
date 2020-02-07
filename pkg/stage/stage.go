package stage

import (
	"github.com/oclaussen/dodo/pkg/types"
)

type Stage interface {
	Initialize(string, *types.Stage) error
	Cleanup()
	Create() error
	Start() error
	Stop() error
	Remove(bool, bool) error
	Exist() (bool, error)
	Available() (bool, error)
	GetSSHOptions() (*SSHOptions, error)
	GetDockerOptions() (*DockerOptions, error)
}

type SSHOptions struct {
	Hostname       string
	Port           int
	Username       string
	PrivateKeyFile string
}

type DockerOptions struct {
	Version  string
	Host     string
	CAFile   string
	CertFile string
	KeyFile  string
}
