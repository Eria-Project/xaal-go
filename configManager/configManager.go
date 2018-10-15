package configmanager

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
	Address       string // mcast address
	Port          uint16 // mcast port
	Hops          uint8  // mcast hop
	Key           string
	CipherWindow  uint16 // Time Window in seconds to avoid replay attacks
	AliveTimer    uint16 // Time between two alive msg
	XAALBcastAddr string
}

var once sync.Once

var config = XaalConfiguration{
	StackVersion:  "0.5",
	Address:       "224.0.29.200",
	Port:          1235,
	Hops:          10,
	CipherWindow:  120,
	AliveTimer:    100,
	XAALBcastAddr: "00000000-0000-0000-0000-000000000000",
}

/*GetXAALConfig : Get the config instance*/
func GetXAALConfig() *XaalConfiguration {
	once.Do(func() {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		configFile := fmt.Sprintf("%s/.xaal/xaal.json", usr.HomeDir)
		err = LoadConfig(configFile, &config)
		if err != nil {
			panic(err)
		}
	})
	return &config
}

// LoadConfig : Load config from file
func LoadConfig(filename string, s interface{}) error {
	log.Printf("Loading config from %s", filename)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		// TODO What to do if file doesn't exists
		return err
	}

	err = json.Unmarshal(bytes, s)
	if err != nil {
		// TODO What to do if not json file
		return err
	}
	return nil
}

/*SaveConfig : Save config to file */
func SaveConfig(filename string, s interface{}) error {
	bytes, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, 0644)
}
