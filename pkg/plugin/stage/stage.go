package stage

import (
	"github.com/hashicorp/go-plugin"
	dodo "github.com/wabenet/dodo-core/pkg/plugin"
	"github.com/wabenet/dodo-core/pkg/plugin/builder"
	"github.com/wabenet/dodo-core/pkg/plugin/runtime"
	api "github.com/wabenet/dodo-stage/api/v1alpha2"
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
		return nil, dodo.InvalidError{
			Plugin:  p.PluginInfo().Name,
			Message: "plugin does not implement Stage API",
		}
	}
	return &grpcPlugin{Impl: config}, nil
}

type grpcPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	Impl Stage
}

func (p *grpcPlugin) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &client{stageClient: api.NewStagePluginClient(conn)}, nil
}

func (p *grpcPlugin) GRPCServer(_ *plugin.GRPCBroker, s *grpc.Server) error {
	api.RegisterStagePluginServer(s, &server{impl: p.Impl})
	return nil
}

type Stage interface {
	dodo.Plugin

	GetStage(string) (*api.GetStageResponse, error)
	CreateStage(*api.Stage) error
	DeleteStage(string, bool, bool) error
	StartStage(string) error
	StopStage(string) error
	ProvisionStage(string) error
	GetContainerRuntime(string) (runtime.ContainerRuntime, error)
	GetImageBuilder(string) (builder.ImageBuilder, error)
}
