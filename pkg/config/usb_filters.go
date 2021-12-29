package config

import (
	"cuelang.org/go/cue"
	api "github.com/dodo-cli/dodo-stage/api/v1alpha1"
	"github.com/hashicorp/go-multierror"
)

func USBFiltersFromValue(v cue.Value) ([]*api.UsbFilter, error) {
	var errs error

	if out, err := USBFiltersFromMap(v); err == nil {
		return out, err
	} else {
		errs = multierror.Append(errs, err)
	}

	if out, err := USBFiltersFromList(v); err == nil {
		return out, err
	} else {
		errs = multierror.Append(errs, err)
	}

	return nil, errs
}

func USBFiltersFromMap(v cue.Value) ([]*api.UsbFilter, error) {
	out := []*api.UsbFilter{}

	err := eachInMap(v, func(name string, v cue.Value) error {
		r, err := USBFilterFromValue(name, v)
		if err == nil {
			out = append(out, r)
		}

		return err

	})

	return out, err
}

func USBFiltersFromList(v cue.Value) ([]*api.UsbFilter, error) {
	out := []*api.UsbFilter{}

	err := eachInList(v, func(v cue.Value) error {
		r, err := USBFilterFromValue("", v)
		if err == nil {
			out = append(out, r)
		}

		return err
	})

	return out, err
}

func USBFilterFromValue(name string, v cue.Value) (*api.UsbFilter, error) {
	var errs error

	if out, err := USBFilterFromStruct(name, v); err == nil {
		return out, err
	} else {
		errs = multierror.Append(errs, err)
	}

	return nil, errs
}

func USBFilterFromStruct(name string, v cue.Value) (*api.UsbFilter, error) {
	out := &api.UsbFilter{}

	if p, ok := property(v, "name"); ok {
		if v, err := StringFromValue(p); err != nil {
			return nil, err
		} else {
			out.Name = v
		}
	}

	if p, ok := property(v, "vendorid"); ok {
		if v, err := StringFromValue(p); err != nil {
			return nil, err
		} else {
			out.VendorId = v
		}
	}

	if p, ok := property(v, "productid"); ok {
		if v, err := StringFromValue(p); err != nil {
			return nil, err
		} else {
			out.ProductId = v
		}
	}

	return out, nil
}
