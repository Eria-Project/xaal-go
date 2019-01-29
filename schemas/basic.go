package schemas

import (
	"github.com/Eria-Project/xaal-go/device"
)

// Basic : Generic schema for any devices
func Basic(addr string) *device.Device {
	dev, _ := device.New("basic.basic", addr)
	return dev
}
