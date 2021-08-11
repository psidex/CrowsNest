package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type RepositoryConfig struct {
	Directory string
	Remote    string // TODO: Can we find this in dir/.git?
	GitFlags  string
	Interval  int
}

type Config struct {
	Respositories map[string]RepositoryConfig
}

func Get(path string) Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")

	if path != "" {
		viper.AddConfigPath(path)
	}

	viper.SetDefault("Interval", 60)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var c Config

	if err := viper.Unmarshal(&c); err != nil {
		panic(fmt.Errorf("unable to decode into struct: %w", err))
	}

	return c
}
