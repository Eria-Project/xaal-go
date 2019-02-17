package device

import (
	"errors"
	"fmt"
)

// Function declaration
type Function func(*Device, map[string]interface{}) map[string]interface{}

// Method : xAAL methods
type Method struct {
	Function Function
	Args     []string
}

// AddMethod : Associate a new method
func (d *Device) AddMethod(name string, f Function) (*Method, error) {
	if name == "" {
		return nil, errors.New("No name has been provided for the method")
	}
	d.Methods[name] = &Method{
		Function: f,
	}
	return d.Methods[name], nil
}

// GetMethods : return the list of device methods
func (d *Device) GetMethods() map[string]*Method {
	return d.Methods
}

// GetMethod : return the list on arguments for a given method
func (d *Device) GetMethod(name string) (*Method, error) {
	if _, in := d.Methods[name]; !in {
		return nil, fmt.Errorf("Method not found")
	}
	return d.Methods[name], nil
}
