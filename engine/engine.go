package engine

import (
	"log"
	configmanager "xaal-go/configManager"
	"xaal-go/device"
	"xaal-go/message"
	"xaal-go/messagefactory"
	"xaal-go/network"
)

var _config configmanager.XaalConfiguration

/*InitWithConfig : init the engine using the config file parameters */
func InitWithConfig() {
	_config := configmanager.GetXAALConfig()

	Init(_config.Interface, _config.Address, _config.Port, _config.Hops, _config.Key)
}

/*Init : init the engine */
func Init(ifaceName string, address string, port uint16, hops uint8, key string) {
	_rxHandlers = append(_rxHandlers, handleRequest) // message receive workflow
	/* start network */
	network.Init(ifaceName, address, port, hops)
	/* start msg worker */
	messagefactory.Init(key)
}

/*******************
* Mainloops & run ..
********************/
var _running = make(chan bool)
var _started = false // engine is running or not

/* Start the core engine: send queue alive msg */
func start() {
	if !_started {
		network.Connect()
		_started = true
	}
}

// Stop all mainloops
func Stop() {
	log.Println("Stopping...")
	close(_queueMsgTx)
	_tickerAlive.Stop() // Stop Alives
	_running <- false
}

// Run all mainloops
func Run() {
	start()
	/* Process xAAL msg received, filter msg and process request*/
	go processRxMsg()
	/* Process xAAL msgs to send */
	go processTxMsg()
	/* Process attributes change for devices*/
	//TODO	self.process_attributesChange()
	/* Process timers */
	go processTimers()
	// Process Alives
	go processAlives()
	<-_running // Listen the channel to stop
	log.Println("Stopped")
}

// loop throught the devices, to find which are
// expected w/ the msg
// - Filter on devTypes for isAlive request.
// - Filter on device address
func filterMsgForDevices(msg *message.Message, devices []*device.Device) []*device.Device {
	var results []*device.Device
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
		/*
			if len(msg.Targets) == 0 { // if target list is empty == broadcast
				results = devices
			} else {
				for i := 0; i < len(devices); i++ {
					for i := range msg.Targets {
						if msg.Targets[i] == devices[i].Address {
							results = append(results, devices[i])
						}
					}
				}
			}*/
	}
	return results
}
