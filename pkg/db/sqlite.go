package db

import (
	"strings"

	"github.com/chenjr0719/golang-boilerplate/pkg/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectSqlite(databaseUri string) (*gorm.DB, error) {
	log.Info().Msg("Connect to SQLite")

	databaseUri = strings.Replace(databaseUri, "sqlite://", "", 1)
	db, err := gorm.Open(sqlite.Open(databaseUri), &gorm.Config{})

	return db, err
}
