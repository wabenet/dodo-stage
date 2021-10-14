package stage

import (
	"fmt"

	api "github.com/dodo-cli/dodo-stage/api/v1alpha1"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
)

type server struct {
	impl Stage
}

func (s *server) GetPluginInfo(_ context.Context, _ *empty.Empty) (*api.PluginInfo, error) {
	info := s.impl.PluginInfo()

	return &api.PluginInfo{
		Name: &api.PluginName{Name: info.Name.Name, Type: info.Name.Type},
	}, nil
}

func (s *server) InitPlugin(_ context.Context, _ *empty.Empty) (*api.InitPluginResponse, error) {
	config, err := s.impl.Init()
	if err != nil {
		return nil, fmt.Errorf("could not initialize plugin: %w", err)
	}

	return &api.InitPluginResponse{Config: config}, nil
}

func (s *server) GetStage(ctx context.Context, request *api.GetStageRequest) (*api.GetStageResponse, error) {
	return s.impl.GetStage(request.Name)
}

func (s *server) CreateStage(ctx context.Context, request *api.CreateStageRequest) (*empty.Empty, error) {
	return &empty.Empty{}, s.impl.CreateStage(request.Config)
}

func (s *server) DeleteStage(ctx context.Context, request *api.DeleteStageRequest) (*empty.Empty, error) {
	return &empty.Empty{}, s.impl.DeleteStage(request.Name, request.Force, request.Volumes)
}

func (s *server) StartStage(ctx context.Context, request *api.StartStageRequest) (*empty.Empty, error) {
	return &empty.Empty{}, s.impl.StartStage(request.Name)
}

func (s *server) StopStage(ctx context.Context, request *api.StopStageRequest) (*empty.Empty, error) {
	return &empty.Empty{}, s.impl.StopStage(request.Name)
}
