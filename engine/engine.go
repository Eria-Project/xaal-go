package engine

import (
	"xaal-go/configmanager"
	"xaal-go/network"

	"github.com/Eria-Project/logger"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//	logger.SetFormatter(&logger.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	//logger.SetOutput(os.Stdout)
}

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
	logger.Info("Stopping...", logger.Fields{"-module": "engine"})
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
	logger.Info("Stopped", logger.Fields{"-module": "engine"})
}
