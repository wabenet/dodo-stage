package plugin

import (
	"os"

	"github.com/dodo-cli/dodo-core/pkg/appconfig"
	dodo "github.com/dodo-cli/dodo-core/pkg/plugin"
	"github.com/dodo-cli/dodo-stage/pkg/command"
	log "github.com/hashicorp/go-hclog"
)

func RunMe() int {
	if os.Getenv(dodo.MagicCookieKey) == dodo.MagicCookieValue {
		dodo.ServePlugins()
		return 0
	} else {
		log.SetDefault(log.New(appconfig.GetLoggerOptions()))
		p := &command.Command{}
		if err := p.Init(); err != nil {
			return 1
		}
		if err := p.GetCobraCommand().Execute(); err != nil {
			return 1
		}
		return 0
	}
}

func IncludeMe() {
	dodo.IncludePlugins(&command.Command{})
}
