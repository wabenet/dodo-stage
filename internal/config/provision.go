package config

import (
	"cuelang.org/go/cue"
	"github.com/hashicorp/go-multierror"
	"github.com/wabenet/dodo-config/pkg/cuetils"
)

type Provision struct {
	StagehandURL string
	Script       []string
}

func ProvisionFromValue(v cue.Value) (*Provision, error) {
	var errs error

	if out, err := ProvisionFromStruct(v); err == nil {
		return out, err
	} else {
		errs = multierror.Append(errs, err)
	}

	return nil, errs
}

func ProvisionFromStruct(v cue.Value) (*Provision, error) {
	out := &Provision{}

	if p, ok := cuetils.Get(v, "stagehand_url"); ok {
		if n, err := StringFromValue(p); err != nil {
			return nil, err
		} else {
			out.StagehandURL = n
		}
	}

	if p, ok := cuetils.Get(v, "script"); ok {
		if n, err := StringListFromValue(p); err != nil {
			return nil, err
		} else {
			out.Script = n
		}
	}

	return out, nil
}
