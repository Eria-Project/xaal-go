package schemas

import (
	"github.com/project-eria/xaal-go/device"
)

// Thermometer : Simple thermometer
func Thermometer(addr string) (*device.Device, error) {
	dev, err := Basic(addr)
	dev.SetDevType("thermometer.basic")

	// -- Attributes --
	// Temperature (float)
	dev.NewAttribute("temperature", nil)
	return dev, err
}
