package main

import (
	"os"

	plugin "github.com/wabenet/dodo-core/pkg/plugin"
	"github.com/wabenet/dodo-core/pkg/plugin/command"
	core "github.com/wabenet/dodo-core/plugin"
	stagehand "github.com/wabenet/dodo-stage/internal/plugin/command-stagehand"
)

const (
	ExitCodeInternalError = 1
)

func main() {
	os.Exit(execute())
}

func execute() int {
	m := plugin.Init()

	core.IncludeMe(m)

	m.LoadPlugins()

	defer m.UnloadPlugins()

	cmd := stagehand.New(m).GetCobraCommand()

	if err := cmd.Execute(); err != nil {
		return ExitCodeInternalError
	}

	return command.GetExitCode(cmd)
}
