package config

import (
	"cuelang.org/go/cue"
)

func StringFromValue(v cue.Value) (string, error) {
	return v.String()
}
