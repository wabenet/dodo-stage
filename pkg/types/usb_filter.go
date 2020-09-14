package types

import (
	"reflect"

	"github.com/dodo-cli/dodo-core/pkg/decoder"
)

func NewUsbFilter() decoder.Producer {
	return func() (interface{}, decoder.Decoding) {
		target := &UsbFilter{}
		return &target, DecodeUsbFilter(&target)
	}
}

func DecodeUsbFilter(target interface{}) decoder.Decoding {
	usb := *(target.(**UsbFilter))

	return decoder.Kinds(map[reflect.Kind]decoder.Decoding{
		reflect.Map: decoder.Keys(map[string]decoder.Decoding{
			"name":       decoder.String(&usb.Name),
			"vendor_id":  decoder.String(&usb.VendorId),
			"product_id": decoder.String(&usb.ProductId),
		}),
	})
}
