package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/viper"
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/fatih/structs"

)

type config struct {
	Host           string
	Debug          bool
	Config         string
}

var Config config

func ConstructConfig(c *cli.Context) error {
	v := readConfigFromFile()

	err := v.Unmarshal(&Config)
	if err != nil {
		return err
	}

	s := structs.New(&Config)
	for _, name := range c.GlobalFlagNames() {
		if !c.IsSet(name) { continue }

		field := s.Field(strings.Title(name))

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
		"config": Config,
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
	v.SetDefault("Host", "demo.vapor.io")

	v.ReadInConfig()

	log.WithFields(log.Fields{
		"file": v.ConfigFileUsed(),
		"settings": v.AllSettings(),
	}).Debug("loading config")

	return v
}
