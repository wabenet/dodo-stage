package provision

import (
	core "github.com/wabenet/dodo-core/api/core/v1alpha5"
	coreconfig "github.com/wabenet/dodo-core/pkg/config"
	"github.com/wabenet/dodo-core/pkg/plugin"
	api "github.com/wabenet/dodo-stage/api/provision/v1alpha1"
	stage "github.com/wabenet/dodo-stage/api/stage/v1alpha3"
	"github.com/wabenet/dodo-stage/internal/plugin/provision-fixed/config"
	"github.com/wabenet/dodo-stage/pkg/plugin/provision"
	"github.com/wabenet/dodo-stage/pkg/proxy"
)

const (
	name = "fixed"
)

var _ provision.Provisioner = &Provisioner{}

type Provisioner struct {
	proxyClient *proxy.Client
}

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

func (*Provisioner) ProvisionStage(*stage.StageInfo, *stage.SSHOptions) error {
	return nil
}

func (p *Provisioner) GetClient(info *stage.StageInfo) (*proxy.Client, error) {
	if p.proxyClient != nil {
		return p.proxyClient, nil
	}

	stages, err := config.GetAllStages(coreconfig.GetConfigFiles()...)
	if err != nil {
		return nil, err
	}
	cfg := stages[info.Name].Provision

	pc, err := proxy.NewClient(&api.ProxyConfig{
		Url:      cfg.Address,
		CaPath:   cfg.CaPath,
		CertPath: cfg.CertPath,
		KeyPath:  cfg.KeyPath,
	})
	if err != nil {
		return nil, err
	}

	p.proxyClient = pc

	return p.proxyClient, nil
}

func (*Provisioner) CleanStage(*stage.StageInfo) error {
	return nil
}
