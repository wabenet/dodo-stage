package command

import (
	"github.com/spf13/cobra"
	core "github.com/wabenet/dodo-core/api/core/v1alpha5"
	"github.com/wabenet/dodo-core/pkg/plugin"
	"github.com/wabenet/dodo-core/pkg/plugin/command"
)

const name = "stage"

var _ command.Command = &Command{}

type Command struct {
	cmd *cobra.Command
}

func (p *Command) Type() plugin.Type {
	return command.Type
}

func (p *Command) PluginInfo() *core.PluginInfo {
	return &core.PluginInfo{
		Name: &core.PluginName{Name: name, Type: command.Type.String()},
	}
}

func (*Command) Init() (plugin.Config, error) {
	return map[string]string{}, nil
}

func (*Command) Cleanup() {}

func (p *Command) GetCobraCommand() *cobra.Command {
	return p.cmd
}
