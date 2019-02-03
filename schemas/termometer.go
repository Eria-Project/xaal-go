package schemas

import (
	"github.com/project-eria/xaal-go/device"
)

// Thermometer : Simple thermometer
func Thermometer(addr string) *device.Device {
	dev, _ := device.New("thermometer.basic", addr)

	// -- Attributes --
	// Temperature
	dev.NewAttribute("temperature", nil)
	return dev
}
