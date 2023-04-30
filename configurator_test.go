package configurator_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jimenezmaximiliano/configurator"
)

func TestGettingATrueBoolean(test *testing.T) {
	test.Parallel()

	const key = "TRUE_BOOLEAN"
	err := os.Setenv(key, "true")
	require.NoError(test, err)
	defer func() {
		err := os.Unsetenv(key)
		assert.NoError(test, err)
	}()

	config, err := configurator.NewConfiguratorFromOSEnvironment()
	require.NoError(test, err)

	booleanConfig, err := config.GetBoolean(key)

	assert.NoError(test, err)
	assert.Equal(test, true, booleanConfig)
}

func TestGettingAFalseBoolean(test *testing.T) {
	test.Parallel()

	const key = "FALSE_BOOLEAN"
	err := os.Setenv(key, "false")
	require.NoError(test, err)
	defer func() {
		err := os.Unsetenv(key)
		assert.NoError(test, err)
	}()

	config, err := configurator.NewConfiguratorFromOSEnvironment()
	require.NoError(test, err)

	booleanConfig, err := config.GetBoolean(key)

	assert.NoError(test, err)
	assert.Equal(test, false, booleanConfig)
}

func TestGettingAnInvalidBoolean(test *testing.T) {
	test.Parallel()

	const key = "INVALID_BOOLEAN"
	err := os.Setenv(key, "obladi")
	require.NoError(test, err)
	defer func() {
		err := os.Unsetenv(key)
		assert.NoError(test, err)
	}()

	config, err := configurator.NewConfiguratorFromOSEnvironment()
	require.NoError(test, err)

	booleanConfig, err := config.GetBoolean(key)

	assert.Error(test, err)
	assert.Equal(test, false, booleanConfig)
}

func TestGettingAString(test *testing.T) {
	test.Parallel()

	const key = "STRING"
	err := os.Setenv(key, "var")
	require.NoError(test, err)
	defer func() {
		err := os.Unsetenv(key)
		assert.NoError(test, err)
	}()

	config, err := configurator.NewConfiguratorFromOSEnvironment()
	require.NoError(test, err)

	stringConfig, err := config.GetString(key)

	assert.NoError(test, err)
	assert.Equal(test, "var", stringConfig)
}

func TestGettingARequiredStringFails(test *testing.T) {
	test.Parallel()

	config, err := configurator.NewConfiguratorFromOSEnvironment()
	require.NoError(test, err)

	value, err := config.GetString("REQUIRED_STRING")
	assert.Equal(test, "", value)
	assert.Error(test, err)
}

func TestGettingARequiredBooleanFails(test *testing.T) {
	test.Parallel()

	config, err := configurator.NewConfiguratorFromOSEnvironment()
	require.NoError(test, err)

	value, err := config.GetBoolean("REQUIRED_BOOLEAN")
	assert.Equal(test, false, value)
	assert.Error(test, err)
}

func TestGettingARequiredInt(test *testing.T) {
	test.Parallel()

	const key = "REQUIRED_INT"
	err := os.Setenv(key, "1")
	require.NoError(test, err)
	defer func() {
		err := os.Unsetenv(key)
		assert.NoError(test, err)
	}()

	config, err := configurator.NewConfiguratorFromOSEnvironment()
	require.NoError(test, err)

	value, err := config.GetInteger(key)
	assert.Equal(test, int64(1), value)
	assert.NoError(test, err)
}

func TestGettingARequiredIntFailsWhenNotFound(test *testing.T) {
	test.Parallel()

	config, err := configurator.NewConfiguratorFromOSEnvironment()
	require.NoError(test, err)

	value, err := config.GetInteger("REQUIRED_INT_MISSING")
	assert.Equal(test, int64(0), value)
	assert.Error(test, err)
}

func TestGettingARequiredIntFailsWhenInvalid(test *testing.T) {
	test.Parallel()

	const key = "REQUIRED_INT_INVALID"
	err := os.Setenv(key, "foo")
	require.NoError(test, err)
	defer func() {
		err := os.Unsetenv(key)
		assert.NoError(test, err)
	}()

	config, err := configurator.NewConfiguratorFromOSEnvironment()
	require.NoError(test, err)

	value, err := config.GetInteger(key)
	assert.Equal(test, int64(0), value)
	assert.Error(test, err)
}
