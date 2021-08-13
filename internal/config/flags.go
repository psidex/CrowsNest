package config

import (
	"fmt"
	"log"
	"os"
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

// ValidateFlags validates the given Flags and sets up logging options.
// Returns the file (if there is one) that the log is being written to.
func ValidateFlags(f Flags) (*os.File, error) {
	validate := validator.New()

	errs := validate.Struct(f)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			return nil, fmt.Errorf(
				"invalid value for flag %s",
				strings.ToLower(err.Field()),
			)
		}
	}

	if f.LogPath != "" {
		file, err := os.OpenFile(f.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		log.SetOutput(file)
		return file, nil
	}

	return nil, nil
}
