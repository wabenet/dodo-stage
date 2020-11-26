package types

import (
	"reflect"

	"github.com/dodo-cli/dodo-core/pkg/decoder"
	api "github.com/dodo-cli/dodo-stage/api/v1alpha1"
)

func NewStage() decoder.Producer {
	return func() (interface{}, decoder.Decoding) {
		target := &api.Stage{
			Box: &api.Box{},
			Resources: &api.Resources{
				Volumes:    []*api.PersistentVolume{},
				UsbFilters: []*api.UsbFilter{},
			},
		}
		return &target, DecodeStage(&target)
	}
}

func DecodeStage(target interface{}) decoder.Decoding {
	stage := *(target.(**api.Stage))

	return decoder.Kinds(map[reflect.Kind]decoder.Decoding{
		reflect.Map: decoder.Keys(map[string]decoder.Decoding{
			"type":      decoder.String(&stage.Type),
			"box":       DecodeBox(&stage.Box),
			"resources": DecodeResources(&stage.Resources),
		}),
	})
}
