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
	Args     *[]string
}

// AddMethod : Associate a new method
func (d *Device) AddMethod(name string, f Function, args *[]string) (*Method, error) {
	if name == "" {
		return nil, errors.New("No name has been provided for the method")
	}
	d.Methods[name] = &Method{
		Function: f,
		Args:     args,
	}
	return d.Methods[name], nil
}

// HandleMethod : Replace the method handler
// TODO Allow to have more than 1 handler
func (d *Device) HandleMethod(name string, f Function) error {
	method, ok := d.Methods[name]
	if !ok {
		return fmt.Errorf("Method %s not found", name)
	}
	method.Function = f
	return nil
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
