package schemas

import (
	"github.com/project-eria/xaal-go/device"
)

// Switch : Simple switch button device
func Switch(addr string) *device.Device {
	dev, _ := device.New("switch.basic", addr)

	// -- Attributes --
	// State of the switch
	dev.NewAttribute("position", nil)

	return dev
}
