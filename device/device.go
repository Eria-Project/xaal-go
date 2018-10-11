package device

import (
	"errors"
	"fmt"
	configmanager "xaal-go/configManager"
	"xaal-go/tools"
)

// Device : xAAL device
type Device struct {
	DevType     string // xaal devtype
	Address     string // xaal addr
	alivePeriod int    // time in sec between two alive
	VendorID    string // vendor ID ie : ACME
	ProductID   string // product ID
	Version     string // product release
	URL         string // product URL
	Info        string // additionnal info

	/*
		self.devtype = devtype          # xaal devtype
		self.address = addr             # xaal addr
		self.hw_id = None               # hardware info
		self.group_id = None            # group devices
		# Some useless attributes, only for compatibility
		self.bus_addr = config.address
		self.bus_port = config.port
		self.hops = config.hops
	*/
	// Unsupported stuffs
	unsupportedAttributes []string
	//	unsupportedMethods = []
	//	unsupportedNotifications = []

	// Default attributes & methods
	attributes []Attribute
	/*
		self.methods = {'getAttributes' : self._get_attributes,
						'getDescription': self._get_description }
		self.engine = engine
	*/
}

// Attribute : xAAL internal attributes
type Attribute struct {
	name         string
	defaultValue interface{}
	value        interface{}
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
	_config := configmanager.GetXAALConfig()
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

// ---------------- attributes ----------------

// NewAttribute : Create a new attribute
func (d *Device) NewAttribute(name string, defaultValue interface{}) Attribute {
	attr := Attribute{
		name:         name,
		defaultValue: defaultValue,
	}
	d.AddAttribute(&attr)
	return attr
}

// AddAttribute : Add the attribute to the device list of attributes
func (d *Device) AddAttribute(attr *Attribute) {
	if attr != nil {
		// TODO attr.device = self
		d.attributes = append(d.attributes, *attr)
	}
}

// AddUnsupportedAttribute : Add the attribute string to the list of unsupported
func (d *Device) AddUnsupportedAttribute(name string) {
	d.unsupportedAttributes = append(d.unsupportedAttributes, name)
	d.DelAttribute(name)
}

// DelAttribute : Remove the attribute from device list of attributes
func (d *Device) DelAttribute(name string) error {
	if name != "" {
		// Find element index
		i := d.findAttributeIndex(name)
		if i > -1 {
			// Delete index (See https://github.com/golang/go/wiki/SliceTricks)
			d.attributes = append(d.attributes[:i], d.attributes[i+1:]...)
			return nil
		}
	}
	return errors.New("Attribute not found")
}

// SetAttributeValue : Set the attribute value
func (d *Device) SetAttributeValue(name string, value interface{}) {
	// Search element index
	i := d.findAttributeIndex(name)
	if i > -1 {
		d.attributes[i].value = value
	}
}

func (d *Device) findAttributeIndex(name string) int {
	for i, e := range d.attributes {
		if e.name == name {
			return i
		}
	}
	return -1
}

/*
def get_attribute(self,name):
	for attr in self.__attributes:
		if attr.name == name:
			return attr
	return None
*/
