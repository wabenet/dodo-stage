package proxy

import (
	"fmt"
	"net"

	coreapi "github.com/wabenet/dodo-core/api/v1alpha4"
	"github.com/wabenet/dodo-core/pkg/plugin"
	"github.com/wabenet/dodo-core/pkg/plugin/builder"
	"github.com/wabenet/dodo-core/pkg/plugin/runtime"
	api "github.com/wabenet/dodo-stage/api/v1alpha2"
	"google.golang.org/grpc"
)

type Server struct {
	listener net.Listener
	server   *grpc.Server
	plugins  plugin.Manager
}

func NewServer(m plugin.Manager, c *api.ProxyConfig) (*Server, error) {
	protocol, addr, err := DialOptions(c)
	if err != nil {
		return nil, fmt.Errorf("invalid connection config: %w", err)
	}

	if _, err = net.Dial(protocol, addr); err == nil {
		return nil, fmt.Errorf("server already exists at %s: %w", addr, err)
	}

	creds, err := TLSServerOptions(c)
	if err != nil {
		return nil, err
	}

	listener, err := net.Listen(protocol, addr)
	if err != nil {
		return nil, fmt.Errorf("could not start server socket: %w", err)
	}

	return &Server{
		plugins:  m,
		listener: listener,
		server:   grpc.NewServer(grpc.Creds(creds)),
	}, nil
}

func (s *Server) Listen() error {
	defer s.listener.Close()

	if rt, err := runtime.GetByName(s.plugins, ""); err == nil {
		coreapi.RegisterRuntimePluginServer(s.server, runtime.NewGRPCServer(rt))
	}

	if b, err := builder.GetByName(s.plugins, ""); err == nil {
		coreapi.RegisterBuilderPluginServer(s.server, builder.NewGRPCServer(b))
	}

	return s.server.Serve(s.listener)
}
