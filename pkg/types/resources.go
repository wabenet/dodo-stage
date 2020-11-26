package types

import (
	"reflect"

	"github.com/dodo-cli/dodo-core/pkg/decoder"
	api "github.com/dodo-cli/dodo-stage/api/v1alpha1"
)

func DecodeResources(target interface{}) decoder.Decoding {
	res := *(target.(**api.Resources))

	return decoder.Kinds(map[reflect.Kind]decoder.Decoding{
		reflect.Map: decoder.Keys(map[string]decoder.Decoding{
			"cpu": decoder.Int(&res.Cpu),
			"memory": Bytes(&res.Memory),
			"volumes": decoder.Kinds(map[reflect.Kind]decoder.Decoding{
				reflect.Map: decoder.Singleton(NewVolume(), &res.Volumes),
				reflect.Slice:  decoder.Slice(NewVolume(), &res.Volumes),
			}),
			"usb": decoder.Kinds(map[reflect.Kind]decoder.Decoding{
				reflect.Map: decoder.Singleton(NewUsbFilter(), &res.UsbFilters),
				reflect.Slice:  decoder.Slice(NewUsbFilter(), &res.UsbFilters),
			}),
		}),
	})
}
