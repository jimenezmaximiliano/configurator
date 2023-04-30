package configurator

import (
	"strconv"

	"github.com/pkg/errors"
)

// Configurator contains all the methods supported by this library.
// Create your own interface including only the methods that you need.
type Configurator interface {
	GetString(key string) (string, error)
	GetBoolean(key string) (bool, error)
	GetInteger(key string) (int64, error)
}

var _ Configurator = Config{}

type Config struct {
	env map[string]string
}

// GetBoolean return a boolean value or an error if the key couldn't be found, or its value is invalid.
func (conf Config) GetBoolean(key string) (bool, error) {
	keyValue, keyIsPresent := conf.env[key]
	if keyIsPresent {
		if keyValue == "true" {
			return true, nil
		}

		if keyValue == "false" {
			return false, nil
		}

		return false, errors.Errorf("key [%s] has an invalid boolean value [%s]", key, keyValue)
	}

	return false, errors.Errorf("key [%s] not found in env", key)
}

// GetString returns a string or an error if the key couldn't be found.
func (conf Config) GetString(key string) (string, error) {
	keyValue, keyIsPresent := conf.env[key]
	if keyIsPresent {
		return keyValue, nil
	}

	return "", errors.Errorf("key [%s] not found in env", key)
}

// GetInteger returns an integer or an error if the key couldn't be found, or its value is invalid.
func (conf Config) GetInteger(key string) (int64, error) {
	keyValue, keyIsPresent := conf.env[key]
	if !keyIsPresent {
		return 0, errors.Errorf("key [%s] not found in env", key)
	}

	integer, err := strconv.ParseInt(keyValue, 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "key [%s] has an invalid integer value [%s]", key, keyValue)
	}

	return integer, nil
}
