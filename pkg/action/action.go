package action

import (
	"errors"

	"github.com/mitchellh/mapstructure"
	"github.com/wabenet/dodo-stage/pkg/stagehand/docker"
	"github.com/wabenet/dodo-stage/pkg/stagehand/hostname"
	"github.com/wabenet/dodo-stage/pkg/stagehand/network"
	"github.com/wabenet/dodo-stage/pkg/stagehand/proxy"
	"github.com/wabenet/dodo-stage/pkg/stagehand/script"
	"github.com/wabenet/dodo-stage/pkg/stagehand/ssh"
)

type Action interface {
	Type() string
	Execute() error
}

type actionConfig struct {
	ID     string                 `mapstructure:"id"`
	Type   string                 `mapstructure:"type"`
	Config map[string]interface{} `mapstructure:",remain"`
}

func New(name string, config interface{}) (Action, error) {
	at, ac := decode(name, config)

	act, err := getByType(at)
	if err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(ac, &act); err != nil {
		return nil, err
	}

	return act, nil
}

func decode(key string, value interface{}) (string, map[string]interface{}) {
	ac := actionConfig{}
	if err := mapstructure.Decode(value, &ac); err != nil {
		return key, map[string]interface{}{
			"config": value,
		}
	}

	if t := ac.Type; t != "" {
		return t, ac.Config
	}

	if t := ac.ID; t != "" {
		return t, ac.Config
	}

	return key, ac.Config
}

func getByType(t string) (Action, error) {
	switch t {
	case docker.Type:
		return &docker.Action{}, nil
	case hostname.Type:
		return &hostname.Action{}, nil
	case network.Type:
		return &network.Action{}, nil
	case proxy.Type:
		return &proxy.Action{}, nil
	case script.Type:
		return &script.Action{}, nil
	case ssh.Type:
		return &ssh.Action{}, nil
	default:
		return nil, errors.New("Unknown action")
	}
}
