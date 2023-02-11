package configurator

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

func NewConfiguratorFromFile(path string) (Configurator, error) {
	envVars, err := getEnvVarsFromFile(path)
	if err != nil {
		return nil, err
	}

	return config{
		env: envVars,
	}, nil
}

func NewConfiguratorFromOSEnvironment() (Configurator, error) {
	envVars, err := godotenv.Unmarshal(strings.Join(os.Environ(), "\n"))
	if err != nil {
		return nil, errors.Wrap(err, "could not parse env vars from the OS")
	}

	return config{
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
