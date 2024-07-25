package builder

import (
	"fmt"

	log "github.com/hashicorp/go-hclog"
	"github.com/oclaussen/go-gimme/configfiles"
	core "github.com/wabenet/dodo-core/api/core/v1alpha5"
	"github.com/wabenet/dodo-core/pkg/plugin"
	"github.com/wabenet/dodo-core/pkg/plugin/builder"
	"github.com/wabenet/dodo-stage/internal/plugin/builder/config"
	"github.com/wabenet/dodo-stage/pkg/plugin/provision"
	"github.com/wabenet/dodo-stage/pkg/plugin/stage"
)

var _ builder.ImageBuilder = &ImageBuilder{}

type ImageBuilder struct {
	name    string
	config  *config.Stage
	manager plugin.Manager
}

func (b *ImageBuilder) Type() plugin.Type {
	return builder.Type
}

func (b *ImageBuilder) PluginInfo() *core.PluginInfo {
	return &core.PluginInfo{
		Name: &core.PluginName{
			Name: b.name,
			Type: builder.Type.String(),
		},
		Dependencies: []*core.PluginName{
			{Name: "stagehand", Type: provision.Type.String()},
		},
	}
}

func (b *ImageBuilder) Init() (plugin.Config, error) {
	i, err := b.get()
	if err != nil {
		return nil, err
	}

	return i.Init()
}

func (b *ImageBuilder) Cleanup() {
	i, err := b.get()
	if err != nil {
		log.L().Error("plugin reset error", "error", err)

		return
	}

	i.Init()
}

func (b *ImageBuilder) CreateImage(config *core.BuildInfo, stream *plugin.StreamConfig) (string, error) {
	i, err := b.get()
	if err != nil {
		return "", err
	}

	return i.CreateImage(config, stream)
}

func GetAllBuilderPlugins(m plugin.Manager) []plugin.Plugin {
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
		plugins = append(plugins, &ImageBuilder{
			name:    name,
			config:  config,
			manager: m,
		})
	}

	return plugins
}

func (b *ImageBuilder) get() (builder.ImageBuilder, error) {
	s, err := loadStagePlugin(b.manager, b.config.Type)
	if err != nil {
		return nil, err
	}

	p, err := loadProvisionPlugin(b.manager, b.config.Provision.Type)
	if err != nil {
		return nil, err
	}

	status, err := s.GetStage(b.name)
	if err != nil {
		return nil, err
	}

	client, err := p.GetClient(status.Info)
	if err != nil {
		return nil, err
	}

	return client.ImageBuilder, nil
}

func loadStagePlugin(m plugin.Manager, name string) (stage.Stage, error) {
	for _, p := range m.GetPlugins(stage.Type.String()) {
		s := p.(stage.Stage)
		if info := s.PluginInfo(); info.Name.Name == name {
			return s, nil
		}
	}

	return nil, fmt.Errorf("could not find any stage plugin for type '%s'", name)
}

func loadProvisionPlugin(m plugin.Manager, name string) (provision.Provisioner, error) {
	for _, p := range m.GetPlugins(provision.Type.String()) {
		s := p.(provision.Provisioner)
		if info := s.PluginInfo(); info.Name.Name == name {
			return s, nil
		}
	}

	return nil, fmt.Errorf("could not find any stage provisioner plugin for type '%s'", name)
}
