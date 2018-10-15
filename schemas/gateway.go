package schemas

import (
	"xaal-go/device"
)

// Gateway : Simple gateway that manage physical devices
func Gateway(addr string) *device.Device {
	dev, _ := device.New("gateway.basic", addr)

	// -- Attributes --
	// Embeded devices
	dev.NewAttribute("embedded", nil)

	return dev
}
