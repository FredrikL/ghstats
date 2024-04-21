package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Repo       []string `mapstructure:"repos"`
	Bots       []string `mapstructure:"bots"`
	OrgsToSkip []string `mapstructure:"orgskip"`
}

var config *Config

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		return
	}
}

func GetConfig() *Config {
	return config
}
