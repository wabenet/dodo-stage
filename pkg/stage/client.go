package stage

import (
	"errors"

	coreapi "github.com/dodo-cli/dodo-core/api/v1alpha1"
	"github.com/dodo-cli/dodo-core/pkg/plugin"
	"github.com/dodo-cli/dodo-core/pkg/plugin/runtime"
	api "github.com/dodo-cli/dodo-stage/api/v1alpha1"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
)

var _ Stage = &client{}

type client struct {
	stageClient api.StagePluginClient
}

func (c *client) Type() plugin.Type {
	return Type
}

func (c *client) Init() error {
	_, err := c.stageClient.Init(context.Background(), &empty.Empty{})
	return err
}

func (c *client) PluginInfo() (*coreapi.PluginInfo, error) {
	info, err := c.stageClient.GetPluginInfo(context.Background(), &empty.Empty{})
	if err != nil {
		return nil, err
	}

	return &coreapi.PluginInfo{
		Name:    info.Name,
		Version: info.Version,
	}, nil
}

func (client *client) ListStages() ([]*api.Stage, error) {
	response, err := client.stageClient.ListStages(context.Background(), &empty.Empty{})
	if err != nil {
		return []*api.Stage{}, err
	}

	return response.Stages, nil
}

func (client *client) GetStage(name string) (*api.GetStageResponse, error) {
	return client.stageClient.GetStage(context.Background(), &api.GetStageRequest{Name: name})
}

func (client *client) CreateStage(config *api.Stage) error {
	_, err := client.stageClient.CreateStage(context.Background(), &api.CreateStageRequest{Config: config})
	return err
}

func (client *client) DeleteStage(name string, force bool, volumes bool) error {
	_, err := client.stageClient.DeleteStage(context.Background(), &api.DeleteStageRequest{Name: name, Force: force, Volumes: volumes})
	return err
}

func (client *client) StartStage(name string) error {
	_, err := client.stageClient.StartStage(context.Background(), &api.StartStageRequest{Name: name})
	return err
}

func (client *client) StopStage(name string) error {
	_, err := client.stageClient.StopStage(context.Background(), &api.StopStageRequest{Name: name})
	return err
}

func (client *client) GetContainerRuntime(name string) (runtime.ContainerRuntime, error) {
	return nil, errors.New("container runtime over grpc not implemented")
}
