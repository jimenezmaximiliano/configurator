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

	assert.Equal(test, true, config.GetBoolean(key, false))
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

	assert.Equal(test, false, config.GetBoolean(key, true))
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

	assert.Equal(test, true, config.GetBoolean(key, true))
}

func TestGettingADefaultBoolean(test *testing.T) {
	test.Parallel()

	config, err := configurator.NewConfiguratorFromOSEnvironment()
	require.NoError(test, err)

	assert.Equal(test, true, config.GetBoolean("FOO", true))
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

	assert.Equal(test, "var", config.GetString(key, ""))
}

func TestGettingADefaultString(test *testing.T) {
	test.Parallel()

	config, err := configurator.NewConfiguratorFromOSEnvironment()
	require.NoError(test, err)

	assert.Equal(test, "var", config.GetString("DEFAULT_STRING", "var"))
}

func TestGettingARequiredStringFails(test *testing.T) {
	test.Parallel()

	config, err := configurator.NewConfiguratorFromOSEnvironment()
	require.NoError(test, err)

	value, err := config.MustGetString("REQUIRED_STRING")
	assert.Equal(test, "", value)
	assert.Error(test, err)
}

func TestGettingARequiredBooleanFails(test *testing.T) {
	test.Parallel()

	config, err := configurator.NewConfiguratorFromOSEnvironment()
	require.NoError(test, err)

	value, err := config.MustGetBoolean("REQUIRED_BOOLEAN")
	assert.Equal(test, false, value)
	assert.Error(test, err)
}
