package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Persist writes the current configuration to file.
func Persist() error {

	// First, check if an existing configuration file exists.
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	// If an existing file does not exist, create one
	if configPath == "" {
		configPath = filepath.Join(os.Getenv("HOME"), configName+".yml")
	}

	log.WithFields(log.Fields{
		"path": configPath,
	}).Debug("Persisting config")

	data, err := yaml.Marshal(Config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(configPath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// getConfigPath
func getConfigPath(paths ...string) (string, error) {

	var found []string

	for _, path := range paths {

		if strings.HasPrefix(path, "$HOME") {
			path = os.Getenv("HOME") + path[5:]
		}

		// First check if a configuration file exists, either in the current working
		// directory or in the $HOME directory.
		fullPath := filepath.Join(path, configName)
		matches, err := filepath.Glob(fullPath + ".*")
		if err != nil {
			return "", err
		}

		for _, match := range matches {
			ext := filepath.Ext(match)

			switch strings.ToLower(ext) {
			case ".yaml", ".yml":
				found = append(found, match)

			default:
			}
		}
	}

	if len(found) == 0 {
		return "", nil
	}

	if len(found) >= 2 {
		return "", fmt.Errorf("found more than one possible configurations - can only have one")
	}

	// Only one was found, so return it
	return found[0], nil
}
