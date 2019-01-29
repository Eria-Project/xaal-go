// Package engine : Devices Management
package engine

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Eria-Project/xaal-go/device"
	"github.com/Eria-Project/xaal-go/message"
	"github.com/Eria-Project/xaal-go/tools"
)

var _devices map[string]*device.Device // list of devices / use (un)register_devices()

func init() {
	// Initialize the map of devices
	_devices = make(map[string]*device.Device)
}

// AddDevice : register a new device
func AddDevice(dev *device.Device) {
	_devices[dev.Address] = dev
}

// AddDevices : register new devices
func AddDevices(devs []*device.Device) {
	for _, dev := range devs {
		AddDevice(dev)
	}
}

// RemoveDevice : unregister a device
func RemoveDevice(dev *device.Device) {
	delete(_devices, dev.Address)
}

// loop throught the devices, to find which are
// expected w/ the msg
// - Filter on devTypes for isAlive request.
// - Filter on device address
func filterMsgForDevices(msg *message.Message, devices map[string]*device.Device) map[string]*device.Device {
	var results map[string]*device.Device
	if msg.Header.Action == "isAlive" {
		if _, in := msg.Body["devTypes"]; in {
			var devTypes []string
			if reflect.TypeOf(msg.Body["devTypes"]).String() == "string" {
				// Convert into array
				devTypes = []string{msg.Body["devTypes"].(string)}
			} else { // Array
				devTypes = msg.Body["devTypes"].([]string)
			}
			if _, in := tools.SliceContains(&devTypes, "any.any"); in {
				// If request alive for all devices
				results = devices
			} else {
				results = map[string]*device.Device{}
				for _, dev := range devices {
					devTypesSplit := strings.Split(dev.DevType, ".")
					anySubtype := fmt.Sprintf("%s.any", devTypesSplit[0])
					if _, in := tools.SliceContains(&devTypes, dev.DevType); in {
						results[dev.Address] = dev
					} else if _, in := tools.SliceContains(&devTypes, anySubtype); in {
						results[dev.Address] = dev
					}
				}
			}
		}
	} else {
		if len(msg.Targets) == 0 { // if target list is empty == broadcast
			results = devices
		} else {
			results = map[string]*device.Device{}
			for _, dev := range devices {
				for i := range msg.Targets {
					if msg.Targets[i] == dev.Address {
						results[dev.Address] = dev
					}
				}
			}
		}
	}
	return results
}
