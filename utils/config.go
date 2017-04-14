package utils

import (
	"fmt"
	"reflect"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/fatih/structs"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

type config struct {
	VaporHost string
	Debug     bool
	Config    string
}

var Config config

// ConstructConfig takes in the cli context and builds the current config from
// the cascade of configuration sources. It prioritizes configruation options
// from sources in the following order, with top of the list being highest priority.
//
// 	- Run time CLI flags
// 	- Environment variables
// 	- Configuration files
// 		- .vesh.yaml in the local directory
// 		- .vesh.yaml in the home (~) directory
//
// All fields in the configuration file are optional.
func ConstructConfig(c *cli.Context) error {
	v := readConfigFromFile()

	v.RegisterAlias("VaporHost", "vapor_host") // FIXME: This is really hacky, but works for now

	err := v.Unmarshal(&Config)
	if err != nil {
		return err
	}

	s := structs.New(&Config)
	for _, name := range c.GlobalFlagNames() {
		if !c.IsSet(name) {
			continue
		}

		field := s.Field(strings.Replace(strings.Title(name), "-", "", -1))

		val := reflect.ValueOf(c.Generic(name)).Elem()

		var err error
		if val.Kind() == reflect.Bool {
			err = field.Set(val.Bool())
		} else {
			err = field.Set(val.String())
		}

		if err != nil {
			fmt.Println("%v", err)
		}
	}

	log.WithFields(log.Fields{
		"config": fmt.Sprintf("%+v", Config),
	}).Debug("final config")

	return nil
}

// We don't care about being unable to read in the config as it is a non-terminal state.
// Log the issue as debug and move on.
func readConfigFromFile() *viper.Viper {
	v := viper.New()
	v.SetConfigName(".vesh")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")      // Try local first
	v.AddConfigPath("$HOME/") // Then try home

	// Defaults
	v.SetDefault("VaporHost", "demo.vapor.io")

	v.ReadInConfig()

	log.WithFields(log.Fields{
		"file":     v.ConfigFileUsed(),
		"settings": v.AllSettings(),
	}).Debug("loading config")

	return v
}
