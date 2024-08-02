package config

import (
	"cuelang.org/go/cue"
	"github.com/hashicorp/go-multierror"
	"github.com/wabenet/dodo-config/pkg/cuetils"
	api "github.com/wabenet/dodo-stage/api/stage/v1alpha4"
)

func SSHConfigFromValue(v cue.Value) (*api.SSHOptions, error) {
	var errs error

	if out, err := SSHConfigFromStruct(v); err == nil {
		return out, nil
	} else {
		errs = multierror.Append(errs, err)
	}

	return nil, errs
}

func SSHConfigFromStruct(v cue.Value) (*api.SSHOptions, error) {
	out := &api.SSHOptions{}

	if p, ok := cuetils.Get(v, "host"); ok {
		if n, err := StringFromValue(p); err != nil {
			return nil, err
		} else {
			out.Hostname = n
		}
	}

	if p, ok := cuetils.Get(v, "port"); ok {
		if n, err := IntFromValue(p); err != nil {
			return nil, err
		} else {
			out.Port = n
		}
	}

	if p, ok := cuetils.Get(v, "username"); ok {
		if n, err := StringFromValue(p); err != nil {
			return nil, err
		} else {
			out.Username = n
		}
	}

	if p, ok := cuetils.Get(v, "private_key_file"); ok {
		if n, err := StringFromValue(p); err != nil {
			return nil, err
		} else {
			out.PrivateKeyFile = n
		}
	}

	return out, nil
}
