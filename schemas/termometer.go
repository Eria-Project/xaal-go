package schemas

import (
	"xaal-go/device"
	"xaal-go/tools"
)

// Thermometer : Simple thermometer
func Thermometer(addr string) device.Device {
	if addr == "" {
		addr = tools.GetRandomUUID()
	}
	dev := device.Device{
		DevType: "thermometer.basic", Address: addr}

	// -- Attributes --
	// Temperature
	dev.NewAttribute("temperature", nil)
	return dev
}
