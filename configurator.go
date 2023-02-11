package configurator

import (
	"github.com/pkg/errors"
)

// Configurator contains all the methods supported by this library.
// Create your own interface including only the methods that you need.
type Configurator interface {
	MustGetString(key string) (string, error)
	GetString(key string, defaultValue string) string
	MustGetBoolean(key string) (bool, error)
	GetBoolean(key string, defaultValue bool) bool
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

// GetBoolean returns a boolean value based on the provided key or the default value if it couldn't be found,
// or its value is invalid.
func (conf config) GetBoolean(key string, defaultValue bool) bool {
	keyValue, err := conf.MustGetBoolean(key)
	if err != nil {
		return defaultValue
	}

	return keyValue
}
