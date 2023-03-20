package main

import (
	"os"

	core "github.com/wabenet/dodo-core"
	plugin "github.com/wabenet/dodo-core/pkg/plugin"
	"github.com/wabenet/dodo-core/pkg/plugin/command"
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
