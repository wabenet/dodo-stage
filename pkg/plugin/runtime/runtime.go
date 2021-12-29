package runtime

import (
	"fmt"

	coreapi "github.com/dodo-cli/dodo-core/api/v1alpha2"
	"github.com/dodo-cli/dodo-core/pkg/plugin"
	"github.com/dodo-cli/dodo-core/pkg/plugin/runtime"
	api "github.com/dodo-cli/dodo-stage/api/v1alpha1"
	"github.com/dodo-cli/dodo-stage/pkg/config"
	"github.com/dodo-cli/dodo-stage/pkg/plugin/stage"
	log "github.com/hashicorp/go-hclog"
	"github.com/oclaussen/go-gimme/configfiles"
)

var _ runtime.ContainerRuntime = &ContainerRuntime{}

type ContainerRuntime struct {
	name    string
	config  *api.Stage
	manager plugin.Manager
}

func (c *ContainerRuntime) Type() plugin.Type {
	return runtime.Type
}

func (c *ContainerRuntime) PluginInfo() *coreapi.PluginInfo {
	return &coreapi.PluginInfo{
		Name: &coreapi.PluginName{
			Name: c.name,
			Type: runtime.Type.String(),
		},
		Dependencies: []*coreapi.PluginName{
			{Name: c.config.Type, Type: stage.Type.String()},
		},
	}
}

func (c *ContainerRuntime) Init() (plugin.PluginConfig, error) {
	r, err := c.get()
	if err != nil {
		return nil, err
	}

	return r.Init()
}

func (c *ContainerRuntime) ResolveImage(spec string) (string, error) {
	r, err := c.get()
	if err != nil {
		return "", err
	}

	return r.ResolveImage(spec)
}

func (c *ContainerRuntime) CreateContainer(config *coreapi.Backdrop, tty bool, stdio bool) (string, error) {
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

	s, err := p.GetContainerRuntime(c.name)
	if err != nil {
		return nil, err
	}

	return s, nil
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
