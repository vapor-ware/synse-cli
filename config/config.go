package config

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

// CliConfig specifies the configuration for the CLI
type CliConfig struct {
	Debug      bool
	ActiveHost *HostConfig
	Hosts      map[string]*HostConfig
}

// AddHost adds the given host to the configuration. If the host already exists
// in the config, an error is returned. If there is no current active host when
// a new host is being added, the new host will become the active host.
func (c *CliConfig) AddHost(host *HostConfig) error {
	if c.Hosts[host.Name] != nil {
		return fmt.Errorf("host '%v' already exists in configuration", host.Name)
	}
	c.Hosts[host.Name] = host
	if c.ActiveHost == nil {
		c.ActiveHost = host
	}
	return nil
}

// HostConfig holds the configuration information for a single Synse Server host.
type HostConfig struct {
	Name    string `json:"name" yaml:"name"`
	Address string `json:"address" yaml:"address"`
}

// IsActiveHost checks if the host is the current active host for the CLI.
func (c *HostConfig) IsActiveHost() bool {
	return Config.ActiveHost != nil && *c == *Config.ActiveHost
}

// Config is a new variable containing the config object
var Config CliConfig

var configName = ".synse"

// NewHostConfig creates a new instance of HostConfig with the given values.
func NewHostConfig(name, address string) *HostConfig {
	return &HostConfig{
		Name:    name,
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
	Config.AddHost(&HostConfig{ // nolint
		Name:    "local",
		Address: "localhost:5000",
	})

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
	v.SetDefault("hosts", []HostConfig{})

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
