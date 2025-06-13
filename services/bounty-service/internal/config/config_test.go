package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTempEnvFile(t *testing.T, content string) string {
	t.Helper()
	tempEnvFile := ".env.test"
	err := os.WriteFile(tempEnvFile, []byte(content), 0644)
	assert.NoError(t, err, "Failed to write temporary .env file")
	return tempEnvFile
}

func TestLoadConfig(t *testing.T) {
	// Create and defer cleanup of temporary .env file
	tempEnvFile := createTempEnvFile(t, "DATABASE_URL=test_db_url\nSERVER_PORT=8081")
	defer func() {
		_ = os.Remove(tempEnvFile) // Safe cleanup even if already deleted
	}()

	// Ensure related environment variables are unset
	_ = os.Unsetenv("DATABASE_URL")
	_ = os.Unsetenv("SERVER_PORT")

	cfg, err := LoadConfig(tempEnvFile)
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "test_db_url", cfg.DatabaseURL)
	assert.Equal(t, "8081", cfg.ServerPort)

	// Remove .env file and test fallback
	_ = os.Remove(tempEnvFile)
	cfg, err = LoadConfig()
	assert.NoError(t, err)
	assert.NotNil(t, cfg)

	// Clean up env vars set by test
	_ = os.Unsetenv("DATABASE_URL")
	_ = os.Unsetenv("SERVER_PORT")
}

func TestLoadConfig_EnvVarsOverride(t *testing.T) {
	_ = os.Setenv("DATABASE_URL", "env_db_url")
	_ = os.Setenv("SERVER_PORT", "9000")

	tempEnvFile := createTempEnvFile(t, "DATABASE_URL=file_db_url\nSERVER_PORT=8082")
	defer func() {
		_ = os.Remove(tempEnvFile)
	}()

	cfg, err := LoadConfig(tempEnvFile)
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "env_db_url", cfg.DatabaseURL) // Should use env var
	assert.Equal(t, "9000", cfg.ServerPort)        // Should use env var
}
