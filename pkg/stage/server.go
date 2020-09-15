package stage

import (
	"encoding/json"

	"github.com/dodo-cli/dodo-stage/pkg/types"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type server struct {
	impl Stage
}

func (server *server) Initialize(ctx context.Context, request *types.InitRequest) (*types.Empty, error) {
	var config types.Stage
	if err := json.Unmarshal([]byte(request.Config), &config); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal json")
	}
	return &types.Empty{}, server.impl.Initialize(request.Name, &config)
}

func (server *server) Create(ctx context.Context, _ *types.Empty) (*types.Empty, error) {
	return &types.Empty{}, server.impl.Create()
}

func (server *server) Remove(ctx context.Context, request *types.RemoveRequest) (*types.Empty, error) {
	return &types.Empty{}, server.impl.Remove(request.Force, request.Volumes)
}

func (server *server) Start(ctx context.Context, _ *types.Empty) (*types.Empty, error) {
	return &types.Empty{}, server.impl.Start()
}

func (server *server) Stop(ctx context.Context, _ *types.Empty) (*types.Empty, error) {
	return &types.Empty{}, server.impl.Stop()
}

func (server *server) Exist(ctx context.Context, _ *types.Empty) (*types.ExistResponse, error) {
	exist, err := server.impl.Exist()
	if err != nil {
		return nil, err
	}
	return &types.ExistResponse{Exist: exist}, nil
}

func (server *server) Available(ctx context.Context, _ *types.Empty) (*types.AvailableResponse, error) {
	available, err := server.impl.Available()
	if err != nil {
		return nil, err
	}
	return &types.AvailableResponse{Available: available}, nil
}

func (server *server) GetSSHOptions(ctx context.Context, _ *types.Empty) (*types.SSHOptionsResponse, error) {
	opts, err := server.impl.GetSSHOptions()
	if err != nil {
		return nil, err
	}
	return &types.SSHOptionsResponse{
		Hostname:       opts.Hostname,
		Port:           int32(opts.Port),
		Username:       opts.Username,
		PrivateKeyFile: opts.PrivateKeyFile,
	}, nil
}

func (server *server) GetDockerOptions(ctx context.Context, _ *types.Empty) (*types.DockerOptionsResponse, error) {
	opts, err := server.impl.GetDockerOptions()
	if err != nil {
		return nil, err
	}
	return &types.DockerOptionsResponse{
		Version:  opts.Version,
		Host:     opts.Host,
		CaFile:   opts.CAFile,
		CertFile: opts.CertFile,
		KeyFile:  opts.KeyFile,
	}, nil
}
