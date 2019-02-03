package schemas

import (
	"github.com/project-eria/xaal-go/device"
)

// Barometer : Simple barometer
func Barometer(addr string) *device.Device {
	dev, _ := device.New("barometer.basic", addr)

	// -- Attributes --
	// Temperature
	dev.NewAttribute("pressure", nil)
	return dev
}
