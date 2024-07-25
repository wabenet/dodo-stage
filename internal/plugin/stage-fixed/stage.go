package stage

import (
	log "github.com/hashicorp/go-hclog"
	coreapi "github.com/wabenet/dodo-core/api/core/v1alpha5"
	coreconfig "github.com/wabenet/dodo-core/pkg/config"
	"github.com/wabenet/dodo-core/pkg/plugin"
	api "github.com/wabenet/dodo-stage/api/stage/v1alpha3"
	"github.com/wabenet/dodo-stage/internal/plugin/stage-fixed/config"
	"github.com/wabenet/dodo-stage/pkg/plugin/stage"
)

const (
	name = "fixed"
)

var _ stage.Stage = &Stage{}

type Stage struct{}

func New() *Stage {
	return &Stage{}
}

func (*Stage) Type() plugin.Type {
	return stage.Type
}

func (s *Stage) PluginInfo() *coreapi.PluginInfo {
	return &coreapi.PluginInfo{
		Name: &coreapi.PluginName{Name: name, Type: stage.Type.String()},
	}
}

func (*Stage) Init() (plugin.Config, error) {
	return map[string]string{}, nil
}

func (s *Stage) Cleanup() {}

func (s *Stage) GetStage(name string) (*api.GetStageResponse, error) {
	resp := &api.GetStageResponse{
		Info: &api.StageInfo{
			Name:   name,
			Status: api.StageStatus_NONE,
		},
	}

	stages, err := config.GetAllStages(coreconfig.GetConfigFiles()...)
	if err != nil {
		return resp, err
	}

	if _, ok := stages[name]; ok {
		resp.Info.Status = api.StageStatus_UP
	}

	return resp, nil
}

func (s *Stage) CreateStage(name string) error {
	log.L().Info("Assuming remote stage already exists", "name", name)

	return nil
}

func (s *Stage) StartStage(name string) error {
	log.L().Info("Assuming remote state already running", "name", name)

	return nil
}

func (s *Stage) StopStage(name string) error {
	log.L().Info("Will not stop unmanaged remote stage", "name", name)

	return nil
}

func (s *Stage) DeleteStage(name string, force bool, _ bool) error {
	log.L().Info("Will not delete unmanaged remote stage", "name", name)

	return nil
}
