package main

import (
	"os"

	"github.com/wabenet/dodo-stage/pkg/plugin"
)

func main() {
	os.Exit(plugin.RunMe())
}
