package stage

import (
	"errors"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/hashicorp/go-hclog"
	coreapi "github.com/wabenet/dodo-core/api/v1alpha4"
	"github.com/wabenet/dodo-core/pkg/plugin"
	"github.com/wabenet/dodo-core/pkg/plugin/builder"
	"github.com/wabenet/dodo-core/pkg/plugin/runtime"
	api "github.com/wabenet/dodo-stage/api/v1alpha2"
	"golang.org/x/net/context"
)

var _ Stage = &client{}

type client struct {
	stageClient api.StagePluginClient
}

func (c *client) Type() plugin.Type {
	return Type
}

func (c *client) PluginInfo() *coreapi.PluginInfo {
	info, err := c.stageClient.GetPluginInfo(context.Background(), &empty.Empty{})
	if err != nil {
		return &coreapi.PluginInfo{
			Name:   &coreapi.PluginName{Type: Type.String(), Name: plugin.FailedPlugin},
			Fields: map[string]string{"error": err.Error()},
		}
	}

	return &coreapi.PluginInfo{
		Name: &coreapi.PluginName{Name: info.Name.Name, Type: info.Name.Type},
	}
}

func (c *client) Init() (plugin.PluginConfig, error) {
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

func (c *client) GetStage(name string) (*api.GetStageResponse, error) {
	return c.stageClient.GetStage(context.Background(), &api.GetStageRequest{Name: name})
}

func (c *client) CreateStage(config *api.Stage) error {
	_, err := c.stageClient.CreateStage(context.Background(), &api.CreateStageRequest{Config: config})

	return err
}

func (c *client) DeleteStage(name string, force bool, volumes bool) error {
	_, err := c.stageClient.DeleteStage(context.Background(), &api.DeleteStageRequest{Name: name, Force: force, Volumes: volumes})

	return err
}

func (c *client) StartStage(name string) error {
	_, err := c.stageClient.StartStage(context.Background(), &api.StartStageRequest{Name: name})

	return err
}

func (c *client) StopStage(name string) error {
	_, err := c.stageClient.StopStage(context.Background(), &api.StopStageRequest{Name: name})

	return err
}

func (c *client) ProvisionStage(name string) error {
	_, err := c.stageClient.ProvisionStage(context.Background(), &api.ProvisionStageRequest{Name: name})

	return err
}

func (c *client) GetContainerRuntime(name string) (runtime.ContainerRuntime, error) {
	return nil, errors.New("container runtime over grpc not implemented")
}

func (c *client) GetImageBuilder(name string) (builder.ImageBuilder, error) {
	return nil, errors.New("image builder over grpc not implemented")
}
