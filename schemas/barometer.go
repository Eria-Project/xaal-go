package schemas

import (
	"xaal-go/device"
	"xaal-go/tools"
)

// Barometer : Simple barometer
func Barometer(addr string) device.Device {
	if addr == "" {
		addr = tools.GetRandomUUID()
	}
	dev := device.Device{
		DevType: "barometer.basic", Address: addr}

	// -- Attributes --
	// Temperature
	dev.NewAttribute("pressure", nil)
	return dev
}
