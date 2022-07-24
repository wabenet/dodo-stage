package proxy

import (
	"context"
	"fmt"
	"net"

	api "github.com/wabenet/dodo-core/api/v1alpha4"
	"github.com/wabenet/dodo-core/pkg/plugin/builder"
	"github.com/wabenet/dodo-core/pkg/plugin/runtime"
	"google.golang.org/grpc"
)

type Client struct {
	runtime.ContainerRuntime
	builder.ImageBuilder

	conn *grpc.ClientConn
}

func NewClient(c *Config) (*Client, error) {
	protocol, addr, err := c.DialOptions()
	if err != nil {
		return nil, fmt.Errorf("invalid connection config: %w", err)
	}

	creds, err := c.TLSClientOptions()
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(creds),
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			return net.Dial(protocol, addr)
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("could not connect to server: %w", err)
	}

	return &Client{
		conn:             conn,
		ContainerRuntime: runtime.NewGRPCClient(api.NewRuntimePluginClient(conn)),
		ImageBuilder:     builder.NewGRPCClient(api.NewBuilderPluginClient(conn)),
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}
