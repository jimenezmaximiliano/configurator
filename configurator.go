package configurator

import (
	"strconv"

	"github.com/pkg/errors"
)

// Configurator contains all the methods supported by this library.
// Create your own interface including only the methods that you need.
type Configurator interface {
	MustGetString(key string) (string, error)
	GetString(key string, defaultValue string) string
	MustGetBoolean(key string) (bool, error)
	GetBoolean(key string, defaultValue bool) bool
	MustGetInteger(key string) (int64, error)
	GetInteger(key string, defaultValue int64) int64
}

type config struct {
	env map[string]string
}

// MustGetBoolean return a boolean value or an error if the key couldn't be found, or its value is invalid.
func (conf config) MustGetBoolean(key string) (bool, error) {
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

// GetBoolean returns a boolean value based on the provided key or the default value if it couldn't be found,
// or its value is invalid.
func (conf config) GetBoolean(key string, defaultValue bool) bool {
	keyValue, err := conf.MustGetBoolean(key)
	if err != nil {
		return defaultValue
	}

	return keyValue
}

// MustGetString returns a string or an error if the key couldn't be found.
func (conf config) MustGetString(key string) (string, error) {
	keyValue, keyIsPresent := conf.env[key]
	if keyIsPresent {
		return keyValue, nil
	}

	return "", errors.Errorf("key [%s] not found in env", key)
}

// GetString returns a string value based on the provided key or the default value if it couldn't be found.
func (conf config) GetString(key string, defaultValue string) string {
	keyValue, err := conf.MustGetString(key)
	if err != nil {
		return defaultValue
	}

	return keyValue
}

// MustGetInteger returns an integer or an error if the key couldn't be found, or its value is invalid.
func (conf config) MustGetInteger(key string) (int64, error) {
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

// GetInteger returns an integer value based on the provided key or the default value if it couldn't be found,
// or its value is invalid.
func (conf config) GetInteger(key string, defaultValue int64) int64 {
	keyValue, err := conf.MustGetInteger(key)
	if err != nil {
		return defaultValue
	}

	return keyValue
}
