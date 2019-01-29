package schemas

import (
	"github.com/Eria-Project/xaal-go/device"
)

// Windgauge : Simple windgauge
func Windgauge(addr string) *device.Device {
	dev, _ := device.New("windgauge.basic", addr)

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
