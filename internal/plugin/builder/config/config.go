package config

import (
	"fmt"

	"cuelang.org/go/cue"
	"github.com/hashicorp/go-multierror"
	"github.com/wabenet/dodo-config/pkg/cuetils"
	"github.com/wabenet/dodo-config/pkg/includes"
	"github.com/wabenet/dodo-stage/pkg/spec"
)

type Stage struct {
	Name      string
	Type      string
	Provision *Provision
}

type Provision struct {
	Type string
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

	if p, ok, err := cuetils.Extract(v, "provision", ProvisionFromStruct); err != nil {
		return nil, err
	} else if ok {
		out.Provision = p
	}

	return out, nil
}

func ProvisionFromStruct(_ string, v cue.Value) (*Provision, error) {
	out := &Provision{}

	if p, ok, err := cuetils.Extract(v, "type", cuetils.String); err != nil {
		return nil, fmt.Errorf("invalid config for %s: %w", "type", err)
	} else if ok {
		out.Type = p
	}

	return out, nil
}
