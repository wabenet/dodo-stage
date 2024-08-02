package provision

import (
	core "github.com/wabenet/dodo-core/api/core/v1alpha5"
	"github.com/wabenet/dodo-core/pkg/plugin"
	stage "github.com/wabenet/dodo-stage/api/stage/v1alpha4"
	"github.com/wabenet/dodo-stage/pkg/plugin/provision"
)

const (
	name = "none"
)

var _ provision.Provisioner = &Provisioner{}

type Provisioner struct{}

func New() *Provisioner {
	return &Provisioner{}
}

func (*Provisioner) Type() plugin.Type {
	return provision.Type
}

func (*Provisioner) PluginInfo() *core.PluginInfo {
	return &core.PluginInfo{
		Name: &core.PluginName{Name: name, Type: provision.Type.String()},
	}
}

func (*Provisioner) Init() (plugin.Config, error) {
	return map[string]string{}, nil
}

func (*Provisioner) Cleanup() {}

func (*Provisioner) ProvisionStage(string, *stage.SSHOptions) error {
	return nil
}

func (*Provisioner) CleanStage(string) error {
	return nil
}
