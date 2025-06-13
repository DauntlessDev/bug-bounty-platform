package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Create a temporary .env file for testing
	tempEnvFile := ".env.test"
	err := os.WriteFile(tempEnvFile, []byte("DATABASE_URL=test_db_url\nSERVER_PORT=8081"), 0644)
	assert.NoError(t, err)
	defer func() {
		err := os.Remove(tempEnvFile)
		assert.NoError(t, err)
	}()

	// Unset environment variables that might interfere with the test
	_ = os.Unsetenv("DATABASE_URL")
	_ = os.Unsetenv("SERVER_PORT")

	cfg, err := LoadConfig(tempEnvFile)

	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "test_db_url", cfg.DatabaseURL)
	assert.Equal(t, "8081", cfg.ServerPort)

	// Test with missing .env file (should not return an error)
	err = os.Remove(tempEnvFile)
	assert.NoError(t, err)

	cfg, err = LoadConfig()
	assert.NoError(t, err)
	assert.NotNil(t, cfg)

	// Clean up environment variables set by the test
	_ = os.Unsetenv("DATABASE_URL")
	_ = os.Unsetenv("SERVER_PORT")
}

func TestLoadConfig_EnvVarsOverride(t *testing.T) {
	// Set environment variables directly
	err := os.Setenv("DATABASE_URL", "env_db_url")
	assert.NoError(t, err)

	err = os.Setenv("SERVER_PORT", "9000")
	assert.NoError(t, err)

	// Create a temporary .env file (should be overridden by env vars)
	tempEnvFile := ".env.test"
	err = os.WriteFile(tempEnvFile, []byte("DATABASE_URL=file_db_url\nSERVER_PORT=8082"), 0644)
	assert.NoError(t, err)
	defer func() {
		err := os.Remove(tempEnvFile)
		assert.NoError(t, err)
	}()

	cfg, err := LoadConfig(tempEnvFile)

	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "env_db_url", cfg.DatabaseURL) // Should be overridden
	assert.Equal(t, "9000", cfg.ServerPort)        // Should be overridden
}
