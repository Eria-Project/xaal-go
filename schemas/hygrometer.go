package schemas

import (
	"xaal-go/device"
	"xaal-go/tools"
)

// Hygrometer : Simple hygrometer
func Hygrometer(addr string) device.Device {
	if addr == "" {
		addr = tools.GetRandomUUID()
	}
	dev := device.Device{
		DevType: "hygrometer.basic", Address: addr}

	// -- Attributes --
	// Temperature
	dev.NewAttribute("humidity", nil)
	return dev
}
