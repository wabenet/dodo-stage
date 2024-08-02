package config

import (
	"cuelang.org/go/cue"
	"github.com/wabenet/dodo-config/pkg/cuetils"
	api "github.com/wabenet/dodo-stage/api/stage/v1alpha4"
)

type Stage struct {
	Name       string
	Type       string
	Address    string
	CaPath     string
	CertPath   string
	KeyPath    string
	SSHOptions *api.SSHOptions
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

	if p, ok := cuetils.Get(v, "address"); ok {
		if n, err := StringFromValue(p); err != nil {
			return nil, err
		} else {
			out.Address = n
		}
	}

	if p, ok := cuetils.Get(v, "ca_path"); ok {
		if n, err := StringFromValue(p); err != nil {
			return nil, err
		} else {
			out.CaPath = n
		}
	}

	if p, ok := cuetils.Get(v, "cert_path"); ok {
		if n, err := StringFromValue(p); err != nil {
			return nil, err
		} else {
			out.CertPath = n
		}
	}

	if p, ok := cuetils.Get(v, "key_path"); ok {
		if n, err := StringFromValue(p); err != nil {
			return nil, err
		} else {
			out.KeyPath = n
		}
	}

	if p, ok := cuetils.Get(v, "ssh_config"); ok {
		if n, err := SSHConfigFromValue(p); err != nil {
			return nil, err
		} else {
			out.SSHOptions = n
		}
	}

	return out, nil
}
