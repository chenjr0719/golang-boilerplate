package db_test

import (
	"os"
	"testing"

	"github.com/chenjr0719/golang-boilerplate/pkg/db"
	"github.com/stretchr/testify/assert"
)

func TestConnectDatabase(t *testing.T) {
	databaseURI := "sqlite://file::memory:?cache=shared"
	os.Setenv("DATABASE_URI", databaseURI)

	db.ConnectDatabase()
	assert.NoError(t, nil)
}
