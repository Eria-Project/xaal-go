package device

import (
	"errors"
	"fmt"

	logger "github.com/project-eria/eria-logger"
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
	attributes, in := args["attributes"]
	if in && len(attributes.([]string)) > 0 {
		// Process attributes filter
		for _, name := range attributes.([]string) {
			if attr, in := d.Attributes[name]; in {
				result[name] = attr.Value
			} else {
				logger.Module("xaal:device").WithField("name", name).Debug("Attribute not found")
			}
		}
	} else {
		// Process all attributes
		for name, attr := range d.Attributes {
			result[name] = attr.Value
		}
	}
	return result
}
