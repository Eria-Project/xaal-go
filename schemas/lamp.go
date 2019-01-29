package schemas

import (
	"github.com/Eria-Project/xaal-go/device"

	"github.com/Eria-Project/logger"
)

// Lamp : Simple switch lamp
func Lamp(addr string) *device.Device {
	dev, _ := device.New("lamp.basic", addr)

	// -- Attributes --
	// State of the lamp
	dev.NewAttribute("light", nil)

	// -- Methods --
	dev.AddMethod("on", defaultOn)
	dev.AddMethod("off", defaultOff)

	return dev
}

func defaultOn(d *device.Device, args map[string]interface{}) map[string]interface{} {
	// """Switch on the lamp"""
	logger.Debug("defaultOn()")
	return nil
}

func defaultOff(d *device.Device, args map[string]interface{}) map[string]interface{} {
	// """Switch off the lamp"""
	logger.Debug("defaultOff()")
	return nil
}
