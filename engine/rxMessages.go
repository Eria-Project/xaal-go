// Package engine : xAAL messages rx handlers
package engine

import (
	"log"
	"xaal-go/device"
	"xaal-go/message"
	"xaal-go/messagefactory"
	"xaal-go/network"
)

var _rxHandlers []func(*message.Message)
var _queueMsgRx = make(chan []byte)

// receiveData : // Loop to received data
func receiveData() {
	for {
		_queueMsgRx <- network.GetData()
	}
}

// AddRxHandler : Add function to the list of msg handlers
func AddRxHandler(handler func(*message.Message)) {
	_rxHandlers = append(_rxHandlers, handler)
}

/*
def remove_rx_hanlder(self,func):
	self.rx_handlers.remove(func)
*/

// processRxMsg : process incomming messages
func processRxMsg() {
	go receiveData()

	for data := range _queueMsgRx {
		if data != nil {
			msg, err := messagefactory.DecodeMsg(data)
			if err != nil {
				log.Println(err)
			}
			if msg != nil {
				for i := 0; i < len(_rxHandlers); i++ {
					_rxHandlers[i](msg)
				}
			}
		}
	}
}

// handlerequest : Filter msg for devices according default xAAL API then process the
// request for each targets identied in the engine
func handleRequest(msg *message.Message) {
	if msg.IsRequest() {
		targets := filterMsgForDevices(msg, _devices)
		if len(targets) > 0 {
			processRequest(msg, targets)
		}
	}
}

// processRequest by device and add related response
// if reply necessary in the Tx fifo
// Note: xAAL attributes change are managed separately
func processRequest(msg *message.Message, targets []device.Device) {
	for i := 0; i < len(targets); i++ {
		target := targets[i]
		if msg.Header.Action == "isAlive" {
			SendAlive(target)
		} else {
			handleMethodRequest(msg, target)
		}
	}
}

// handleMethodRequest : Run method (xAAL exposed method) on device:
// - None is returned if device method do not return anything
// - result is returned if device method gives a response
// - Errors are raised if an error occured:
// 	* Internal error
//	* error returned on the xAAL bus
func handleMethodRequest(msg *message.Message, target device.Device) {
	/*
		try:
			result = run_action(msg, target)
			if result != None:
				self.send_reply(dev=target,targets=[msg.source],action=msg.action,body=result)
		except CallbackError as e:
			self.send_error(target, e.code, e.description)
		except XAALError as e:
			logger.error(e)
	*/
}
