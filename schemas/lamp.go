package schemas

import (
	"github.com/project-eria/xaal-go/device"

	logger "github.com/project-eria/eria-logger"
)

// Lamp : Simple switch lamp
func Lamp(addr string) (*device.Device, error) {
	dev, err := Basic(addr)
	dev.SetDevType("lamp.basic")

	// -- Attributes --
	// State of the lamp (bool)
	dev.NewAttribute("light", nil)

	// -- Methods --
	// On: Switch on the lamp
	dev.AddMethod("on", defaultOn, nil)

	// Off: Switch off the lamp
	dev.AddMethod("off", defaultOff, nil)

	return dev, err
}

// LampDimmer : Lamp with a dimmer
func LampDimmer(addr string) (*device.Device, error) {
	// Extend "lamp.basic"
	dev, err := Lamp(addr)
	dev.SetDevType("lamp.dimmer")

	// -- Attributes --
	// Level of the dimmer (int percentage unit:%, minimum:0, maximum:100)
	dev.NewAttribute("dimmer", nil)

	// -- Methods --
	// Dim: Change the dimmer of the lamp
	// params:
	// - target: "Target of the dimmer" (int percentage unit:%, minimum:0, maximum:100)
	dev.AddMethod("dim", defaultDim, &[]string{"target"})

	return dev, err
}

func defaultOn(d *device.Device, args map[string]interface{}) map[string]interface{} {
	// """Switch on the lamp"""
	logger.Module("xaal:schema-lamp").Debug("defaultOn()")
	return nil
}

func defaultOff(d *device.Device, args map[string]interface{}) map[string]interface{} {
	// """Switch off the lamp"""
	logger.Module("xaal:schema-lamp").Debug("defaultOff()")
	return nil
}

func defaultDim(d *device.Device, args map[string]interface{}) map[string]interface{} {
	// """Dim the lamp"""
	logger.Module("xaal:schema-lamp").WithField("target", args["target"]).Debug("defaultDim()")
	return nil
}
