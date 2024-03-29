package configurator

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// NewConfiguratorFromFile reads configuration values from a .env file given a path.
func NewConfiguratorFromFile(path string) (Configurator, error) {
	envVars, err := getEnvVarsFromFile(path)
	if err != nil {
		return nil, err
	}

	return Config{
		env: envVars,
	}, nil
}

// NewConfiguratorFromOSEnvironment reads environment variables from the OS.
func NewConfiguratorFromOSEnvironment() (Config, error) {
	envVars, err := godotenv.Unmarshal(strings.Join(os.Environ(), "\n"))
	if err != nil {
		return Config{}, errors.Wrap(err, "could not parse env vars from the OS")
	}

	return Config{
		env: envVars,
	}, nil
}

func getEnvVarsFromFile(path string) (map[string]string, error) {
	envFilePath, err := filepath.Abs(path)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get absolute path for [%s]", envFilePath)
	}

	envVars, err := godotenv.Read(envFilePath)
	if err != nil {
		return nil, errors.Wrapf(err, "could not parse env vars")
	}

	return envVars, nil
}
