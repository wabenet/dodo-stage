package plugin

import (
	"os"

	"github.com/dodo-cli/dodo-core/pkg/plugin"
	"github.com/dodo-cli/dodo-stage/pkg/command"
	"github.com/dodo-cli/dodo-stage/pkg/plugin/builder"
	"github.com/dodo-cli/dodo-stage/pkg/plugin/runtime"
	"github.com/dodo-cli/dodo-stage/pkg/plugin/stage"
)

func RunMe() int {
	m := plugin.Init()

	if os.Getenv(plugin.MagicCookieKey) == plugin.MagicCookieValue {
		plugins := []plugin.Plugin{}
		plugins = append(plugins, runtime.GetAllRuntimePlugins(m)...)
		plugins = append(plugins, builder.GetAllBuilderPlugins(m)...)

		m.ServePlugins(plugins...)

		return 0
	} else {
		if err := command.New(m).GetCobraCommand().Execute(); err != nil {
			return 1
		}

		return 0
	}
}

func IncludeMe(m plugin.Manager) {
	plugins := []plugin.Plugin{command.New(m)}
	plugins = append(plugins, runtime.GetAllRuntimePlugins(m)...)
	plugins = append(plugins, builder.GetAllBuilderPlugins(m)...)

	m.RegisterPluginTypes(stage.Type)
	m.IncludePlugins(plugins...)
}
