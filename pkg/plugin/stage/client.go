package stage

import (
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/hashicorp/go-hclog"
	core "github.com/wabenet/dodo-core/api/core/v1alpha5"
	"github.com/wabenet/dodo-core/pkg/plugin"
	stage "github.com/wabenet/dodo-stage/api/stage/v1alpha3"
	"github.com/wabenet/dodo-stage/pkg/proxy"
	"golang.org/x/net/context"
)

var _ Stage = &client{}

type client struct {
	stageClient stage.PluginClient
}

func (c *client) Type() plugin.Type {
	return Type
}

func (c *client) PluginInfo() *core.PluginInfo {
	info, err := c.stageClient.GetPluginInfo(context.Background(), &empty.Empty{})
	if err != nil {
		return &core.PluginInfo{
			Name:   &core.PluginName{Type: Type.String(), Name: plugin.FailedPlugin},
			Fields: map[string]string{"error": err.Error()},
		}
	}

	return &core.PluginInfo{
		Name: &core.PluginName{Name: info.Name.Name, Type: info.Name.Type},
	}
}

func (c *client) Init() (plugin.Config, error) {
	resp, err := c.stageClient.InitPlugin(context.Background(), &empty.Empty{})
	if err != nil {
		return nil, fmt.Errorf("could not initialize plugin: %w", err)
	}

	return resp.Config, nil
}

func (c *client) Cleanup() {
	_, err := c.stageClient.ResetPlugin(context.Background(), &empty.Empty{})
	if err != nil {
		log.L().Error("plugin reset error", "error", err)
	}
}

func (c *client) GetStage(name string) (*stage.GetStageResponse, error) {
	return c.stageClient.GetStage(context.Background(), &stage.GetStageRequest{Name: name})
}

func (c *client) CreateStage(name string) error {
	_, err := c.stageClient.CreateStage(context.Background(), &stage.CreateStageRequest{Name: name})

	return err
}

func (c *client) DeleteStage(name string, force bool, volumes bool) error {
	_, err := c.stageClient.DeleteStage(context.Background(), &stage.DeleteStageRequest{Name: name, Force: force, Volumes: volumes})

	return err
}

func (c *client) StartStage(name string) error {
	_, err := c.stageClient.StartStage(context.Background(), &stage.StartStageRequest{Name: name})

	return err
}

func (c *client) StopStage(name string) error {
	_, err := c.stageClient.StopStage(context.Background(), &stage.StopStageRequest{Name: name})

	return err
}

func (c *client) ProvisionStage(name string) error {
	_, err := c.stageClient.ProvisionStage(context.Background(), &stage.ProvisionStageRequest{Name: name})

	return err
}

func (c *client) GetClient(name string) (*proxy.Client, error) {
	resp, err := c.stageClient.GetProxy(context.Background(), &stage.GetProxyRequest{Name: name})
	if err != nil {
		return nil, err
	}

	return proxy.NewClient(resp.Config)
}
