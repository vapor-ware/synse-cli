package config

import (
	"fmt"
	//"reflect"
	//"strings"

	log "github.com/Sirupsen/logrus"
	//"github.com/fatih/structs"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

type config struct {
	Debug bool
	ActiveHost *hostConfig
	Hosts map[string]*hostConfig
}

// AddHosts adds the given host to the configuration. If the host already exists
// in the config, an error is returned. If there is no current active host when
// a new host is being added, the new host will become the active host.
func (c *config) AddHost(host *hostConfig) error {
	if c.Hosts[host.Name] != nil {
		return fmt.Errorf("host '%v' already exists in configuration", host.Name)
	}
	c.Hosts[host.Name] = host
	if c.ActiveHost == nil {
		c.ActiveHost = host
	}
	return nil
}

type hostConfig struct {
	Name string
	Address string
}

// Config is a new variable containing the config object
var Config config


var configName = ".synse"


func NewHostConfig(name, address string) *hostConfig {
	return &hostConfig{
		Name: name,
		Address: address,
	}
}

// ConstructConfig takes in the cli context and builds the current config from
// the cascade of configuration sources. It prioritizes configruation options
// from sources in the following order, with top of the list being highest priority.
//
// 	- Run time CLI flags
// 	- Environment variables
// 	- Configuration files
// 		- .synse.yaml in the local directory
// 		- .synse.yaml in the home (~) directory
//
// All fields in the configuration file are optional.
func ConstructConfig(c *cli.Context) error {
	v := readConfigFromFile()

	err := v.Unmarshal(&Config)
	if err != nil {
		return err
	}

	// add a default "local" instance of Synse Server
	Config.AddHost(&hostConfig{
		Name: "local",
		Address: "localhost:5000",
	})

	// FIXME: not sure what this did..
	//s := structs.New(&Config)
	//for _, name := range c.GlobalFlagNames() {
	//	if !c.IsSet(name) {
	//		continue
	//	}
	//
	//	field := s.Field(strings.Replace(strings.Title(name), "-", "", -1))
	//
	//	val := reflect.ValueOf(c.Generic(name)).Elem()
	//
	//	var err error
	//	if val.Kind() == reflect.Bool {
	//		err = field.Set(val.Bool())
	//	} else {
	//		err = field.Set(val.String())
	//	}
	//
	//	if err != nil {
	//		fmt.Printf("%v\n", err)
	//	}
	//}

	log.WithFields(log.Fields{
		"config": fmt.Sprintf("%+v", Config),
	}).Debug("final config")

	return nil
}

// We don't care about being unable to read in the config as it is a non-terminal state.
// Log the issue as debug and move on.
func readConfigFromFile() *viper.Viper {
	v := viper.New()
	v.SetConfigName(configName)
	v.SetConfigType("yaml")

	v.AddConfigPath(".")      // Try local first
	v.AddConfigPath("$HOME/") // Then try home

	// Defaults
	v.SetDefault("debug", false)
	v.SetDefault("hosts", []hostConfig{})

	err := v.ReadInConfig()
	if err != nil {
		log.WithFields(log.Fields{
			"file": v.ConfigFileUsed(),
		}).Debug("config file not found, a new one will be created")
	}

	log.WithFields(log.Fields{
		"file":     v.ConfigFileUsed(),
		"settings": v.AllSettings(),
	}).Debug("loading config")

	return v
}
