package utils

import (
	// "errors"
	"fmt"
	// "net/http"
	//
	"github.com/vapor-ware/vesh/client"

	"github.com/urfave/cli"
	"github.com/spf13/viper"
)

type Config struct {
	VeshHost string `string:"vesh_host"`
	Debug bool	`bool:"debug"`
	ConfigFilePath string
}

func NewConfig(cli *cli.Context) error {
	config = new(Config struct)
	config, err = configFromDefault()
	config, err = configFromFile()
	config, err = configFromEnv(cli *cli.Context)
}

func configFromFile() error {
	return &(GetConfig())
}

func EvaluatePriority()


func GetConfig() (*Config, error) {
	v := viper.New()
	_ = readConfigFile(v)
	config, err := getConfigValuesFromFile(v)
	return config, err
}

func readConfigFile(v *viper.Viper) error {
	viper.SetConfigName(".vesh")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func getConfigValuesFromFile(v *viper.Viper) (*Config, error) {
	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		return &config, err
	}
	return &config, nil
}
