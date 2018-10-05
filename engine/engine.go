package engine

import (
	"log"
	"xaal-go/config"
	"xaal-go/device"
	"xaal-go/message"
	"xaal-go/messagefactory"
	"xaal-go/network"
)

var _config config.XaalConfiguration
var _running = false         // engine started or not
var _started = false         // engine is running or not
var _devices []device.Device // list of devices / use (un)register_devices()
var _rxHandlers []func(*message.Message)

/*InitWithConfig : init the engine using the config file parameters */
func InitWithConfig() {
	_config := config.GetConfig()

	Init(_config.Interface, _config.Address, _config.Port, _config.BindAddress, _config.Hops, _config.Key)
}

/*Init : init the engine */
func Init(ifaceName string, address string, port uint16, bindAddress string, hops uint8, key string) {
	_rxHandlers = append(_rxHandlers, handleRequest) // message receive wrOkflow
	/*
		self.timers = []                         # functions to call periodic
		self.__last_timer = 0                    # last timer check

		self.__attributesChange = []             # list of XAALAttributes instances
		self.__txFifo = collections.deque()      # tx msg fifo
		self.__alives = []                       # list of alive devices
	*/

	/* start network */
	network.Init(ifaceName, address, port, hops, bindAddress)
	/* start msg worker */
	messagefactory.Init(key)
}

/*******************
* Mainloops & run ..
********************/

// loop : Process incomming xAAL msg
// Process attributes change for device
// Process timers
// Process isAlive for device
// Send msgs from the Tx Buffer
func loop() {
	/* Process xAAL msg received, filter msg and process request*/
	processRxMsg()
	/* Process attributes change for devices*/
	//TODO	self.process_attributesChange()
	/* Process timers */
	//TODO	self.process_timers()
	/* Process Alives */
	//TODO	self.process_alives()
	/* Process xAAL msgs to send */
	//TODO	self.process_tx_msg()
}

/* Start the core engine: send queue alive msg */
func start() {
	if !_started {
		network.Connect()
		for i := 0; i < len(_devices); i++ {
			//			dev := _devices[i]
			//	self.send_alive(dev)
			//	dev.update_alive()
		}
		_started = true
	}
}

func stop() {
	_running = false
}

// Run is a var
func Run() {
	start()
	_running = true
	for _running {
		loop()
	}
}

// loop throught the devices, to find which are
// expected w/ the msg
// - Filter on devTypes for isAlive request.
// - Filter on device address
func filterMsgForDevices(msg *message.Message, devices []device.Device) []device.Device {
	var results []device.Device
	if msg.IsAlive() {
		/* TODO
		if _, in := msg.Body["devTypes"]; in {
			devTypes := msg.Body["devTypes"].(string)
			if _, in := devTypes["any.any"]; in {
				results = devices
			} else {
				for i := 0; i < len(devices); i++ {
					devTypesSplit := strings.Split(devices[i].DevType(), ".")
					anySubtype := fmt.Sprintf("%s.any", devTypesSplit[0])
					if _, in := devTypes[devices[i].DevType()]; in {
						results = append(results, devices[i])
					} else if _, in := devTypes[anySubtype]; in {
						results = append(results, devices[i])
					}
				}
			}
		}
		*/
	} else {
		if len(msg.Targets) == 0 { // if target list is empty == broadcast
			results = devices
		} else {
			for i := 0; i < len(devices); i++ {
				for i := range msg.Targets {
					if msg.Targets[i] == devices[i].Address() {
						results = append(results, devices[i])
					}
				}
			}
		}
	}
	return results
}

/**************************
* xAAL messages rx handlers
**************************/
/*receiveMsg : return new received message or nil */
func receiveMsg() *message.Message {
	var msg *message.Message
	data := network.GetData()
	if data != nil {
		var err error
		msg, err = messagefactory.DecodeMsg(data)
		if err != nil {
			log.Println(err)
			return nil
		}
	}
	return msg
}

/*AddRxHandler : Add function to the list of msg handlers */
func AddRxHandler(handler func(*message.Message)) {
	_rxHandlers = append(_rxHandlers, handler)
}

/*
def remove_rx_hanlder(self,func):
	self.rx_handlers.remove(func)
*/

/*processRxMsg : process incomming messages */
func processRxMsg() {
	msg := receiveMsg()
	if msg != nil {
		for i := 0; i < len(_rxHandlers); i++ {
			_rxHandlers[i](msg)
		}
	}
}

/*handlerequest : Filter msg for devices according default xAAL API then process the
request for each targets identied in the engine
*/
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
		if msg.Action() == "isAlive" {
			// TODO	self.send_alive(target)
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
