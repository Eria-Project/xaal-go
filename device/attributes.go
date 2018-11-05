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

// GetMethods : return the list of device methods
func (d *Device) GetMethods() map[string]func(*Device, map[string]interface{}) map[string]interface{} {
	return d.Methods
}

// GetMethodArgs : return the list on arguments for a given method
func (d *Device) GetMethodArgs(method string) ([]string, error) {
	if _, in := d.MethodArgs[method]; !in {
		return nil, fmt.Errorf("Method not found")
	}
	return d.MethodArgs[method], nil
}

// default public methods

func getDescription(d *Device, args map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	if d.VendorID != "" {
		result["vendorId"] = d.VendorID
	}
	if d.ProductID != "" {
		result["productId"] = d.ProductID
	}
	if d.Version != "" {
		result["version"] = d.Version
	}
	if d.URL != "" {
		result["url"] = d.URL
	}
	if d.Info != "" {
		result["info"] = d.Info
	}
	if d.HwID != "" {
		result["hwId"] = d.HwID
	}
	if d.GroupID != "" {
		result["groupId"] = d.GroupID
	}

	result["unsupportedMethods"] = d.unsupportedMethods
	result["unsupportedNotifications"] = d.unsupportedNotifications
	result["unsupportedAttributes"] = d.unsupportedAttributes
	return result
}

// attributes:
// - None = body empty and means request all attributes
// - Empty array means request all attributes
// - Array of attributes (string) and means request attributes in the list
//
// TODO: (Waiting for spec. decision) add test on attribute devices
// - case physical sensor not responding or value not ready add error
// with specific error code and with value = suspicious/stale/cached
func getAttributes(d *Device, args map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	/* TODO
	dev_attr = {attr.name: attr for attr in self.__attributes}
	if _attributes:
		"""Process attributes filter"""
		for attr in _attributes:
			if attr in dev_attr.keys():
				result.update({dev_attr[attr].name: dev_attr[attr].value})
			else:
				logger.debug("Attribute %s not found" % attr)
	else:
		"""Process all attributes"""
		for attr in dev_attr.values():
			result.update({attr.name: attr.value})
	*/
	return result
}
