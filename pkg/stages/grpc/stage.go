package grpc

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	dodo "github.com/dodo-cli/dodo-core/pkg/plugin"
	"github.com/dodo-cli/dodo-stage/pkg/stage"
	"github.com/dodo-cli/dodo-stage/pkg/types"
	"github.com/hashicorp/go-plugin"
	"github.com/oclaussen/go-gimme/configfiles"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Plugin struct {
	plugin.NetRPCUnsupportedPlugin
	Impl stage.Stage
}

func (p *Plugin) GRPCServer(_ *plugin.GRPCBroker, server *grpc.Server) error {
	types.RegisterDockerStageServer(server, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *Plugin) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, client *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: types.NewDockerStageClient(client)}, nil
}

type Stage struct {
	wrapped stage.Stage
	client  *plugin.Client
}

func (s *Stage) Initialize(name string, conf *types.Stage) error {
	path, err := findPluginExecutable(conf.Type)
	if err != nil {
		return err
	}

	s.client = plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  dodo.ProtocolVersion,
			MagicCookieKey:   dodo.MagicCookieKey,
			MagicCookieValue: dodo.MagicCookieValue,
		},
		Plugins:          map[string]plugin.Plugin{"stage": &Plugin{}},
		Cmd:              exec.Command(path),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})

	c, err := s.client.Client()
	if err != nil {
		return err
	}
	raw, err := c.Dispense("stage")
	if err != nil {
		return err
	}

	s.wrapped = raw.(stage.Stage)
	return s.wrapped.Initialize(name, conf)
}

func findPluginExecutable(name string) (string, error) {
	directories, err := configfiles.GimmeConfigDirectories(&configfiles.Options{
		Name:                      "dodo",
		IncludeWorkingDirectories: true,
	})
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("plugin-%s_%s_%s", name, runtime.GOOS, runtime.GOARCH)
	for _, dir := range directories {
		path := filepath.Join(dir, ".dodo", "plugins", filename)
		if stat, err := os.Stat(path); err == nil && stat.Mode().Perm()&0111 != 0 {
			return path, nil
		}
	}

	return "", errors.New("could not find a suitable plugin for the stage anywhere")
}

func (s *Stage) Cleanup() {
	if s.wrapped != nil {
		s.wrapped.Cleanup()
	}
	if s.client != nil {
		s.client.Kill()
	}
}

func (s *Stage) Create() error {
	return s.wrapped.Create()
}

func (s *Stage) Start() error {
	return s.wrapped.Start()
}

func (s *Stage) Stop() error {
	return s.wrapped.Stop()
}

func (s *Stage) Remove(force bool, volumes bool) error {
	return s.wrapped.Remove(force, volumes)
}

func (s *Stage) Exist() (bool, error) {
	return s.wrapped.Exist()
}

func (s *Stage) Available() (bool, error) {
	return s.wrapped.Available()
}

func (s *Stage) GetSSHOptions() (*stage.SSHOptions, error) {
	return s.wrapped.GetSSHOptions()
}

func (s *Stage) GetDockerOptions() (*stage.DockerOptions, error) {
	return s.wrapped.GetDockerOptions()
}
