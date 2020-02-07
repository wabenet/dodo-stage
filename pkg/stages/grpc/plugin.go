package grpc

import (
	"github.com/dodo/dodo-stage/pkg/stage"
	"github.com/dodo/dodo-stage/proto"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const ProtocolVersion = 1

var PluginMap = map[string]plugin.Plugin{
	"stage": &Plugin{},
}

func HandshakeConfig(stageType string) plugin.HandshakeConfig {
	return plugin.HandshakeConfig{
		ProtocolVersion:  ProtocolVersion,
		MagicCookieKey:   "DODO_STAGE",
		MagicCookieValue: stageType,
	}
}

type Plugin struct {
	plugin.NetRPCUnsupportedPlugin
	Impl stage.Stage
}

func (p *Plugin) GRPCServer(_ *plugin.GRPCBroker, server *grpc.Server) error {
	proto.RegisterStageServer(server, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *Plugin) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, client *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewStageClient(client)}, nil
}
