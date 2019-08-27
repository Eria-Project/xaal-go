package engine

import (
	"github.com/project-eria/logger"
	"github.com/project-eria/xaal-go/device"
	"github.com/project-eria/xaal-go/messagefactory"
	"github.com/project-eria/xaal-go/network"
)

// XaalConfig is a struct that describe the xAAL config
type XaalConfig struct {
	StackVersion  string `default:"0.5"`          // protocol version
	Address       string `default:"224.0.29.200"` // mcast address
	Port          uint16 `default:"1235"`         // mcast port
	Hops          uint8  `default:"10"`           // mcast hop
	Key           string `required:"true"`
	CipherWindow  uint16 `default:"120"` // Time Window in seconds to avoid replay attacks
	AliveTimer    uint16 `default:"60"`  // Time between two alive msg
	XAALBcastAddr string `default:"00000000-0000-0000-0000-000000000000"`
}

var _config XaalConfig

func init() {
}

// Init : init the engine using the config
func Init(config XaalConfig) {
	_config = config
	messagefactory.Init(config.StackVersion, config.Key, config.CipherWindow)
	device.Init(config.XAALBcastAddr, config.AliveTimer)
	_rxHandlers = append(_rxHandlers, handleRequest)       // message receive workflow
	network.Init(config.Address, config.Port, config.Hops) // Start network
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
