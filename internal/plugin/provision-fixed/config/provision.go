package config

import (
	"cuelang.org/go/cue"
	"github.com/hashicorp/go-multierror"
	"github.com/wabenet/dodo-config/pkg/cuetils"
)

type Provision struct {
	Type     string
	Address  string
	CaPath   string
	CertPath string
	KeyPath  string
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

	return out, nil
}
