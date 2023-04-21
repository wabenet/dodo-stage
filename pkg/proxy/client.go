package proxy

import (
	"context"
	"fmt"
	"net"

	"github.com/wabenet/dodo-core/pkg/plugin/builder"
	"github.com/wabenet/dodo-core/pkg/plugin/runtime"
	provision "github.com/wabenet/dodo-stage/api/provision/v1alpha1"
	"google.golang.org/grpc"
)

type Client struct {
	runtime.ContainerRuntime
	builder.ImageBuilder

	Config *provision.ProxyConfig
	conn   *grpc.ClientConn
}

func NewClient(c *provision.ProxyConfig) (*Client, error) {
	protocol, addr, err := DialOptions(c)
	if err != nil {
		return nil, fmt.Errorf("invalid connection config: %w", err)
	}

	creds, err := TLSClientOptions(c)
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
		Config:           c,
		conn:             conn,
		ContainerRuntime: runtime.NewGRPCClient(conn),
		ImageBuilder:     builder.NewGRPCClient(conn),
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}
