package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Create a temporary .env file for testing
	tempEnvFile := ".env.test"
	os.WriteFile(tempEnvFile, []byte("DATABASE_URL=test_db_url\nSERVER_PORT=8081"), 0644)
	defer os.Remove(tempEnvFile) // Clean up the temporary file

	// Unset environment variables that might interfere with the test
	os.Unsetenv("DATABASE_URL")
	os.Unsetenv("SERVER_PORT")

	cfg, err := LoadConfig(tempEnvFile)

	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "test_db_url", cfg.DatabaseURL)
	assert.Equal(t, "8081", cfg.ServerPort)

	// Test with missing .env file (should not return an error)
	os.Remove(tempEnvFile)
	cfg, err = LoadConfig()
	assert.NoError(t, err)
	assert.NotNil(t, cfg)

	// Clean up environment variables set by the test
	os.Unsetenv("DATABASE_URL")
	os.Unsetenv("SERVER_PORT")
}

func TestLoadConfig_EnvVarsOverride(t *testing.T) {
	// Set environment variables directly
	os.Setenv("DATABASE_URL", "env_db_url")
	os.Setenv("SERVER_PORT", "9000")
	// Create a temporary .env file (should be overridden by direct env vars)
	tempEnvFile := ".env.test"
	os.WriteFile(tempEnvFile, []byte("DATABASE_URL=file_db_url\nSERVER_PORT=8082"), 0644)
	defer os.Remove(tempEnvFile)

	cfg, err := LoadConfig(tempEnvFile)

	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "env_db_url", cfg.DatabaseURL) // Should be overridden
	assert.Equal(t, "9000", cfg.ServerPort)        // Should be overridden
}
