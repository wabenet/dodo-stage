package types

import (
	"fmt"
	"reflect"

	"github.com/alecthomas/units"
	"github.com/dodo-cli/dodo-core/pkg/decoder"
)

func Bytes(target interface{}) decoder.Decoding {
	return decoder.Kinds(map[reflect.Kind]decoder.Decoding{
		reflect.Int:   decoder.Int(target),
		reflect.Int8:  decoder.Int(target),
		reflect.Int16: decoder.Int(target),
		reflect.Int32: decoder.Int(target),
		reflect.Int64: decoder.Int(target),
		reflect.String: func(d *decoder.Decoder, config interface{}) {
			decoded, ok := config.(string)
			if !ok {
				d.Error(fmt.Errorf("could not decode string: %w", decoder.ErrUnexpectedType))
				return
			}

			templated, err := decoder.ApplyTemplate(d, decoded)
			if err != nil {
				d.Error(err)
				return
			}

                        result, err := units.ParseStrictBytes(templated)
			if err != nil {
				d.Error(err)
				return
			}

			reflect.ValueOf(target).Elem().SetInt(result)
		},
	})
}
