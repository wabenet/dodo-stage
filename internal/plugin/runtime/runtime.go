package runtime

import (
	"fmt"
	"os"

	log "github.com/hashicorp/go-hclog"
	"github.com/oclaussen/go-gimme/configfiles"
	core "github.com/wabenet/dodo-core/api/core/v1alpha5"
	"github.com/wabenet/dodo-core/pkg/plugin"
	"github.com/wabenet/dodo-core/pkg/plugin/runtime"
	"github.com/wabenet/dodo-stage/internal/config"
	"github.com/wabenet/dodo-stage/pkg/plugin/stage"
)

var _ runtime.ContainerRuntime = &ContainerRuntime{}

type ContainerRuntime struct {
	name    string
	config  *config.Stage
	manager plugin.Manager
}

func (c *ContainerRuntime) Type() plugin.Type {
	return runtime.Type
}

func (c *ContainerRuntime) PluginInfo() *core.PluginInfo {
	return &core.PluginInfo{
		Name: &core.PluginName{
			Name: c.name,
			Type: runtime.Type.String(),
		},
		Dependencies: []*core.PluginName{
			{Name: c.config.Type, Type: stage.Type.String()},
		},
	}
}

func (c *ContainerRuntime) Init() (plugin.Config, error) {
	r, err := c.get()
	if err != nil {
		return nil, err
	}

	return r.Init()
}

func (c *ContainerRuntime) Cleanup() {
	r, err := c.get()
	if err != nil {
		log.L().Error("plugin reset error", "error", err)

		return
	}

	r.Init()
}

func (c *ContainerRuntime) ResolveImage(spec string) (string, error) {
	r, err := c.get()
	if err != nil {
		return "", err
	}

	return r.ResolveImage(spec)
}

func (c *ContainerRuntime) CreateContainer(config *core.Backdrop, tty bool, stdio bool) (string, error) {
	r, err := c.get()
	if err != nil {
		return "", err
	}

	return r.CreateContainer(config, tty, stdio)
}

func (c *ContainerRuntime) StartContainer(id string) error {
	r, err := c.get()
	if err != nil {
		return err
	}

	return r.StartContainer(id)
}

func (c *ContainerRuntime) DeleteContainer(id string) error {
	r, err := c.get()
	if err != nil {
		return err
	}

	return r.DeleteContainer(id)
}

func (c *ContainerRuntime) KillContainer(id string, s os.Signal) error {
	r, err := c.get()
	if err != nil {
		return err
	}

	return r.KillContainer(id, s)
}

func (c *ContainerRuntime) ResizeContainer(id string, height uint32, width uint32) error {
	r, err := c.get()
	if err != nil {
		return err
	}

	return r.ResizeContainer(id, height, width)
}

func (c *ContainerRuntime) StreamContainer(id string, stream *plugin.StreamConfig) (*runtime.Result, error) {
	r, err := c.get()
	if err != nil {
		return nil, err
	}

	return r.StreamContainer(id, stream)
}

func GetAllRuntimePlugins(m plugin.Manager) []plugin.Plugin {
	plugins := []plugin.Plugin{}
	filenames := []string{}

	configfiles.GimmeConfigFiles(&configfiles.Options{
		Name:                      "dodo",
		Extensions:                []string{"yaml", "yml", "json"},
		IncludeWorkingDirectories: true,
		Filter: func(configFile *configfiles.ConfigFile) bool {
			filenames = append(filenames, configFile.Path)
			return false
		},
	})

	stages, err := config.GetAllStages(filenames...)
	if err != nil {
		log.L().Error(err.Error())
	}

	for name, config := range stages {
		plugins = append(plugins, &ContainerRuntime{
			name:    name,
			config:  config,
			manager: m,
		})
	}

	return plugins
}

func (c *ContainerRuntime) get() (runtime.ContainerRuntime, error) {
	p, err := loadPlugin(c.manager, c.config.Type)
	if err != nil {
		return nil, err
	}

	s, err := p.GetClient(c.name)
	if err != nil {
		return nil, err
	}

	return s.ContainerRuntime, nil
}

func loadPlugin(m plugin.Manager, name string) (stage.Stage, error) {
	for _, p := range m.GetPlugins(stage.Type.String()) {
		s := p.(stage.Stage)
		if info := s.PluginInfo(); info.Name.Name == name {
			return s, nil
		}
	}

	return nil, fmt.Errorf("could not find any stage plugin for type '%s'", name)
}
