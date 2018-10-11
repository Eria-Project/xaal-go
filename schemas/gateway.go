package schemas

import (
	"xaal-go/device"
	"xaal-go/tools"
)

// Gateway : Simple gateway that manage physical devices
func Gateway(addr string) device.Device {
	if addr == "" {
		addr = tools.GetRandomUUID()
	}

	dev := device.Device{
		DevType: "gateway.basic", Address: addr}

	// -- Attributes --
	// Embeded devices
	dev.NewAttribute("embedded", nil)

	return dev
}
