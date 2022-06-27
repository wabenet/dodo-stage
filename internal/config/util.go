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
