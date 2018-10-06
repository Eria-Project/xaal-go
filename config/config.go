package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

/*XaalConfiguration : structure for xaal config file */
type XaalConfiguration struct {
	StackVersion string // default = "0.5"
	Interface    string
	Address      string
	Port         uint16
	Hops         uint8
	Key          string
	CipherWindow uint16
	AliveTimer   uint16
}

var instance *XaalConfiguration
var once sync.Once

/*GetConfig : Get the config instance*/
func GetConfig() *XaalConfiguration {
	once.Do(func() {
		config, err := loadConfig("xaal.json")
		if err != nil {
			panic(err)
		}
		instance = &config
		instance.StackVersion = "0.5"
	})
	return instance
}

/*loadConfig : Load config from file */
func loadConfig(filename string) (XaalConfiguration, error) {
	log.Printf("Loading config from %s", filename)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return XaalConfiguration{}, err
	}

	var c XaalConfiguration
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return XaalConfiguration{}, err
	}

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
