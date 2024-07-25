package script

import (
	"log"
	"os/exec"
)

const (
	Type = "run-script"
)

type Action struct {
	Script []string `mapstructure:"script"`
}

func (a *Action) Type() string {
	return Type
}

func (a *Action) Execute() error {
	log.Printf("running provision script...")

	for _, script := range a.Script {
		if _, err := exec.Command("/bin/sh", "-c", script).CombinedOutput(); err != nil {
			return err
		}
	}

	return nil
}
