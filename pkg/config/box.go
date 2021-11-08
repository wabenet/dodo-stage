package config

import (
	"cuelang.org/go/cue"
	api "github.com/dodo-cli/dodo-stage/api/v1alpha1"
)

func BoxFromValue(v cue.Value) (*api.Box, error) {
	if out, err := BoxFromStruct(v); err == nil {
		return out, err
	}

	return nil, ErrUnexpectedSpec
}

func BoxFromStruct(v cue.Value) (*api.Box, error) {
	out := &api.Box{}

	if p, ok := property(v, "user"); ok {
		if n, err := StringFromValue(p); err != nil {
			return nil, err
		} else {
			out.User = n
		}
	}

	if p, ok := property(v, "name"); ok {
		if n, err := StringFromValue(p); err != nil {
			return nil, err
		} else {
			out.Name = n
		}
	}

	if p, ok := property(v, "version"); ok {
		if n, err := StringFromValue(p); err != nil {
			return nil, err
		} else {
			out.Version = n
		}
	}

	if p, ok := property(v, "access_token"); ok {
		if n, err := StringFromValue(p); err != nil {
			return nil, err
		} else {
			out.AccessToken = n
		}
	}

	return out, nil
}
