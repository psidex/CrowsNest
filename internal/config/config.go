package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type CliBinOpts struct {
	BinaryPath string `validate:"omitempty,file"`
	Flags      []string
	// NOTE: It seems it's not possible to validate as dir when using required_with.
	WorkingDirectory string `validate:"required_with=BinaryPath"`
}

type RepositoryConfig struct {
	Directory    string `validate:"required,dir"`
	GitPullFlags []string
	Interval     int `validate:"gt=0"`
	PrePullCmd   CliBinOpts
	PostPullCmd  CliBinOpts
}

type Config struct {
	Repositories map[string]*RepositoryConfig
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
		return Config{}, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	if len(c.Repositories) == 0 {
		return Config{}, errors.New("no repositoires found")
	}

	validate := validator.New()
	for cfgName, cfg := range c.Repositories {
		if cfg.Interval == 0 {
			cfg.Interval = 60
		}

		err := validate.Struct(cfg)

		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				return Config{}, fmt.Errorf(
					"invalid configuration for %s: %s",
					cfgName,
					strings.ToLower(err.Namespace()),
				)
			}
		}

	}

	return c, nil
}
