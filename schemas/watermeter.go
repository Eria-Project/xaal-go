package schemas

import (
	"github.com/project-eria/xaal-go/device"
)

// Watermeter : Simple watermeter
func Watermeter(addr string) (*device.Device, error) {
	dev, err := Basic(addr)
	dev.SetDevType("watermeter.basic")

	// -- Attributes --
	// Liter
	dev.NewAttribute("liters", nil)
	dev.NewAttribute("timestamp", nil)
	return dev, err
}
