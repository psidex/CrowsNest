package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type RepositoryConfig struct {
	Directory string `validate:"required,dir"`
	Remote    string `validate:"required"` // TODO: Can we find this in dir/.git?
	GitFlags  []string
	Interval  int `validate:"gt=0"`
	Method    string
}

type Config struct {
	Respositories map[string]*RepositoryConfig
}

func Get(path string) (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")

	if path != "" {
		viper.AddConfigPath(path)
	}

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("fatal error config file: %w", err)
	}

	var c Config

	if err := viper.Unmarshal(&c); err != nil {
		return Config{}, fmt.Errorf("unable to decode into struct: %w", err)
	}

	if len(c.Respositories) == 0 {
		return Config{}, errors.New("no repositoires found")
	}

	validate := validator.New()
	for cfgName, cfg := range c.Respositories {
		if cfg.Method == "" {
			cfg.Method = "pull"
		} else if cfg.Method != "pull" && cfg.Method != "checkpull" {
			return Config{}, errors.New("invalid configuration for key \"method\"")
		}

		if cfg.Interval == 0 {
			cfg.Interval = 60
		}

		err := validate.Struct(cfg)

		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				return Config{}, fmt.Errorf(
					"invalid configuration for respositories.%s.%s",
					cfgName,
					strings.ToLower(err.Field()),
				)
			}
		}

	}

	return c, nil
}
