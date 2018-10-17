// Package engine : Alive messages
package engine

import (
	"time"
	"xaal-go/configmanager"
	"xaal-go/device"
	"xaal-go/messagefactory"

	"xaal-go/log"
)

var _tickerAlive *time.Ticker

//TODO		self.__alives = []                       # list of alive devices

// SendAlive : Send a Alive message for a given device
func SendAlive(dev *device.Device) {
	timeout := dev.GetTimeout()
	msg, err := messagefactory.BuildAliveFor(dev, timeout)
	if err != nil {
		log.Error("Cannot build alive message", log.Fields{"-module": "engine", "err": err})
	} else {
		_queueMsgTx <- msg
	}
}

func sendAlives() {
	for i := 0; i < len(_devices); i++ {
		dev := _devices[i]
		SendAlive(dev)
	}
}

// SendIsAlive : Send a isAlive message, w/ devTypes filtering
func SendIsAlive(dev *device.Device, devTypes string) {
	body := make(map[string]interface{})
	body["devTypes"] = devTypes
	msg, err := messagefactory.BuildMsg(dev, []string{}, "request", "isAlive", body)
	if err != nil {
		log.Error("Cannot build isAlive message", log.Fields{"-module": "engine", "err": err})
	} else {
		_queueMsgTx <- msg
	}
}

// processAlives : Periodic sending alive messages
func processAlives() {
	_config := configmanager.GetXAALConfig()
	_tickerAlive = time.NewTicker(time.Duration(_config.AliveTimer) * time.Second)
	go func() {
		log.Debug("Send initial alive messages", log.Fields{"-module": "engine"})
		sendAlives()
		for range _tickerAlive.C {
			log.Debug("Send alive messages", log.Fields{"-module": "engine"})
			sendAlives()
		}
	}()
}
