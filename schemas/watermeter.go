package schemas

import (
	"github.com/project-eria/xaal-go/device"
)

// Watermeter : Simple watermeter
func Watermeter(addr string) *device.Device {
	dev, _ := device.New("watermeter.basic", addr)

	// -- Attributes --
	// Liter
	dev.NewAttribute("liters", nil)
	dev.NewAttribute("timestamp", nil)
	return dev
}
