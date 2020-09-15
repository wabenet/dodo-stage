package stage

import (
	dodo "github.com/dodo-cli/dodo-core/pkg/plugin"
	"github.com/dodo-cli/dodo-stage/pkg/types"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const Type pluginType = "stage"

type pluginType string

func (t pluginType) String() string {
	return string(t)
}

func (t pluginType) GRPCClient() (plugin.Plugin, error) {
	return &grpcPlugin{}, nil
}

func (t pluginType) GRPCServer(p dodo.Plugin) (plugin.Plugin, error) {
	config, ok := p.(Stage)
	if !ok {
		return nil, dodo.ErrPluginInvalid
	}
	return &grpcPlugin{Impl: config}, nil
}

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

type grpcPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	Impl Stage
}

func (p *grpcPlugin) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &client{stageClient: types.NewDockerStageClient(conn)}, nil
}

func (p *grpcPlugin) GRPCServer(_ *plugin.GRPCBroker, s *grpc.Server) error {
	types.RegisterDockerStageServer(s, &server{impl: p.Impl})
	return nil
}
