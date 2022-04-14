package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// LoadConfig
// load config file
func LoadConfig(config interface{}, configFile string) error {
	file, err := os.OpenFile(configFile, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	str, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(str, config); err != nil {
		return err
	}
	return nil
}
