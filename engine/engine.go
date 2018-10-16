package engine

import (
	"log"
	"xaal-go/configmanager"
	"xaal-go/device"
	"xaal-go/message"
	"xaal-go/network"
)

var _config = configmanager.GetXAALConfig()

/*InitWithConfig : init the engine using the config file parameters */
func InitWithConfig() {
	Init(_config.Address, _config.Port, _config.Hops)
}

/*Init : init the engine */
func Init(address string, port uint16, hops uint8) {
	_rxHandlers = append(_rxHandlers, handleRequest) // message receive workflow
	/* start network */
	network.Init(address, port, hops)
}

/*******************
* Mainloops & run ..
********************/
var _running = make(chan bool)

// IsStarted : if engine is running or not
var IsStarted = false

/* Start the core engine: send queue alive msg */
func start() {
	if !IsStarted {
		network.Connect()
		IsStarted = true
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
	go processAttributesChange()
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
