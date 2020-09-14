package main

import (
	"os"

	"github.com/dodo-cli/dodo-stage/plugin"
)

func main() {
	os.Exit(plugin.RunMe())
}
