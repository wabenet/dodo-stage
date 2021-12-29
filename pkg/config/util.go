package config

import (
	"cuelang.org/go/cue"
	"github.com/alecthomas/units"
)

func StringFromValue(v cue.Value) (string, error) {
	return v.String()
}

func IntFromValue(v cue.Value) (int64, error) {
	return v.Int64()
}

func BytesFromValue(v cue.Value) (int64, error) {
	num, err := v.String()
	if err != nil {
		return 0, err
	}

	return units.ParseStrictBytes(num)
}

func property(v cue.Value, name string) (cue.Value, bool) {
	p := v.LookupPath(cue.MakePath(cue.Str(name)))
	return p, p.Exists()
}

func eachInList(v cue.Value, f func(cue.Value) error) error {
	iter, err := v.List()
	if err != nil {
		return err
	}

	for iter.Next() {
		if err := f(iter.Value()); err != nil {
			return err
		}
	}

	return nil
}

func eachInMap(v cue.Value, f func(string, cue.Value) error) error {
	iter, err := v.Fields()
	if err != nil {
		return err
	}

	for iter.Next() {
		name := iter.Selector().String()

		if err := f(name, iter.Value()); err != nil {
			return err
		}
	}

	return nil
}
