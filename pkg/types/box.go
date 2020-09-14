package types

import (
	"reflect"

	"github.com/dodo-cli/dodo-core/pkg/decoder"
)

func DecodeBox(target interface{}) decoder.Decoding {
	box := *(target.(**Box))

	return decoder.Kinds(map[reflect.Kind]decoder.Decoding{
		reflect.Map: decoder.Keys(map[string]decoder.Decoding{
			"user":         decoder.String(&box.User),
			"name":         decoder.String(&box.User),
			"version":      decoder.String(&box.Version),
			"access_token": decoder.String(&box.AccessToken),
		}),
	})
}
