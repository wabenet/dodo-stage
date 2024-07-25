package config

import (
	"cuelang.org/go/cue"
	"github.com/wabenet/dodo-config/pkg/cuetils"
)

type Stage struct {
	Name      string
	Type      string
	Provision *Provision
}

func StagesFromValue(v cue.Value) (map[string]*Stage, error) {
	return StagesFromMap(v)
}

func StagesFromMap(v cue.Value) (map[string]*Stage, error) {
	out := map[string]*Stage{}

	err := cuetils.IterMap(v, func(name string, v cue.Value) error {
		r, err := StageFromStruct(name, v)
		if err == nil {
			out[name] = r
		}

		return err

	})

	return out, err
}

func StageFromStruct(name string, v cue.Value) (*Stage, error) {
	out := &Stage{Name: name}

	if p, ok := cuetils.Get(v, "name"); ok {
		if n, err := StringFromValue(p); err != nil {
			return nil, err
		} else {
			out.Name = n
		}
	}

	if p, ok := cuetils.Get(v, "type"); ok {
		if n, err := StringFromValue(p); err != nil {
			return nil, err
		} else {
			out.Type = n
		}
	}

	if p, ok := cuetils.Get(v, "provision"); ok {
		if n, err := ProvisionFromValue(p); err != nil {
			return nil, err
		} else {
			out.Provision = n
		}
	}

	return out, nil
}
