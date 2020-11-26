package types

import (
	"reflect"

	"github.com/dodo-cli/dodo-core/pkg/decoder"
	api "github.com/dodo-cli/dodo-stage/api/v1alpha1"
)

func NewVolume() decoder.Producer {
	return func() (interface{}, decoder.Decoding) {
		target := &api.PersistentVolume{}
		return &target, DecodeVolume(&target)
	}
}

func DecodeVolume(target interface{}) decoder.Decoding {
	vol := *(target.(**api.PersistentVolume))

	return decoder.Kinds(map[reflect.Kind]decoder.Decoding{
		reflect.Map: decoder.Keys(map[string]decoder.Decoding{
			"size": Bytes(&vol.Size),
		}),
	})
}
