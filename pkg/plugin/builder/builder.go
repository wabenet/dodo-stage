package builder

import (
	"fmt"

	coreapi "github.com/dodo-cli/dodo-core/api/v1alpha2"
	"github.com/dodo-cli/dodo-core/pkg/decoder"
	"github.com/dodo-cli/dodo-core/pkg/plugin"
	"github.com/dodo-cli/dodo-core/pkg/plugin/builder"
	api "github.com/dodo-cli/dodo-stage/api/v1alpha1"
	"github.com/dodo-cli/dodo-stage/pkg/plugin/stage"
	"github.com/dodo-cli/dodo-stage/pkg/types"
	"github.com/oclaussen/go-gimme/configfiles"
)

var _ builder.ImageBuilder = &ImageBuilder{}

type ImageBuilder struct {
	name    string
	config  *api.Stage
	manager plugin.Manager
}

func (b *ImageBuilder) Type() plugin.Type {
	return builder.Type
}

func (b *ImageBuilder) PluginInfo() *coreapi.PluginInfo {
	return &coreapi.PluginInfo{
		Name: &coreapi.PluginName{
			Name: b.name,
			Type: builder.Type.String(),
		},
		Dependencies: []*coreapi.PluginName{
			&coreapi.PluginName{
				Name: b.config.Type,
				Type: stage.Type.String(),
			},
		},
	}
}

func (b *ImageBuilder) Init() (plugin.PluginConfig, error) {
	i, err := b.get()
	if err != nil {
		return nil, err
	}

	return i.Init()
}

func (b *ImageBuilder) CreateImage(config *coreapi.BuildInfo, stream *plugin.StreamConfig) (string, error) {
	i, err := b.get()
	if err != nil {
		return "", err
	}

	return i.CreateImage(config, stream)
}

func GetAllBuilderPlugins(m plugin.Manager) []plugin.Plugin {
	plugins := []plugin.Plugin{}
	stages := map[string]*api.Stage{}

	configfiles.GimmeConfigFiles(&configfiles.Options{
		Name:                      "dodo",
		Extensions:                []string{"yaml", "yml", "json"},
		IncludeWorkingDirectories: true,
		Filter: func(configFile *configfiles.ConfigFile) bool {
			d := decoder.New(configFile.Path)
			d.DecodeYaml(configFile.Content, &stages, map[string]decoder.Decoding{
				"stages": decoder.Map(types.NewStage(), &stages),
			})

			return false
		},
	})

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
	p, err := loadPlugin(b.manager, b.config.Type)
	if err != nil {
		return nil, err
	}

	s, err := p.GetImageBuilder(b.name)
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
