package provision

import (
	"github.com/hashicorp/go-plugin"
	dodo "github.com/wabenet/dodo-core/pkg/plugin"
	provision "github.com/wabenet/dodo-stage/api/provision/v1alpha2"
	stage "github.com/wabenet/dodo-stage/api/stage/v1alpha4"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const Type pluginType = "stage-provisioner"

type pluginType string

func (t pluginType) String() string {
	return string(t)
}

func (t pluginType) GRPCClient() (plugin.Plugin, error) {
	return &grpcPlugin{}, nil
}

func (t pluginType) GRPCServer(p dodo.Plugin) (plugin.Plugin, error) {
	config, ok := p.(Provisioner)
	if !ok {
		return nil, dodo.InvalidError{
			Plugin:  p.PluginInfo().Name,
			Message: "plugin does not implement Stage Provision API",
		}
	}
	return &grpcPlugin{Impl: config}, nil
}

type grpcPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	Impl Provisioner
}

func (p *grpcPlugin) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &client{provisionClient: provision.NewPluginClient(conn)}, nil
}

func (p *grpcPlugin) GRPCServer(_ *plugin.GRPCBroker, s *grpc.Server) error {
	provision.RegisterPluginServer(s, &server{impl: p.Impl})
	return nil
}

type Provisioner interface {
	dodo.Plugin

	ProvisionStage(name string, sshOptions *stage.SSHOptions) error
	CleanStage(name string) error
}
