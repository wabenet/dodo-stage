package config

import (
	"errors"
	"strings"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
	"cuelang.org/go/encoding/yaml"
	api "github.com/dodo-cli/dodo-stage/api/v1alpha1"
	"github.com/dodo-cli/dodo-config/pkg/template"
	"github.com/dodo-cli/dodo-stage/pkg/spec"
	"github.com/hashicorp/go-multierror"
)

type Config struct {
	Stages   map[string]*api.Stage
	Includes []string
}

func GetAllStages(filenames ...string) (map[string]*api.Stage, error) {
	var errs error
	stages := map[string]*api.Stage{}

	for _, filename := range filenames {
		config, err := ParseConfig(filename)
		if err != nil {
			errs = multierror.Append(errs, err)
			continue
		}

		for name, stage := range config.Stages {
			stages[name] = stage
		}

		for _, include := range config.Includes {
			included, err := GetAllStages(include)
			if err != nil {
				errs = multierror.Append(errs, err)
				continue
			}

			for name, stage := range included {
				stages[name] = stage
			}
		}
	}

	return stages, errs
}

func ParseConfig(filename string) (*Config, error) {
	ctx := cuecontext.New()

	bis := load.Instances([]string{"-"}, &load.Config{
		Stdin: strings.NewReader(spec.CueSpec),
	})

	if len(bis) != 1 {
		return nil, errors.New("expected exactly one instance")
	}

	bi := bis[0]

	if bi.Err != nil {
		return nil, bi.Err
	}

	yamlFile, err := yaml.Extract(filename, nil)
	if err != nil {
		return nil, err
	}

	yamlFile, err = template.TemplateCueAST(yamlFile)
	if err != nil {
		return nil, err
	}

	if err := bi.AddSyntax(yamlFile); err != nil {
		return nil, err
	}

	value := ctx.BuildInstance(bi)
	if err := value.Err(); err != nil {
		return nil, err
	}

	if err := value.Validate(cue.Concrete(true), cue.Final()); err != nil {
		return nil, err
	}

	return ConfigFromValue(value)
}

func ConfigFromValue(v cue.Value) (*Config, error) {
	out := &Config{}

	if p, ok := property(v, "stages"); ok {
		if s, err := StagesFromValue(p); err != nil {
			return nil, err
		} else {
			out.Stages = s
		}
	}

	if p, ok := property(v, "include"); ok {
		if is, err := IncludesFromValue(p); err != nil {
			return nil, err
		} else {
			out.Includes = is
		}
	}

	return out, nil
}
