package proxy

import (
	"fmt"
	"net"

	buildapi "github.com/wabenet/dodo-core/api/build/v1alpha1"
	runtimeapi "github.com/wabenet/dodo-core/api/runtime/v1alpha1"
	"github.com/wabenet/dodo-core/pkg/plugin"
	"github.com/wabenet/dodo-core/pkg/plugin/builder"
	"github.com/wabenet/dodo-core/pkg/plugin/runtime"
	provision "github.com/wabenet/dodo-stage/api/provision/v1alpha1"
	"google.golang.org/grpc"
)

type Server struct {
	listener net.Listener
	server   *grpc.Server
	plugins  plugin.Manager
}

func NewServer(m plugin.Manager, c *provision.ProxyConfig) (*Server, error) {
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
		runtimeapi.RegisterPluginServer(s.server, runtime.NewGRPCServer(rt))
	}

	if b, err := builder.GetByName(s.plugins, ""); err == nil {
		buildapi.RegisterPluginServer(s.server, builder.NewGRPCServer(b))
	}

	return s.server.Serve(s.listener)
}
