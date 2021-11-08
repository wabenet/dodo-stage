package config

import (
	"cuelang.org/go/cue"
	api "github.com/dodo-cli/dodo-stage/api/v1alpha1"
)

func VolumesFromValue(v cue.Value) ([]*api.PersistentVolume, error) {
	if out, err := VolumesFromMap(v); err == nil {
		return out, err
	}

	if out, err := VolumesFromList(v); err == nil {
		return out, err
	}

	return nil, ErrUnexpectedSpec
}

func VolumesFromMap(v cue.Value) ([]*api.PersistentVolume, error) {
	out := []*api.PersistentVolume{}

	err := eachInMap(v, func(name string, v cue.Value) error {
		r, err := VolumeFromValue(name, v)
		if err == nil {
			out = append(out, r)
		}

		return err

	})

	return out, err
}

func VolumesFromList(v cue.Value) ([]*api.PersistentVolume, error) {
	out := []*api.PersistentVolume{}

	err := eachInList(v, func(v cue.Value) error {
		r, err := VolumeFromValue("", v)
		if err == nil {
			out = append(out, r)
		}

		return err
	})

	return out, err
}

func VolumeFromValue(name string, v cue.Value) (*api.PersistentVolume, error) {
	if out, err := VolumeFromStruct(name, v); err == nil {
		return out, err
	}

	return nil, ErrUnexpectedSpec
}

func VolumeFromStruct(name string, v cue.Value) (*api.PersistentVolume, error) {
	out := &api.PersistentVolume{}

	if p, ok := property(v, "size"); ok {
		if v, err := BytesFromValue(p); err != nil {
			return nil, err
		} else {
			out.Size = v
		}
	}

	return out, nil
}
