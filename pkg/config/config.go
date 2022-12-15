// Synse CLI
// Copyright (c) 2019 Vapor IO
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package config

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"

	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func init() {
	config = Config{
		Contexts:       []ContextRecord{},
		CurrentContext: map[string]string{},
	}
}

// configFile is the name of file which will hold the persisted CLI
// configuration.
const configFile = ".synse.yml"

// config is an internal instance of the CLI configuration. When the
// CLI is run, it will load the configuration (via Load) into this
// variable so all sub-commands can access it.
var config Config

// Config specifies the persisted configuration for the CLI.
type Config struct {
	Contexts       []ContextRecord   `json:"contexts" yaml:"contexts" mapstructure:"contexts"`
	CurrentContext map[string]string `json:"current_context" yaml:"current_context" mapstructure:"current_context"`
}

// ContextRecord describes the record for a Synse component
// and how the CLI should connect to it.
type ContextRecord struct {
	Name    string  `json:"name" yaml:"name" mapstructure:"name"`
	Type    string  `json:"type" yaml:"type" mapstructure:"type"`
	Context Context `json:"context" yaml:"context" mapstructure:"context"`
}

// Context specifies any contextual information associated
// with a ContextRecord that can be used by the CLI to connect
// to the Synse component.
type Context struct {
	Address    string `json:"address" yaml:"address" mapstructure:"address"`
	ClientCert string `json:"client_cert" yaml:"client_cert" mapstructure:"client_cert"`
}

// Load loads the configuration for the CLI. If a configuration file
// cannot be found, this will load a new empty Config instance.
func Load() error {
	log.Debug("loading cli configuration")
	v := readConfigFromFile()

	// Use mapstructure.Decode here instead of Viper's built-in Unmarshal
	// since that appears to be broken and does not properly load all keys.
	err := mapstructure.Decode(v.AllSettings(), &config)
	if err != nil {
		return err
	}

	log.WithField("config", config).Debug("unmarshaled config")
	return nil
}

// Persist saves the CLI configuration to disk, writing back the
// in-memory config to the configuration YAML file.
func Persist() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	var configPath = filepath.Join(wd, configFile)
	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		log.Debug("persisting cli config to HOME directory")
		configPath = filepath.Join(home, configFile)
	} else if err != nil {
		log.Debug(fmt.Sprintf("stat cli config: %v", err))
		return err
	}

	log.WithFields(log.Fields{
		"path":   configPath,
		"config": fmt.Sprintf("%+v", config),
	}).Debug("persisting config")

	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	log.Debug("cli configuration persisted")
	return os.WriteFile(configPath, data, 0644)
}

// GetContexts gets the contexts for the default configuration.
func GetContexts() []ContextRecord {
	return config.Contexts
}

// GetContext gets the named context for the config. If a context
// with the given name does not exist, nil is returned.
func (c *Config) GetContext(name string) *ContextRecord {
	for _, ctx := range c.Contexts {
		if ctx.Name == name {
			return &ctx
		}
	}
	return nil
}

// GetContext gets the named context from the default configuration.
func GetContext(name string) *ContextRecord {
	return config.GetContext(name)
}

// AddContext adds a context to the configuration. If a context with the
// same name already exists, this will return an error.
func (c *Config) AddContext(ctx *ContextRecord) error {
	var contextExists bool
	for _, context := range c.Contexts {
		if context.Name == ctx.Name {
			contextExists = true
			break
		}
	}
	if contextExists {
		return fmt.Errorf("cannot add context '%s': name already exists", ctx.Name)
	}

	config.Contexts = append(config.Contexts, *ctx)
	log.WithField("context", ctx.Name).Debug("added context")
	return nil
}

// AddContext adds a context to the default configuration.
func AddContext(ctx *ContextRecord) error {
	return config.AddContext(ctx)
}

// RemoveContext removes a context from the configuration. If the given
// name does not correspond to a context, this has no effect.
//
// If the context being removed is the current context, the current context
// will be cleared.
func (c *Config) RemoveContext(name string) {
	if name == "" {
		log.Debug("no context to remove")
		return
	}

	var context ContextRecord
	var idx *int
	for i, ctx := range c.Contexts {
		if ctx.Name == name {
			context = ctx
			idx = &i
			break
		}
	}
	if idx != nil {
		c.Contexts = append(c.Contexts[:*idx], c.Contexts[*idx+1:]...)

		if c.CurrentContext[context.Type] == context.Name {
			delete(c.CurrentContext, context.Type)
		}
	}
	log.WithField("context", name).Debug("removed context")
}

// RemoveContext removes a context from the default configuration.
func RemoveContext(name string) {
	config.RemoveContext(name)
}

// Purge removes all contexts from the config and clears the current
// context.
func (c *Config) Purge() {
	c.CurrentContext = map[string]string{}
	c.Contexts = []ContextRecord{}
	log.Debug("cli contexts purged")
}

// Purge removes all contexts from the default configuration.
func Purge() {
	config.Purge()
}

// IsCurrentContext checks if the specified ContextRecord is currently active.
func (c *Config) IsCurrentContext(ctx *ContextRecord) bool {
	current, ok := c.CurrentContext[ctx.Type]
	if !ok {
		return false
	}
	return current == ctx.Name
}

// IsCurrentContext checks if the context is the current context for the
// default configuration.
func IsCurrentContext(ctx *ContextRecord) bool {
	return config.IsCurrentContext(ctx)
}

// SetCurrentContext sets the named context as the current active context. If
// the given name does not correspond to a ContextRecord, an error is returned.
func (c *Config) SetCurrentContext(name string) error {
	var context *ContextRecord
	for _, ctx := range c.Contexts {
		if ctx.Name == name {
			context = &ctx
			break
		}
	}
	if context == nil {
		return fmt.Errorf("cannot set '%s' as current context: no such context", name)
	}

	c.CurrentContext[context.Type] = context.Name
	log.WithField("context", name).Debug("set current context")
	return nil
}

// SetCurrentContext sets the current context for the default configuration.
func SetCurrentContext(name string) error {
	return config.SetCurrentContext(name)
}

// GetCurrentContext gets the ContextRecords for the current context, if set.
func (c *Config) GetCurrentContext() map[string]*ContextRecord {
	var current = make(map[string]*ContextRecord)

	if c.CurrentContext == nil || len(c.CurrentContext) == 0 {
		log.Debug("config has no current context")
		return current
	}

	for t, name := range c.CurrentContext {
		for _, ctx := range c.Contexts {
			if ctx.Type == t && ctx.Name == name {
				current[t] = &ctx
				break
			}
		}
	}
	return current
}

// GetCurrentContext gets the current context for the default configuration.
func GetCurrentContext() map[string]*ContextRecord {
	return config.GetCurrentContext()
}

// readConfigFromFile reads in the CLI configuration from file.
//
// If the config file does not exist, we fall back to using default values.
func readConfigFromFile() *viper.Viper {
	v := viper.New()
	v.SetConfigName(".synse")
	v.SetConfigType("yaml")

	v.AddConfigPath(".")               // Try local first
	v.AddConfigPath(os.Getenv("HOME")) // Then try home

	// Defaults
	v.SetDefault("current_context", map[string]string{})
	v.SetDefault("contexts", []ContextRecord{})

	err := v.ReadInConfig()
	if err != nil {
		log.WithFields(log.Fields{
			"file": v.ConfigFileUsed(),
		}).Debug("config file not found, a new one will be created")
	}

	log.WithFields(log.Fields{
		"file":     v.ConfigFileUsed(),
		"settings": v.AllSettings(),
	}).Debug("loaded cli configuration")
	return v
}
