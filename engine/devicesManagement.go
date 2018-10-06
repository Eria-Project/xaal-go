// Package engine : Devices Management
package engine

import "xaal-go/device"

var _devices []device.Device // list of devices / use (un)register_devices()

// AddDevice : register a new device
func AddDevice(dev device.Device) {
	for i := 0; i < len(_devices); i++ {
		if dev.Address == _devices[i].Address {
			return
		}
	}
	_devices = append(_devices, dev)
}

/*
def add_devices(self, devs):
"""register new devices"""
for dev in devs:
	self.add_device(dev)

def remove_device(self, dev):
"""unregister a device """
dev.engine = None
# Remove dev from devices list
self.devices.remove(dev)
*/
