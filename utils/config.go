package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	VaporHost      string `string:"vesh_host"`
	Debug          bool   `bool:"debug"`
	ConfigFilePath string
}

var VaporHost = ""
var DebugFlag = false
var ConfigFilePath = ""

func ConstructConfig() error {
	config := new(Config)
	v, err := readConfigFromFile()
	if err != nil {
		return err
	}
	err = v.Unmarshal(config)
	if err != nil {
		return err
	}
	switch {
	case config.VaporHost != "" && VaporHost == "":
		VaporHost = config.VaporHost
	case config.Debug && !DebugFlag:
		DebugFlag = config.Debug
	case config.ConfigFilePath != "" && ConfigFilePath == "":
		ConfigFilePath = config.ConfigFilePath
	}
	fmt.Println("populated config", VaporHost, DebugFlag, ConfigFilePath, config)
	return nil
}

func readConfigFromFile() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(".vesh")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")      // Try local first
	v.AddConfigPath("$HOME/") // Then try home
	err := v.ReadInConfig()
	if err != nil {
		return v, err
	}
	return v, nil
}
