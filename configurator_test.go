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

	value, err := config.MustGetInteger(key)
	assert.Equal(test, int64(1), value)
	assert.NoError(test, err)
}

func TestGettingARequiredIntFailsWhenNotFound(test *testing.T) {
	test.Parallel()

	config, err := configurator.NewConfiguratorFromOSEnvironment()
	require.NoError(test, err)

	value, err := config.MustGetInteger("REQUIRED_INT_MISSING")
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

	value, err := config.MustGetInteger(key)
	assert.Equal(test, int64(0), value)
	assert.Error(test, err)
}

func TestGettingAnInt(test *testing.T) {
	test.Parallel()

	const key = "INT"
	err := os.Setenv(key, "1")
	require.NoError(test, err)
	defer func() {
		err := os.Unsetenv(key)
		assert.NoError(test, err)
	}()

	config, err := configurator.NewConfiguratorFromOSEnvironment()
	require.NoError(test, err)

	assert.Equal(test, int64(1), config.GetInteger(key, 0))
}

func TestGettingAnIntWithDefault(test *testing.T) {
	test.Parallel()

	const key = "INT"

	config, err := configurator.NewConfiguratorFromOSEnvironment()
	require.NoError(test, err)

	assert.Equal(test, int64(7), config.GetInteger(key, 7))
}
