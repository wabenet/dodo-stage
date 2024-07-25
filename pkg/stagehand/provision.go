package stagehand

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/wabenet/dodo-stage/pkg/action"
	"github.com/wabenet/dodo-stage/pkg/stagehand/docker"
	"github.com/wabenet/dodo-stage/pkg/stagehand/hostname"
	"github.com/wabenet/dodo-stage/pkg/stagehand/network"
	"github.com/wabenet/dodo-stage/pkg/stagehand/proxy"
	"github.com/wabenet/dodo-stage/pkg/stagehand/script"
	"github.com/wabenet/dodo-stage/pkg/stagehand/ssh"
)

func Provision(input []byte) error {
	initial := map[string][]action.Action{
		docker.Type:   {},
		hostname.Type: {},
		network.Type:  {},
		proxy.Type:    {},
		script.Type:   {},
		ssh.Type:      {},
	}

	var cfg map[string]interface{}
	if err := json.Unmarshal(input, &cfg); err != nil {
		return fmt.Errorf("invalid yaml syntax: %w", err)
	}

	actionsByType, err := sortActions(cfg, initial)
	if err != nil {
		return err
	}

	// TODO: implement something smarter to put the actions in the correct order
	// This list is currently hardcoded, so we have the exact same behavior
	// as before
	sorted := []action.Action{}
	sorted = append(sorted, actionsByType[ssh.Type]...)
	sorted = append(sorted, actionsByType[network.Type]...)
	sorted = append(sorted, actionsByType[hostname.Type]...)
	sorted = append(sorted, actionsByType[script.Type]...)
	sorted = append(sorted, actionsByType[docker.Type]...)
	sorted = append(sorted, actionsByType[proxy.Type]...)

	for _, a := range sorted {
		if err := a.Execute(); err != nil {
			return err
		}
	}

	return nil
}

func sortActions(unsorted map[string]interface{}, actionsByType map[string][]action.Action) (map[string][]action.Action, error) {
	for name, value := range unsorted {
		if name == "actions" {
			subActions := map[string]interface{}{}
			if err := mapstructure.Decode(value, &subActions); err != nil {
				return nil, err
			}

			subSorted, err := sortActions(subActions, actionsByType)
			if err != nil {
				return nil, err
			}

			actionsByType = subSorted
		} else {
			act, err := action.New(name, value)
			if err != nil {
				return nil, fmt.Errorf("could not decode action: %w", err)
			}

			acts := actionsByType[act.Type()]
			acts = append(acts, act)
			actionsByType[act.Type()] = acts
		}
	}

	return actionsByType, nil
}
