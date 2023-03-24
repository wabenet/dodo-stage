package config

import (
	"cuelang.org/go/cue"
	"github.com/wabenet/dodo-config/pkg/cuetils"
	api "github.com/wabenet/dodo-stage/api/stage/v1alpha3"
)

func StagesFromValue(v cue.Value) (map[string]*api.Stage, error) {
	return StagesFromMap(v)
}

func StagesFromMap(v cue.Value) (map[string]*api.Stage, error) {
	out := map[string]*api.Stage{}

	err := cuetils.IterMap(v, func(name string, v cue.Value) error {
		r, err := StageFromStruct(name, v)
		if err == nil {
			out[name] = r
		}

		return err

	})

	return out, err
}

func StageFromStruct(name string, v cue.Value) (*api.Stage, error) {
	out := &api.Stage{Name: name}

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

	return out, nil
}
