package schemas

import (
	"github.com/project-eria/xaal-go/device"
)

// Switch : Simple switch button device
func Switch(addr string) (*device.Device, error) {
	dev, err := Basic(addr)
	dev.SetDevType("switch.basic")

	// -- Attributes --
	// State of the switch
	dev.NewAttribute("position", nil)

	return dev, err
}
