// Package engine : Attributes changes
package engine

import (
	"github.com/project-eria/xaal-go/device"
)

var _queueAttributesChange = make(chan *device.Device)

// NotifyAttributesChange : send notification message for all changes attributes
func NotifyAttributesChange(dev *device.Device) {
	_queueAttributesChange <- dev
}

/* TO REMOVE
def add__queueAttributesChange(self, attr):
"""add a new attribute change to the list"""
self.___queueAttributesChange.append(attr)

def get__queueAttributesChange(self):
"""return the pending attributes changes list"""
return self.___queueAttributesChange
*/

// processAttributesChange : Processes (send notify) attributes changes for all devices
func processAttributesChange() {
	for dev := range _queueAttributesChange {
		var numChanges int
		// Build Body
		body := make(map[string]interface{})
		for _, attr := range dev.Attributes {
			if attr.Changed {
				body[attr.Name] = attr.Value
				attr.Changed = false
				numChanges++
			}
		}
		if numChanges > 0 {
			SendNotification(dev, "attributesChange", body)
		}
	}
}
