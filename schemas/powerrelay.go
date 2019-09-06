package schemas

import (
	logger "github.com/project-eria/eria-logger"
	"github.com/project-eria/xaal-go/device"
)

// Powerrelay returns a simple power relay device
func Powerrelay(addr string) (*device.Device, error) {
	dev, err := Basic(addr)
	dev.SetDevType("powerrelay.basic")

	// -- Attributes --
	// State of the relay
	dev.NewAttribute("power", nil)

	// -- Methods --
	dev.AddMethod("on", defaultRelayOn, nil)
	dev.AddMethod("off", defaultRelayOff, nil)

	return dev, err
}

func defaultRelayOn(d *device.Device, args map[string]interface{}) map[string]interface{} {
	// """Switch on the relay"""
	logger.Module("xaal:schema-powerrelay").Debug("defaultRelayOn()")
	return nil
}

func defaultRelayOff(d *device.Device, args map[string]interface{}) map[string]interface{} {
	// """Switch off the relay"""
	logger.Module("xaal:schema-powerrelay").Debug("defaultRelayOff()")
	return nil
}
