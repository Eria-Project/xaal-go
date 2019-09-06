package schemas

import (
	"github.com/project-eria/xaal-go/device"
)

// Hygrometer : Simple hygrometer
func Hygrometer(addr string) (*device.Device, error) {
	dev, err := Basic(addr)
	dev.SetDevType("hygrometer.basic")

	// -- Attributes --
	// Temperature
	dev.NewAttribute("humidity", nil)
	return dev, err
}
