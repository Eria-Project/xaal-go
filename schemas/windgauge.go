package schemas

import (
	"xaal-go/device"
	"xaal-go/tools"
)

// Windgauge : Simple windgauge
func Windgauge(addr string) device.Device {
	if addr == "" {
		addr = tools.GetRandomUUID()
	}
	dev := device.Device{
		DevType: "windgauge.basic", Address: addr}

	// -- Attributes --
	// Strength of the wind
	dev.NewAttribute("windStrength", nil)
	// Direction of the wind
	dev.NewAttribute("windAngle", nil)
	// Strength of gusts
	dev.NewAttribute("gustStrength", nil)
	// Direction of gusts
	dev.NewAttribute("gustAngle", nil)

	return dev
}
