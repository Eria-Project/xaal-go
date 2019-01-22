package schemas

import (
	"xaal-go/device"
)

// Lamp : Simple switch lamp
func Lamp(addr string) *device.Device {
	dev, _ := device.New("lamp.basic", addr)

	// -- Attributes --
	// State of the lamp
	dev.NewAttribute("light", nil)

	/* TODO
	   # -- Methods --
	   def default_on():
	       """Switch on the lamp"""
	       logger.info("default_on()")

	   def default_off():
	       """Switch off the lamp"""
	       logger.info("default_off()")

	   dev.add_method('on',default_on)
	   dev.add_method('off',default_off)
	*/
	return dev
}
