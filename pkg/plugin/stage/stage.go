package stage

import (
	"github.com/hashicorp/go-plugin"
	dodo "github.com/wabenet/dodo-core/pkg/plugin"
	stage "github.com/wabenet/dodo-stage/api/stage/v1alpha3"
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
	return &client{stageClient: stage.NewPluginClient(conn)}, nil
}

func (p *grpcPlugin) GRPCServer(_ *plugin.GRPCBroker, s *grpc.Server) error {
	stage.RegisterPluginServer(s, &server{impl: p.Impl})
	return nil
}

type Stage interface {
	dodo.Plugin

	GetStage(string) (*stage.GetStageResponse, error)
	CreateStage(string) error
	DeleteStage(string, bool, bool) error
	StartStage(string) error
	StopStage(string) error
}
