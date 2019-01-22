// Package engine : xAAL messages tx handlers
package engine

import (
	"xaal-go/device"
	"xaal-go/messagefactory"
	"xaal-go/network"

	"github.com/ERIA-Project/logger"
)

var _queueMsgTx = make(chan []byte)

// processTxMsg : Process (send) message in tx queue chan called from other tasks
func processTxMsg() {
	for msg := range _queueMsgTx {
		network.SendData(msg)
	}
}

// queue a new request
func sendRequest(dev *device.Device, targets []string, action string, body map[string]interface{}) {
	msg, err := messagefactory.BuildMsg(dev, targets, "request", action, body)
	if err != nil {
		logger.Error("Cannot build request message", logger.Fields{"-module": "engine", "err": err})
	} else {
		logger.Debug("Sending request message", logger.Fields{"-module": "engine", "action": action, "from": dev.Address, "to": targets})
		_queueMsgTx <- msg
	}
}

// queue a new reply
func sendReply(dev *device.Device, targets []string, action string, body map[string]interface{}) {
	msg, err := messagefactory.BuildMsg(dev, targets, "reply", action, body)
	if err != nil {
		logger.Error("Cannot build reply message", logger.Fields{"-module": "engine", "err": err})
	} else {
		logger.Debug("Sending reply message", logger.Fields{"-module": "engine", "action": action, "from": dev.Address, "to": targets})
		_queueMsgTx <- msg
	}
}

// queue a error message
func sendError(dev *device.Device, errcode int, description string) {
	msg, err := messagefactory.BuildErrorMsg(dev, errcode, description)
	if err != nil {
		logger.Error("Cannot build error message", logger.Fields{"-module": "engine", "err": err})
	} else {
		logger.Debug("Sending error message", logger.Fields{"-module": "engine", "from": dev.Address})
		_queueMsgTx <- msg
	}
}

// queue a getDescription request
func sendGetDescription(dev *device.Device, targets []string) {
	sendRequest(dev, targets, "getDescription", nil)
}

// queue a getAttributes request
func sendGetAttributes(dev *device.Device, targets []string) {
	sendRequest(dev, targets, "getAttributes", nil)
}

func sendNotification(dev *device.Device, action string, body map[string]interface{}) {
	msg, err := messagefactory.BuildMsg(dev, []string{}, "notify", action, body)
	if err != nil {
		logger.Error("Cannot build notify message", logger.Fields{"-module": "engine", "err": err})
	} else {
		logger.Debug("Sending notify message", logger.Fields{"-module": "engine", "action": action, "from": dev.Address})
		_queueMsgTx <- msg
	}
}
