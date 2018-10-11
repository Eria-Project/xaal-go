// Package engine : Alive messages
package engine

import (
	"log"
	"time"
	configmanager "xaal-go/configManager"
	"xaal-go/device"
	"xaal-go/messagefactory"
)

var _tickerAlive *time.Ticker

//TODO		self.__alives = []                       # list of alive devices

// SendAlive : Send a Alive message for a given device
func SendAlive(dev *device.Device) {
	timeout := dev.GetTimeout()
	msg := messagefactory.BuildAliveFor(dev, timeout)
	_queueMsgTx <- msg
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
	msg := messagefactory.BuildMsg(dev, []string{}, "request", "isAlive", body)
	_queueMsgTx <- msg
}

// processAlives : Periodic sending alive messages
func processAlives() {
	_config := configmanager.GetXAALConfig()
	_tickerAlive = time.NewTicker(time.Duration(_config.AliveTimer) * time.Second)
	go func() {
		log.Println("Send initial alive messages")
		sendAlives()
		for range _tickerAlive.C {
			log.Println("Send alive messages")
			sendAlives()
		}
	}()
}
