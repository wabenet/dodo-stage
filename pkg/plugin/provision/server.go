package provision

import (
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	core "github.com/wabenet/dodo-core/api/core/v1alpha5"
	provision "github.com/wabenet/dodo-stage/api/provision/v1alpha1"
	"golang.org/x/net/context"
)

type server struct {
	impl Provisioner
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

func (s *server) ProvisionStage(ctx context.Context, request *provision.ProvisionStageRequest) (*empty.Empty, error) {
	if err := s.impl.ProvisionStage(request.Stage, request.SshOptions); err != nil {
		return nil, fmt.Errorf("could not provision stage: %w", err)
	}

	return &empty.Empty{}, nil
}

func (s *server) CleanStage(ctx context.Context, request *provision.CleanStageRequest) (*empty.Empty, error) {
	if err := s.impl.CleanStage(request.Stage); err != nil {
		return nil, fmt.Errorf("could not cleanup stage: %w", err)
	}

	return &empty.Empty{}, nil
}

func (s *server) GetProxy(ctx context.Context, request *provision.GetProxyRequest) (*provision.GetProxyResponse, error) {
	pc, err := s.impl.GetClient(request.Stage)
	if err != nil {
		return nil, err
	}

	return &provision.GetProxyResponse{Config: pc.Config}, nil
}
