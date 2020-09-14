package grpc

import (
	"encoding/json"

	"github.com/dodo-cli/dodo-stage/pkg/stage"
	"github.com/dodo-cli/dodo-stage/pkg/types"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type GRPCServer struct {
	Impl stage.Stage
}

func (server *GRPCServer) Initialize(ctx context.Context, request *types.InitRequest) (*types.Empty, error) {
	var config types.Stage
	if err := json.Unmarshal([]byte(request.Config), &config); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal json")
	}
	return &types.Empty{}, server.Impl.Initialize(request.Name, &config)
}

func (server *GRPCServer) Create(ctx context.Context, _ *types.Empty) (*types.Empty, error) {
	return &types.Empty{}, server.Impl.Create()
}

func (server *GRPCServer) Remove(ctx context.Context, request *types.RemoveRequest) (*types.Empty, error) {
	return &types.Empty{}, server.Impl.Remove(request.Force, request.Volumes)
}

func (server *GRPCServer) Start(ctx context.Context, _ *types.Empty) (*types.Empty, error) {
	return &types.Empty{}, server.Impl.Start()
}

func (server *GRPCServer) Stop(ctx context.Context, _ *types.Empty) (*types.Empty, error) {
	return &types.Empty{}, server.Impl.Stop()
}

func (server *GRPCServer) Exist(ctx context.Context, _ *types.Empty) (*types.ExistResponse, error) {
	exist, err := server.Impl.Exist()
	if err != nil {
		return nil, err
	}
	return &types.ExistResponse{Exist: exist}, nil
}

func (server *GRPCServer) Available(ctx context.Context, _ *types.Empty) (*types.AvailableResponse, error) {
	available, err := server.Impl.Available()
	if err != nil {
		return nil, err
	}
	return &types.AvailableResponse{Available: available}, nil
}

func (server *GRPCServer) GetSSHOptions(ctx context.Context, _ *types.Empty) (*types.SSHOptionsResponse, error) {
	opts, err := server.Impl.GetSSHOptions()
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

func (server *GRPCServer) GetDockerOptions(ctx context.Context, _ *types.Empty) (*types.DockerOptionsResponse, error) {
	opts, err := server.Impl.GetDockerOptions()
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
