package schemas

import (
	logger "github.com/project-eria/eria-logger"
	"github.com/project-eria/xaal-go/device"
)

// Shutter returns a simple shutter
func Shutter(addr string) (*device.Device, error) {
	dev, err := Basic(addr)
	dev.SetDevType("shutter.basic")

	// -- Attributes --
	// Ongoing action of the shutter (string: up/down/stop)
	dev.NewAttribute("action", nil)

	// -- Methods --
	// Up: Up the shutter
	dev.AddMethod("up", defaultUp, nil)
	// Down: Down the shutter
	dev.AddMethod("down", defaultDown, nil)
	// Stop: Stop ongoing action of the shutter
	dev.AddMethod("stop", defaultStop, nil)

	return dev, err
}

// ShutterPosition returns a shutter with a position managment
func ShutterPosition(addr string) (*device.Device, error) {
	// Extend "shutter.basic"
	dev, err := Shutter(addr)
	dev.SetDevType("shutter.position")

	// -- Attributes --
	// Level of aperture of the shutter (int percentage unit:%, minimum:0, maximum:100)
	dev.NewAttribute("position", nil)

	// -- Methods --
	// Position: Change the position of the shutter
	// params:
	// - target: "Target of the position" (int percentage unit:%, minimum:0, maximum:100)
	dev.AddMethod("position", defaultPosition, &[]string{"target"})

	return dev, err
}

func defaultUp(d *device.Device, args map[string]interface{}) map[string]interface{} {
	// """Up the shutter"""
	logger.Module("xaal:schema-shutter").Debug("defaultUp()")
	return nil
}

func defaultDown(d *device.Device, args map[string]interface{}) map[string]interface{} {
	// """Down the shutter"""
	logger.Module("xaal:schema-shutter").Debug("defaultDown()")
	return nil
}

func defaultStop(d *device.Device, args map[string]interface{}) map[string]interface{} {
	// """Stop ongoing action of the shutter"""
	logger.Module("xaal:schema-shutter").Debug("defaultStop()")
	return nil
}

func defaultPosition(d *device.Device, args map[string]interface{}) map[string]interface{} {
	// """Change the position of the shutter"""
	logger.Module("xaal:schema-shutter").WithField("target", args["target"]).Debug("defaultPosition()")
	return nil
}
