package engine

import (
	"xaal-go/device"
	"xaal-go/messagefactory"
	"xaal-go/network"

	"github.com/Eria-Project/config-manager"
	"github.com/Eria-Project/logger"
)

var _config = struct {
	StackVersion  string `default:"0.5"`          // protocol version
	Address       string `default:"224.0.29.200"` // mcast address
	Port          uint16 `default:"1235"`         // mcast port
	Hops          uint8  `default:"10"`           // mcast hop
	Key           string `required:"true"`
	CipherWindow  uint16 `default:"120"` // Time Window in seconds to avoid replay attacks
	AliveTimer    uint16 `default:"60"`  // Time between two alive msg
	XAALBcastAddr string `default:"00000000-0000-0000-0000-000000000000"`
}{}

/*InitWithConfig : init the engine using the config file parameters */
func InitWithConfig(configFile string) {
	configManagerXAAL, err := configmanager.Init(configFile)
	if err != nil {
		logger.WithField("file", configFile).Fatal("Missing config file")
	}

	if err := configManagerXAAL.Load(&_config); err != nil {
		logger.WithError(err).Fatal()
	}
	messagefactory.Init(_config.StackVersion, _config.Key, _config.CipherWindow)
	device.Init(_config.XAALBcastAddr, _config.AliveTimer)
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
	logger.Module("engine").Info("Stopping...")
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
	go processAlives(_config.AliveTimer)
	<-_running // Listen the channel to stop
	logger.Module("engine").Info("Stopped")
}
