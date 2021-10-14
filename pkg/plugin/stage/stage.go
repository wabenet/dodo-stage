package stage

import (
	dodo "github.com/dodo-cli/dodo-core/pkg/plugin"
	"github.com/dodo-cli/dodo-core/pkg/plugin/builder"
	"github.com/dodo-cli/dodo-core/pkg/plugin/runtime"
	api "github.com/dodo-cli/dodo-stage/api/v1alpha1"
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
	GetContainerRuntime(string) (runtime.ContainerRuntime, error)
	GetImageBuilder(string) (builder.ImageBuilder, error)
}
