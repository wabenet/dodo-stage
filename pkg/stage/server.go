package stage

import (
	api "github.com/dodo-cli/dodo-stage/api/v1alpha1"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
)

type server struct {
	impl Stage
}

func (s *server) Init(_ context.Context, _ *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.impl.Init()
}

func (s *server) GetPluginInfo(_ context.Context, _ *empty.Empty) (*api.PluginInfo, error) {
	info, err := s.impl.PluginInfo()
	if err != nil {
		return nil, err
	}

	return &api.PluginInfo{
		Name:    info.Name,
		Version: info.Version,
	}, nil
}

func (server *server) ListStages(ctx context.Context, _ *empty.Empty) (*api.ListStagesResponse, error) {
	response, err := server.impl.ListStages()
	if err != nil {
		return nil, err
	}

	return &api.ListStagesResponse{Stages: response}, nil
}

func (server *server) GetStage(ctx context.Context, request *api.GetStageRequest) (*api.GetStageResponse, error) {
	return server.impl.GetStage(request.Name)
}

func (server *server) CreateStage(ctx context.Context, request *api.CreateStageRequest) (*empty.Empty, error) {
	return &empty.Empty{}, server.impl.CreateStage(request.Config)
}

func (server *server) DeleteStage(ctx context.Context, request *api.DeleteStageRequest) (*empty.Empty, error) {
	return &empty.Empty{}, server.impl.DeleteStage(request.Name, request.Force, request.Volumes)
}

func (server *server) StartStage(ctx context.Context, request *api.StartStageRequest) (*empty.Empty, error) {
	return &empty.Empty{}, server.impl.StartStage(request.Name)
}

func (server *server) StopStage(ctx context.Context, request *api.StopStageRequest) (*empty.Empty, error) {
	return &empty.Empty{}, server.impl.StopStage(request.Name)
}
