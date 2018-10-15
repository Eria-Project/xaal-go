package schemas

import (
	"xaal-go/device"
)

// Hygrometer : Simple hygrometer
func Hygrometer(addr string) *device.Device {
	dev, _ := device.New("hygrometer.basic", addr)

	// -- Attributes --
	// Temperature
	dev.NewAttribute("humidity", nil)
	return dev
}
