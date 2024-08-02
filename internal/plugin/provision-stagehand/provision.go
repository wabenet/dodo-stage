package provision

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/hashicorp/go-hclog"
	core "github.com/wabenet/dodo-core/api/core/v1alpha5"
	coreconfig "github.com/wabenet/dodo-core/pkg/config"
	"github.com/wabenet/dodo-core/pkg/plugin"
	stage "github.com/wabenet/dodo-stage/api/stage/v1alpha4"
	"github.com/wabenet/dodo-stage/internal/plugin/provision-stagehand/config"
	"github.com/wabenet/dodo-stage/pkg/plugin/provision"
	"github.com/wabenet/dodo-stage/pkg/proxy"
	"github.com/wabenet/dodo-stage/pkg/stagehand"
	"github.com/wabenet/dodo-stage/pkg/stagehand/installer"
)

const (
	name        = "stagehand"
	DefaultPort = 20257

	permPrivateDir  = 0o700
	permPrivateFile = 0o600
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

func (p *Provisioner) ProvisionStage(name string, sshOpts *stage.SSHOptions) error {
	stages, err := config.GetAllStages(coreconfig.GetConfigFiles()...)
	if err != nil {
		return err
	}

	cfg := stages[name]
	inst := installer.SSHInstaller{
		DownloadUrl: cfg.Provision.StagehandURL,
		SSHOptions:  sshOpts,
	}

	// TODO: Allow extra options (e.g. replace ssh key?)
	provisionConfig := &stagehand.Config{
		Hostname:    name,
		Script:      cfg.Provision.Script,
		DefaultUser: sshOpts.Username,
	}

	result, err := inst.Install(provisionConfig)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(storagePath(name), permPrivateDir); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(storagePath(name), "ca.pem"), []byte(result.CA), permPrivateFile); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(storagePath(name), "client.pem"), []byte(result.ClientCert), permPrivateFile); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(storagePath(name), "client-key.pem"), []byte(result.ClientKey), permPrivateFile); err != nil {
		return err
	}

	if client, err := proxy.NewClient(&stage.ProxyConfig{
		Url:      fmt.Sprintf("tcp://%s:%d", result.IPAddress, DefaultPort),
		CaPath:   filepath.Join(storagePath(name), "ca.pem"),
		CertPath: filepath.Join(storagePath(name), "client.pem"),
		KeyPath:  filepath.Join(storagePath(name), "client-key.pem"),
	}); err != nil {
		return err
	} else {
		defer client.Close()
	}

	log.L().Info("stage is fully provisioned")

	return nil
}

func (p *Provisioner) CleanStage(name string) error {
	if err := os.Remove(filepath.Join(storagePath(name), "ca.pem")); err != nil {
		return err
	}

	if err := os.Remove(filepath.Join(storagePath(name), "client.pem")); err != nil {
		return err
	}

	if err := os.Remove(filepath.Join(storagePath(name), "client-key.pem")); err != nil {
		return err
	}

	return nil
}

func storagePath(name string) string {
	return filepath.Join(coreconfig.GetAppDir(), "stages", name)
}
