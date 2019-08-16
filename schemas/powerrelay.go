package schemas

import (
	"github.com/project-eria/xaal-go/device"

	"github.com/project-eria/logger"
)

// Powerrelay returns a simple power relay device
func Powerrelay(addr string) *device.Device {
	dev, _ := device.New("powerrelay.basic", addr)

	// -- Attributes --
	// State of the relay
	dev.NewAttribute("power", nil)

	// -- Methods --
	dev.AddMethod("on", defaultRelayOn)
	dev.AddMethod("off", defaultRelayOff)

	return dev
}

func defaultRelayOn(d *device.Device, args map[string]interface{}) map[string]interface{} {
	// """Switch on the relay"""
	logger.Module("schema-powerrelay").Debug("defaultRelayOn()")
	return nil
}

func defaultRelayOff(d *device.Device, args map[string]interface{}) map[string]interface{} {
	// """Switch off the relay"""
	logger.Module("schema-powerrelay").Debug("defaultRelayOff()")
	return nil
}
