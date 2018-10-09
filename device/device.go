package device

import (
	"errors"
	"fmt"
	"xaal-go/config"
	"xaal-go/tools"
)

// Device : xAAL internal attributes for a device
type Device struct {
	DevType     string // xaal devtype
	Address     string // xaal addr
	alivePeriod int    // time in sec between two alive

	/*
		self.devtype = devtype          # xaal devtype
		self.address = addr             # xaal addr
		self.vendor_id = None           # vendor ID ie : ACME
		self.product_id = None          # product ID
		self.version = None             # product release
		self.url = None                 # product URL
		self.info = None                # additionnal info
		self.hw_id = None               # hardware info
		self.group_id = None            # group devices
		# Some useless attributes, only for compatibility
		self.bus_addr = config.address
		self.bus_port = config.port
		self.hops = config.hops
		# Unsupported stuffs
		self.unsupported_attributes = []
		self.unsupported_methods = []
		self.unsupported_notifications = []
		# Default attributes & methods
		self.__attributes = Attributes()
		self.methods = {'getAttributes' : self._get_attributes,
						'getDescription': self._get_description }
		self.engine = engine
	*/
}

/*SetDevType : Set the device type */
func (d *Device) SetDevType(devType string) error {
	if !tools.IsValidDevType(devType) {
		return fmt.Errorf("The devtype %s is not valid", devType)
	}
	d.DevType = devType
	return nil
}

/*SetAddress : Set the device address */
func (d *Device) SetAddress(address string) error {
	_config := config.GetConfig()
	if address == "" {
		d.Address = ""
		return nil
	}
	if !tools.IsValidAddr(address) {
		return fmt.Errorf("The address %s is not valid", address)
	}
	if address == _config.XAALBcastAddr {
		return errors.New("The broadcast address is reserved")
	}
	d.Address = address
	return nil
}

// GetTimeout : return Alive timeout used for isAlive msg
func (d Device) GetTimeout() int {
	return 2 * d.alivePeriod
}
