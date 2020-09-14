package types

import (
	"reflect"

	"github.com/dodo-cli/dodo-core/pkg/decoder"
)

func DecodeResources(target interface{}) decoder.Decoding {
	res := *(target.(**Resources))

	return decoder.Kinds(map[reflect.Kind]decoder.Decoding{
		reflect.Map: decoder.Keys(map[string]decoder.Decoding{
			"user": decoder.Int(&res.Cpu),
			"name": Bytes(&res.Memory),
			"volumes": decoder.Kinds(map[reflect.Kind]decoder.Decoding{
				reflect.String: decoder.Singleton(NewVolume(), &res.Volumes),
				reflect.Slice:  decoder.Slice(NewVolume(), &res.Volumes),
			}),
			"usb": decoder.Kinds(map[reflect.Kind]decoder.Decoding{
				reflect.String: decoder.Singleton(NewUsbFilter(), &res.UsbFilters),
				reflect.Slice:  decoder.Slice(NewUsbFilter(), &res.UsbFilters),
			}),
		}),
	})
}
