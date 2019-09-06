package schemas

import (
	"github.com/project-eria/xaal-go/device"
)

// Gateway : Simple gateway that manage physical devices
func Gateway(addr string) (*device.Device, error) {
	dev, err := Basic(addr)
	dev.SetDevType("gateway.basic")

	// -- Attributes --
	// Embeded devices
	dev.NewAttribute("embedded", nil)

	return dev, err
}
