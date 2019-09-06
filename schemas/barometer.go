package schemas

import (
	"github.com/project-eria/xaal-go/device"
)

// Barometer : Simple barometer
func Barometer(addr string) (*device.Device, error) {
	dev, err := Basic(addr)
	dev.SetDevType("barometer.basic")

	// -- Attributes --
	// Temperature
	dev.NewAttribute("pressure", nil)
	return dev, err
}
