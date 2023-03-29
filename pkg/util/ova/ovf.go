package ova

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
)

const (
	TypeCPU     = "3"
	TypeMemory  = "4"
	TypeIDE     = "5"
	TypeSCSI    = "6"
	TypeNetwork = "10"
	TypeFloppy  = "14"
	TypeCDROM   = "15"
	TypeDisk    = "17"
	TypeUSB     = "23"
)

type Envelope struct {
	File           []References   `xml:"References>File"`
	DiskSection    DiskSection    `xml:"DiskSection"`
	NetworkSection NetworkSection `xml:"NetworkSection"`
	VirtualSystem  VirtualSystem  `xml:"VirtualSystem"`
}

type References struct {
	HRef string `xml:"href,attr"`
	ID   string `xml:"id,attr"`
	Size string `xml:"size,attr"`
}

type DiskSection struct {
	Info string        `xml:"Info"`
	Disk []DiskDetails `xml:"Disk"`
}

type DiskDetails struct {
	Capacity                string `xml:"capacity,attr"`
	CapacityAllocationUnits string `xml:"capacityAllocationUnits,attr"`
	DiskID                  string `xml:"diskId,attr"`
	FileRef                 string `xml:"fileRef,attr"`
	Format                  string `xml:"format,attr"`
	PopulatedSize           string `xml:"populatedSize,attr"`
}

type NetworkSection struct {
	Info    string  `xml:"Info"`
	Network Network `xml:"Network"`
}

type Network struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"Description"`
}

type VirtualSystem struct {
	ID              string          `xml:"id,attr"`
	Info            string          `xml:"Info"`
	Name            string          `xml:"Name"`
	OperatingSystem OperatingSystem `xml:"OperatingSystemSection"`
	VirtualHardware VirtualHardware `xml:"VirtualHardwareSection"`
}

type OperatingSystem struct {
	ID          string `xml:"id,attr"`
	OSType      string `xml:"osType,attr"`
	Info        string `xml:"Info"`
	Description string `xml:"Description"`
}

type VirtualHardware struct {
	Info   string                `xml:"Info"`
	System VirtualHardwareSystem `xml:"System"`
	Items  []VirtualHardwareItem `xml:"Item"`
}

type VirtualHardwareSystem struct {
	ElementName             string `xml:"ElementName"`
	InstanceID              string `xml:"InstanceID"`
	VirtualSystemIdentifier string `xml:"VirtualSystemIdentifier"`
	VirtualSystemType       string `xml:"VirtualSystemType"`
}

type VirtualHardwareItem struct {
	Required            string `xml:"required,attr"`
	Address             string `xml:"Address"`
	AddressOnParent     string `xml:"AddressOnParent"`
	AllocationUnits     string `xml:"AllocationUnits"`
	AutomaticAllocation string `xml:"AutomaticAllocation"`
	Connection          string `xml:"Connection"`
	Description         string `xml:"Description"`
	ElementName         string `xml:"ElementName"`
	HostResource        string `xml:"HostResource"`
	InstanceID          string `xml:"InstanceID"`
	Parent              string `xml:"Parent"`
	ResourceSubType     string `xml:"ResourceSubType"`
	ResourceType        string `xml:"ResourceType"`
	VirtualQuantity     string `xml:"VirtualQuantity"`
}

func ReadOVF(path string) (*Envelope, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("could not determine file path : %w", err)
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read OVF file: %w", err)
	}

	var envelope Envelope
	if err := xml.Unmarshal(bytes, &envelope); err != nil {
		return nil, fmt.Errorf("could not parse OVF file: %w", err)
	}

	return &envelope, nil
}
