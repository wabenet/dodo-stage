package stage

import (
	"encoding/json"

	"github.com/dodo-cli/dodo-stage/pkg/types"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type client struct {
	stageClient types.DockerStageClient
}

func (client *client) Initialize(name string, config *types.Stage) error {
	jsonBytes, err := json.Marshal(config)
	if err != nil {
		return errors.Wrap(err, "could not marshal json")
	}
	_, err = client.stageClient.Initialize(context.Background(), &types.InitRequest{Name: name, Config: string(jsonBytes)})
	return err
}

func (client *client) Cleanup() {}

func (client *client) Create() error {
	_, err := client.stageClient.Create(context.Background(), &types.Empty{})
	return err
}

func (client *client) Remove(force bool, volumes bool) error {
	_, err := client.stageClient.Remove(context.Background(), &types.RemoveRequest{Force: force, Volumes: volumes})
	return err
}

func (client *client) Start() error {
	_, err := client.stageClient.Start(context.Background(), &types.Empty{})
	return err
}

func (client *client) Stop() error {
	_, err := client.stageClient.Stop(context.Background(), &types.Empty{})
	return err
}

func (client *client) Exist() (bool, error) {
	response, err := client.stageClient.Exist(context.Background(), &types.Empty{})
	if err != nil {
		return false, err
	}
	return response.Exist, nil
}

func (client *client) Available() (bool, error) {
	response, err := client.stageClient.Available(context.Background(), &types.Empty{})
	if err != nil {
		return false, err
	}
	return response.Available, nil
}

func (client *client) GetSSHOptions() (*SSHOptions, error) {
	response, err := client.stageClient.GetSSHOptions(context.Background(), &types.Empty{})
	if err != nil {
		return nil, err
	}
	return &SSHOptions{
		Hostname:       response.Hostname,
		Port:           int(response.Port),
		Username:       response.Username,
		PrivateKeyFile: response.PrivateKeyFile,
	}, nil
}

func (client *client) GetDockerOptions() (*DockerOptions, error) {
	response, err := client.stageClient.GetDockerOptions(context.Background(), &types.Empty{})
	if err != nil {
		return nil, err
	}
	return &DockerOptions{
		Version:  response.Version,
		Host:     response.Host,
		CAFile:   response.CaFile,
		CertFile: response.CertFile,
		KeyFile:  response.KeyFile,
	}, nil
}
