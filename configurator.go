package configurator

import (
	"github.com/pkg/errors"
)

type Configurator interface {
	MustGetString(key string) (string, error)
	GetString(key string, defaultValue string) string
	MustGetBoolean(key string) (bool, error)
	GetBoolean(key string, defaultValue bool) bool
}

type config struct {
	env map[string]string
}

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

func (conf config) MustGetString(key string) (string, error) {
	keyValue, keyIsPresent := conf.env[key]
	if keyIsPresent {
		return keyValue, nil
	}

	return "", errors.Errorf("key [%s] not found in env", key)
}

func (conf config) GetString(key string, defaultValue string) string {
	keyValue, err := conf.MustGetString(key)
	if err != nil {
		return defaultValue
	}

	return keyValue
}

func (conf config) GetBoolean(key string, defaultValue bool) bool {
	keyValue, err := conf.MustGetBoolean(key)
	if err != nil {
		return defaultValue
	}

	return keyValue
}
