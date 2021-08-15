package config

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Flags defines the flags for CrowsNest.
type Flags struct {
	RunOnce    bool
	ConfigPath string `validate:"omitempty,file"`
	Verbose    bool
	LogPath    string
}

// ValidateFlags validates the given Flags.
func ValidateFlags(f Flags) error {
	validate := validator.New()
	if errs := validate.Struct(f); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			return fmt.Errorf(
				"invalid value for flag %s",
				strings.ToLower(err.Field()),
			)
		}
	}
	return nil
}
