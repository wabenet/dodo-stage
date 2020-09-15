package types

import (
	"reflect"

	"github.com/dodo-cli/dodo-core/pkg/decoder"
)

func NewStage() decoder.Producer {
	return func() (interface{}, decoder.Decoding) {
		target := &Stage{
			Box: &Box{},
			Resources: &Resources{
				Volumes:    []*PersistentVolume{},
				UsbFilters: []*UsbFilter{},
			},
		}
		return &target, DecodeStage(&target)
	}
}

func DecodeStage(target interface{}) decoder.Decoding {
	stage := *(target.(**Stage))

	return decoder.Kinds(map[reflect.Kind]decoder.Decoding{
		reflect.Map: decoder.Keys(map[string]decoder.Decoding{
			"type":      decoder.String(&stage.Type),
			"box":       DecodeBox(&stage.Box),
			"resources": DecodeResources(&stage.Resources),
		}),
	})
}
