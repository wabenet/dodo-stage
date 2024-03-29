package plugin

import (
	"os"

	"github.com/wabenet/dodo-core/pkg/plugin"
	"github.com/wabenet/dodo-stage/internal/plugin/builder"
	"github.com/wabenet/dodo-stage/internal/plugin/command"
	stagehand "github.com/wabenet/dodo-stage/internal/plugin/provision-stagehand"
	"github.com/wabenet/dodo-stage/internal/plugin/runtime"
	"github.com/wabenet/dodo-stage/pkg/plugin/provision"
	"github.com/wabenet/dodo-stage/pkg/plugin/stage"
)

func RunMe() int {
	m := plugin.Init()

	if os.Getenv(plugin.MagicCookieKey) == plugin.MagicCookieValue {
		plugins := []plugin.Plugin{}
		plugins = append(plugins, stagehand.New())
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
	plugins = append(plugins, stagehand.New())
	plugins = append(plugins, runtime.GetAllRuntimePlugins(m)...)
	plugins = append(plugins, builder.GetAllBuilderPlugins(m)...)

	m.RegisterPluginTypes(stage.Type, provision.Type)
	m.IncludePlugins(plugins...)
}
