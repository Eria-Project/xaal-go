package engine

import (
	"github.com/project-eria/xaal-go/device"
	"github.com/project-eria/xaal-go/messagefactory"
	"github.com/project-eria/xaal-go/network"

	"github.com/project-eria/config-manager"
	"github.com/project-eria/logger"
)

// Version returns the current implementation version
func Version() string {
	return "0.0.1-dev"
}

var (
	_config = struct {
		StackVersion  string `default:"0.5"`          // protocol version
		Address       string `default:"224.0.29.200"` // mcast address
		Port          uint16 `default:"1235"`         // mcast port
		Hops          uint8  `default:"10"`           // mcast hop
		Key           string `required:"true"`
		CipherWindow  uint16 `default:"120"` // Time Window in seconds to avoid replay attacks
		AliveTimer    uint16 `default:"60"`  // Time between two alive msg
		XAALBcastAddr string `default:"00000000-0000-0000-0000-000000000000"`
	}{}
	_configFile = "xaal.json"
)

func init() {
}

// Init : init the engine using the config
func Init() {

	configManagerXAAL, err := configmanager.Init(_configFile, &_config)
	if err != nil {
		logger.Module("engine").WithError(err).WithField("filename", _configFile).Fatal()
	}

	if err := configManagerXAAL.Load(); err != nil {
		logger.Module("engine").WithError(err).Fatal()
	}

	messagefactory.Init(_config.StackVersion, _config.Key, _config.CipherWindow)
	device.Init(_config.XAALBcastAddr, _config.AliveTimer)
	_rxHandlers = append(_rxHandlers, handleRequest)          // message receive workflow
	network.Init(_config.Address, _config.Port, _config.Hops) // Start network
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
