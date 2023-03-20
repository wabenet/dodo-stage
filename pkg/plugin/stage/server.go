package stage

import (
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	core "github.com/wabenet/dodo-core/api/core/v1alpha5"
	stage "github.com/wabenet/dodo-stage/api/stage/v1alpha3"
	"golang.org/x/net/context"
)

type server struct {
	impl Stage
}

func (s *server) GetPluginInfo(_ context.Context, _ *empty.Empty) (*core.PluginInfo, error) {
	info := s.impl.PluginInfo()

	return &core.PluginInfo{
		Name: &core.PluginName{Name: info.Name.Name, Type: info.Name.Type},
	}, nil
}

func (s *server) InitPlugin(_ context.Context, _ *empty.Empty) (*core.InitPluginResponse, error) {
	config, err := s.impl.Init()
	if err != nil {
		return nil, fmt.Errorf("could not initialize plugin: %w", err)
	}

	return &core.InitPluginResponse{Config: config}, nil
}

func (s *server) ResetPlugin(_ context.Context, _ *empty.Empty) (*empty.Empty, error) {
	s.impl.Cleanup()

	return &empty.Empty{}, nil
}

func (s *server) GetStage(ctx context.Context, request *stage.GetStageRequest) (*stage.GetStageResponse, error) {
	return s.impl.GetStage(request.Name)
}

func (s *server) CreateStage(ctx context.Context, request *stage.CreateStageRequest) (*empty.Empty, error) {
	return &empty.Empty{}, s.impl.CreateStage(request.Config)
}

func (s *server) DeleteStage(ctx context.Context, request *stage.DeleteStageRequest) (*empty.Empty, error) {
	return &empty.Empty{}, s.impl.DeleteStage(request.Name, request.Force, request.Volumes)
}

func (s *server) StartStage(ctx context.Context, request *stage.StartStageRequest) (*empty.Empty, error) {
	return &empty.Empty{}, s.impl.StartStage(request.Name)
}

func (s *server) StopStage(ctx context.Context, request *stage.StopStageRequest) (*empty.Empty, error) {
	return &empty.Empty{}, s.impl.StopStage(request.Name)
}

func (s *server) ProvisionStage(ctx context.Context, request *stage.ProvisionStageRequest) (*empty.Empty, error) {
	return &empty.Empty{}, s.impl.ProvisionStage(request.Name)
}

func (s *server) GetProxy(ctx context.Context, request *stage.GetProxyRequest) (*stage.GetProxyResponse, error) {
	pc, err := s.impl.GetClient(request.Name)
	if err != nil {
		return nil, err
	}

	return &stage.GetProxyResponse{Config: pc.Config}, nil
}
