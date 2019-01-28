// Package engine : xAAL messages tx handlers
package engine

import (
	"xaal-go/device"
	"xaal-go/messagefactory"
	"xaal-go/network"

	"github.com/Eria-Project/logger"
)

var _queueMsgTx = make(chan []byte)

// processTxMsg : Process (send) message in tx queue chan called from other tasks
func processTxMsg() {
	for msg := range _queueMsgTx {
		network.SendData(msg)
	}
}

// SendRequest queue a new request
func SendRequest(dev *device.Device, targets []string, action string, body map[string]interface{}) {
	msg, err := messagefactory.BuildMsg(dev, targets, "request", action, body)
	if err != nil {
		logger.Module("engine").WithError(err).Error("Cannot build request message")
	} else {
		logger.Module("engine").WithFields(logger.Fields{"action": action, "from": dev.Address, "to": targets}).Debug("Sending request message")
		_queueMsgTx <- msg
	}
}

// SendReply queue a new reply
func SendReply(dev *device.Device, targets []string, action string, body map[string]interface{}) {
	msg, err := messagefactory.BuildMsg(dev, targets, "reply", action, body)
	if err != nil {
		logger.Module("engine").WithError(err).Error("Cannot build reply message")
	} else {
		logger.Module("engine").WithFields(logger.Fields{"action": action, "from": dev.Address, "to": targets}).Debug("Sending reply message")
		_queueMsgTx <- msg
	}
}

// SendError queue a error message
func SendError(dev *device.Device, errcode int, description string) {
	msg, err := messagefactory.BuildErrorMsg(dev, errcode, description)
	if err != nil {
		logger.Module("engine").WithError(err).Error("Cannot build error message")
	} else {
		logger.Module("engine").WithField("from", dev.Address).Debug("Sending error message")
		_queueMsgTx <- msg
	}
}

// SendGetDescription queue a getDescription request
func SendGetDescription(dev *device.Device, targets []string) {
	SendRequest(dev, targets, "getDescription", nil)
}

// SendGetAttributes queue a getAttributes request
func SendGetAttributes(dev *device.Device, targets []string) {
	SendRequest(dev, targets, "getAttributes", nil)
}

// SendNotification queue a notification message
func SendNotification(dev *device.Device, action string, body map[string]interface{}) {
	msg, err := messagefactory.BuildMsg(dev, []string{}, "notify", action, body)
	if err != nil {
		logger.Module("engine").WithError(err).Error("Cannot build notify message")
	} else {
		logger.Module("engine").WithFields(logger.Fields{"action": action, "from": dev.Address}).Debug("Sending notify message")
		_queueMsgTx <- msg
	}
}
