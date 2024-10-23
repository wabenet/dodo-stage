package config

import (
	"fmt"

	"cuelang.org/go/cue"
	"github.com/hashicorp/go-multierror"
	"github.com/wabenet/dodo-config/pkg/cuetils"
	"github.com/wabenet/dodo-config/pkg/includes"
	api "github.com/wabenet/dodo-stage/api/stage/v1alpha4"
	"github.com/wabenet/dodo-stage/pkg/spec"
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

func GetAllStages(filenames ...string) (map[string]*Stage, error) {
	var errs error
	stages := map[string]*Stage{}

	resolved, err := includes.ResolveIncludes(filenames...)
	if err != nil {
		errs = multierror.Append(errs, err)
		return stages, errs
	}

	for _, filename := range resolved {
		value, err := cuetils.ReadYAMLFileWithSpec(spec.CueSpec, filename)
		if err != nil {
			errs = multierror.Append(errs, err)
			continue
		}

		if p, ok, err := cuetils.Extract(value, "stages", cuetils.Map(StageFromStruct)); err != nil {
			errs = multierror.Append(errs, err)
			continue
		} else if ok {
			for name, stage := range p {
				stages[name] = stage
			}
		}
	}

	return stages, errs
}

func StageFromStruct(name string, v cue.Value) (*Stage, error) {
	out := &Stage{Name: name}

	if p, ok, err := cuetils.Extract(v, "name", cuetils.String); err != nil {
		return nil, fmt.Errorf("invalid config for %s: %w", "name", err)
	} else if ok {
		out.Name = p
	}

	if p, ok, err := cuetils.Extract(v, "type", cuetils.String); err != nil {
		return nil, fmt.Errorf("invalid config for %s: %w", "type", err)
	} else if ok {
		out.Type = p
	}

	if p, ok, err := cuetils.Extract(v, "address", cuetils.String); err != nil {
		return nil, fmt.Errorf("invalid config for %s: %w", "address", err)
	} else if ok {
		out.Type = p
	}

	if p, ok, err := cuetils.Extract(v, "ca_path", cuetils.String); err != nil {
		return nil, fmt.Errorf("invalid config for %s: %w", "ca_path", err)
	} else if ok {
		out.Type = p
	}

	if p, ok, err := cuetils.Extract(v, "cert_path", cuetils.String); err != nil {
		return nil, fmt.Errorf("invalid config for %s: %w", "cert_path", err)
	} else if ok {
		out.Type = p
	}

	if p, ok, err := cuetils.Extract(v, "key_path", cuetils.String); err != nil {
		return nil, fmt.Errorf("invalid config for %s: %w", "key_path", err)
	} else if ok {
		out.Type = p
	}

	if p, ok, err := cuetils.Extract(v, "ssh_config", SSHConfigFromStruct); err != nil {
		return nil, err
	} else if ok {
		out.SSHOptions = p
	}

	return out, nil
}

func SSHConfigFromStruct(_ string, v cue.Value) (*api.SSHOptions, error) {
	out := &api.SSHOptions{}

	if p, ok, err := cuetils.Extract(v, "host", cuetils.String); err != nil {
		return nil, fmt.Errorf("invalid config for %s: %w", "host", err)
	} else if ok {
		out.Hostname = p
	}

	if p, ok, err := cuetils.Extract(v, "port", cuetils.Int); err != nil {
		return nil, fmt.Errorf("invalid config for %s: %w", "port", err)
	} else if ok {
		out.Port = int32(p)
	}

	if p, ok, err := cuetils.Extract(v, "username", cuetils.String); err != nil {
		return nil, fmt.Errorf("invalid config for %s: %w", "username", err)
	} else if ok {
		out.Username = p
	}

	if p, ok, err := cuetils.Extract(v, "private_key_file", cuetils.String); err != nil {
		return nil, fmt.Errorf("invalid config for %s: %w", "private_key_file", err)
	} else if ok {
		out.PrivateKeyFile = p
	}

	return out, nil
}
