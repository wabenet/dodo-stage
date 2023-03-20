package config

import (
	"cuelang.org/go/cue"
	"github.com/hashicorp/go-multierror"
	"github.com/wabenet/dodo-config/pkg/cuetils"
	api "github.com/wabenet/dodo-stage/api/stage/v1alpha3"
)

func ResourcesFromValue(v cue.Value) (*api.Resources, error) {
	var errs error

	if out, err := ResourcesFromStruct(v); err == nil {
		return out, err
	} else {
		errs = multierror.Append(errs, err)
	}

	return nil, errs
}

func ResourcesFromStruct(v cue.Value) (*api.Resources, error) {
	out := &api.Resources{}

	if p, ok := cuetils.Get(v, "cpu"); ok {
		if n, err := IntFromValue(p); err != nil {
			return nil, err
		} else {
			out.Cpu = n
		}
	}

	if p, ok := cuetils.Get(v, "memory"); ok {
		if n, err := BytesFromValue(p); err != nil {
			return nil, err
		} else {
			out.Memory = n
		}
	}

	if p, ok := cuetils.Get(v, "volumes"); ok {
		if n, err := VolumesFromValue(p); err != nil {
			return nil, err
		} else {
			out.Volumes = n
		}
	}

	if p, ok := cuetils.Get(v, "usb_filters"); ok {
		if n, err := USBFiltersFromValue(p); err != nil {
			return nil, err
		} else {
			out.UsbFilters = n
		}
	}

	return out, nil
}
