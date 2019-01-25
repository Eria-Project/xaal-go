// Package engine : Alive messages
package engine

import (
	"time"
	"xaal-go/configmanager"
	"xaal-go/device"
	"xaal-go/messagefactory"

	"github.com/Eria-Project/logger"
)

var _tickerAlive *time.Ticker

//TODO		self.__alives = []                       # list of alive devices

// SendAlive : Send a Alive message for a given device
func SendAlive(dev *device.Device) {
	timeout := dev.GetTimeout()
	msg, err := messagefactory.BuildAliveFor(dev, timeout)
	if err != nil {
		logger.Error("Cannot build alive message", logger.Fields{"-module": "engine", "err": err})
	} else {
		logger.Debug("Sending alive message", logger.Fields{"-module": "engine", "from": dev.Address})
		_queueMsgTx <- msg
	}
}

func sendAlives() {
	for _, dev := range _devices {
		SendAlive(dev)
	}
}

// SendIsAlive : Send a isAlive message, w/ devTypes filtering
func SendIsAlive(dev *device.Device, devTypes string) {
	body := make(map[string]interface{})
	body["devTypes"] = devTypes
	msg, err := messagefactory.BuildMsg(dev, []string{}, "request", "isAlive", body)
	if err != nil {
		logger.Error("Cannot build isAlive message", logger.Fields{"-module": "engine", "err": err})
	} else {
		_queueMsgTx <- msg
	}
}

// processAlives : Periodic sending alive messages
func processAlives() {
	_config := configmanager.GetXAALConfig()
	_tickerAlive = time.NewTicker(time.Duration(_config.AliveTimer) * time.Second)
	go func() {
		logger.Debug("Send initial alive messages", logger.Fields{"-module": "engine"})
		sendAlives()
		for range _tickerAlive.C {
			logger.Debug("Send alive messages", logger.Fields{"-module": "engine"})
			sendAlives()
		}
	}()
}
