package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"sync"
)

/*XaalConfiguration : structure for xaal config file */
type XaalConfiguration struct {
	StackVersion  string // protocol version
	Interface     string
	Address       string // mcast address
	Port          uint16 // mcast port
	Hops          uint8  // mcast hop
	Key           string
	CipherWindow  uint16 // Time Window in seconds to avoid replay attacks
	AliveTimer    uint16 // Time between two alive msg
	XAALBcastAddr string
}

var instance *XaalConfiguration
var once sync.Once

var defaultConfig = XaalConfiguration{
	StackVersion:  "0.5",
	Address:       "224.0.29.200",
	Port:          1235,
	Hops:          10,
	CipherWindow:  120,
	AliveTimer:    100,
	XAALBcastAddr: "00000000-0000-0000-0000-000000000000",
}

func (c *XaalConfiguration) mergeDefault() {
	if c.StackVersion == "" {
		c.StackVersion = defaultConfig.StackVersion
	}
	if c.Address == "" {
		c.Address = defaultConfig.Address
	}
	if c.Port == 0 {
		c.Port = defaultConfig.Port
	}
	if c.Hops == 0 {
		c.Hops = defaultConfig.Hops
	}
	if c.CipherWindow == 0 {
		c.CipherWindow = defaultConfig.CipherWindow
	}
	if c.AliveTimer == 0 {
		c.AliveTimer = defaultConfig.AliveTimer
	}
	if c.XAALBcastAddr == "" {
		c.XAALBcastAddr = defaultConfig.XAALBcastAddr
	}
}

/*GetConfig : Get the config instance*/
func GetConfig() *XaalConfiguration {
	once.Do(func() {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		configFile := fmt.Sprintf("%s/.xaal/xaal.json", usr.HomeDir)
		config, err := loadConfig(configFile)
		if err != nil {
			panic(err)
		}
		instance = &config
	})
	return instance
}

/*loadConfig : Load config from file */
func loadConfig(filename string) (XaalConfiguration, error) {
	log.Printf("Loading config from %s", filename)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return defaultConfig, err
	}

	var c XaalConfiguration
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return defaultConfig, err
	}

	// Load default values
	c.mergeDefault()

	return c, nil
}

/*
func saveConfig(c XaalConfiguration, filename string) error {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, 0644)
}
*/
