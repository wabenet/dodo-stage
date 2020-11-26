package types

import (
	"reflect"

	"github.com/dodo-cli/dodo-core/pkg/decoder"
	api "github.com/dodo-cli/dodo-stage/api/v1alpha1"
)

func DecodeBox(target interface{}) decoder.Decoding {
	box := *(target.(**api.Box))

	return decoder.Kinds(map[reflect.Kind]decoder.Decoding{
		reflect.Map: decoder.Keys(map[string]decoder.Decoding{
			"user":         decoder.String(&box.User),
			"name":         decoder.String(&box.Name),
			"version":      decoder.String(&box.Version),
			"access_token": decoder.String(&box.AccessToken),
		}),
	})
}
