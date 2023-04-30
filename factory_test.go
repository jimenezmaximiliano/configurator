package configurator_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jimenezmaximiliano/configurator"
)

func TestGettingAConfiguratorFromAFile(test *testing.T) {
	test.Parallel()

	const content = `
	FOO="var"
	BOOL="true"
	`
	const path = "./.env"

	err := os.WriteFile(path, []byte(content), 0777)
	require.NoError(test, err)
	defer func() {
		err := os.Remove(path)
		assert.NoError(test, err)
	}()

	config, err := configurator.NewConfiguratorFromFile(path)
	require.NoError(test, err)
	require.NotNil(test, config)

	stringConfig, err := config.GetString("FOO")

	assert.NoError(test, err)
	assert.Equal(test, "var", stringConfig)

	booleanConfig, err := config.GetBoolean("BOOL")

	assert.NoError(test, err)
	assert.Equal(test, true, booleanConfig)
}

func TestGettingAConfiguratorFromAnInvalidFilePath(test *testing.T) {
	test.Parallel()

	const path = ""
	config, err := configurator.NewConfiguratorFromFile(path)

	assert.Error(test, err)
	assert.Nil(test, config)
}

func TestGettingAConfiguratorFromOSEnvVars(test *testing.T) {
	test.Parallel()

	err := os.Setenv("FOO", "var")
	require.NoError(test, err)

	config, err := configurator.NewConfiguratorFromOSEnvironment()
	require.NoError(test, err)

	stringConfig, err := config.GetString("FOO")

	assert.NoError(test, err)
	assert.Equal(test, "var", stringConfig)
}
