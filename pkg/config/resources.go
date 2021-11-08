package config

import (
	"cuelang.org/go/cue"
	api "github.com/dodo-cli/dodo-stage/api/v1alpha1"
)

func ResourcesFromValue(v cue.Value) (*api.Resources, error) {
	if out, err := ResourcesFromStruct(v); err == nil {
		return out, err
	}

	return nil, ErrUnexpectedSpec
}

func ResourcesFromStruct(v cue.Value) (*api.Resources, error) {
	out := &api.Resources{}

	if p, ok := property(v, "cpu"); ok {
		if n, err := IntFromValue(p); err != nil {
			return nil, err
		} else {
			out.Cpu = n
		}
	}

	if p, ok := property(v, "memory"); ok {
		if n, err := BytesFromValue(p); err != nil {
			return nil, err
		} else {
			out.Memory = n
		}
	}

	if p, ok := property(v, "volumes"); ok {
		if n, err := VolumesFromValue(p); err != nil {
			return nil, err
		} else {
			out.Volumes = n
		}
	}

	if p, ok := property(v, "usb_filters"); ok {
		if n, err := USBFiltersFromValue(p); err != nil {
			return nil, err
		} else {
			out.UsbFilters = n
		}
	}

	return out, nil
}
