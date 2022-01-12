package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chenjr0719/golang-boilerplate/pkg/config"
)

func TestLoadConfig(t *testing.T) {
	mode := "release"
	logLevel := "error"
	databaseURI := "sqlite://file::memory:?cache=shared"

	os.Setenv("MODE", mode)
	os.Setenv("LOG_LEVEL", logLevel)
	os.Setenv("DATABASE_URI", databaseURI)

	config := config.LoadConfig()

	assert.Equal(t, config.Mode, mode)
	assert.Equal(t, config.LogLevel, logLevel)
	assert.Equal(t, config.DatabaseURI, databaseURI)

}
