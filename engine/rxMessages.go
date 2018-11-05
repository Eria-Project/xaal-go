// Package engine : xAAL messages rx handlers
package engine

import (
	"xaal-go/device"
	"xaal-go/message"
	"xaal-go/messagefactory"
	"xaal-go/network"

	"xaal-go/log"
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
				log.Error("Cannot decode message", log.Fields{"-module": "engine", "err": err})
			}
			if _, in := _devices[msg.Header.Source]; !in { // Ignore is the msg comes for one of our devices
				if msg != nil {
					for i := 0; i < len(_rxHandlers); i++ {
						_rxHandlers[i](msg)
					}
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
			log.Debug("Received request", log.Fields{"-module": "engine", "action": msg.Header.Action})
			processRequest(msg, targets)
		}
	}
}

// processRequest by device and add related response
// if reply necessary in the Tx fifo
// Note: xAAL attributes change are managed separately
func processRequest(msg *message.Message, targets map[string]*device.Device) {
	for _, target := range targets {
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
func handleMethodRequest(msg *message.Message, target *device.Device) {
	result := runAction(msg, target)
	if result != nil {
		sendReply(target, []string{msg.Header.Source}, msg.Header.Action, result)
		// TODO		except CallbackError as e:
		//			self.send_error(target, e.code, e.description)
		//		except XAALError as e:
		//			logger.error(e)

	}
}

// runAction: Extract & run an action (match with exposed method) from a msg
// on the selected device.
// Return:
//	- nil
//	- result from method if method return something
// Note: If action not found raise error, if wrong parameter raise error
func runAction(msg *message.Message, dev *device.Device) map[string]interface{} {
	methods := dev.GetMethods()
	params := make(map[string]interface{})
	var result map[string]interface{}
	if _, in := methods[msg.Header.Action]; in {
		method := methods[msg.Header.Action]
		if len(msg.Body) > 0 {
			methodParams, _ := dev.GetMethodArgs(msg.Header.Action) // TODO Handle error
			for _, p := range methodParams {
				if value, in := msg.Body[p]; in {
					params[p] = value
				} else {
					log.Info("Wrong method parameter for action", log.Fields{"parameter": p, "action": msg.Header.Action})
				}
			}
		}
		result = method(dev, params)
		//	 except :
		//		raise XAALError("Error in method:%s params:%s" % (msg.action,params))
	}
	//   else:
	//     raise XAALError("Method %s not found" % msg.action)
	return result
}
