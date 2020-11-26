package types

import (
	"reflect"

	"github.com/dodo-cli/dodo-core/pkg/decoder"
	api "github.com/dodo-cli/dodo-stage/api/v1alpha1"
)

func NewUsbFilter() decoder.Producer {
	return func() (interface{}, decoder.Decoding) {
		target := &api.UsbFilter{}
		return &target, DecodeUsbFilter(&target)
	}
}

func DecodeUsbFilter(target interface{}) decoder.Decoding {
	usb := *(target.(**api.UsbFilter))

	return decoder.Kinds(map[reflect.Kind]decoder.Decoding{
		reflect.Map: decoder.Keys(map[string]decoder.Decoding{
			"name":       decoder.String(&usb.Name),
			"vendorid":  decoder.String(&usb.VendorId),
			"productid": decoder.String(&usb.ProductId),
		}),
	})
}
