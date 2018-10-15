package device

import (
	"errors"
	"fmt"
)

// Attribute : xAAL internal attributes
type Attribute struct {
	Name         string
	defaultValue interface{}
	Value        interface{}
	Device       *Device
	Changed      bool
}

// NewAttribute : Create a new attribute
func (d *Device) NewAttribute(name string, defaultValue interface{}) (*Attribute, error) {
	if name == "" {
		return nil, errors.New("No name has been provided for attribute")
	}
	attr := Attribute{
		Name:         name,
		defaultValue: defaultValue,
		Changed:      false,
		Device:       d,
	}
	d.Attributes[name] = &attr
	return &attr, nil
}

// AddUnsupportedAttribute : Add the attribute string to the list of unsupported
func (d *Device) AddUnsupportedAttribute(name string) {
	d.unsupportedAttributes = append(d.unsupportedAttributes, name)
	delete(d.Attributes, name)
}

// SetAttributeValue : Set the attribute value
func (d *Device) SetAttributeValue(name string, value interface{}) error {
	if _, exists := d.Attributes[name]; !exists {
		return fmt.Errorf("Attribute %s doesn't exsists", name)
	}
	attr := d.Attributes[name]
	if attr.Value != value {
		attr.Value = value
		attr.Changed = true
	}

	return nil
}

/*
def get_attribute(self,name):
	for attr in self.__attributes:
		if attr.name == name:
			return attr
	return None
*/
