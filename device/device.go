package device

import (
	"errors"
	"fmt"

	"github.com/project-eria/xaal-go/utils"
)

var (
	_xAALBcastAddr string
	_aliveTimer    uint16
)

// Init the global devices params
func Init(xAALBcastAddr string, aliveTimer uint16) {
	_xAALBcastAddr = xAALBcastAddr
	_aliveTimer = aliveTimer
}

// Device : xAAL device
type Device struct {
	DevType     string // xaal devtype
	Address     string // xaal addr
	alivePeriod uint16 // time in sec between two alive
	VendorID    string // vendor ID ie : ACME
	ProductID   string // product ID
	Version     string // product release
	URL         string // product URL
	Info        string // additionnal info
	HwID        string // hardware info
	GroupID     string // group devices

	// Unsupported stuffs
	unsupportedAttributes    []string
	unsupportedMethods       []string
	unsupportedNotifications []string

	// Default attributes & methods
	Attributes map[string]*Attribute
	Methods    map[string]*Method
}

// New : device constructor
func New(devType string, address string) (*Device, error) {
	if !utils.IsValidDevType(devType) {
		return nil, fmt.Errorf("The devtype %s is not valid", devType)
	}
	if address == "" {
		address = utils.GetRandomUUID()
	} else {
		if !utils.IsValidAddr(address) {
			return nil, fmt.Errorf("The address %s is not valid", address)
		}
		if address == _xAALBcastAddr {
			return nil, errors.New("The broadcast address is reserved")
		}
	}

	device := Device{
		DevType:                  devType,
		Address:                  address,
		unsupportedAttributes:    []string{},
		unsupportedMethods:       []string{},
		unsupportedNotifications: []string{},
		Attributes:               make(map[string]*Attribute),
		Methods: map[string]*Method{
			"getAttributes": &Method{
				Function: getAttributes,
				Args:     []string{"attributes"},
			},
			"getDescription": &Method{
				Function: getDescription,
			},
		},
		alivePeriod: _aliveTimer,
	}
	return &device, nil
}

// SetDevType : Set the device type
// TO REMOVE ??
func (d *Device) SetDevType(devType string) error {
	if !utils.IsValidDevType(devType) {
		return fmt.Errorf("The devtype %s is not valid", devType)
	}
	d.DevType = devType
	return nil
}

/*SetAddress : Set the device address */
func (d *Device) SetAddress(address string, xAALBcastAddr string) error {
	if address == "" {
		d.Address = ""
		return nil
	}
	if !utils.IsValidAddr(address) {
		return fmt.Errorf("The address %s is not valid", address)
	}
	if address == xAALBcastAddr {
		return errors.New("The broadcast address is reserved")
	}
	d.Address = address
	return nil
}

// GetTimeout : return Alive timeout used for isAlive msg
func (d *Device) GetTimeout() uint16 {
	return (2 * d.alivePeriod)
}
