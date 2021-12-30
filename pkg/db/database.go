package db

import (
	"errors"
	"strings"

	"github.com/chenjr0719/golang-boilerplate/pkg/config"
	"github.com/chenjr0719/golang-boilerplate/pkg/log"
	"gorm.io/gorm"
)

var DB *gorm.DB = ConnectDatabase()

func ConnectDatabase() *gorm.DB {
	var dbConnection *gorm.DB
	var err error

	switch {
	case strings.HasPrefix(config.Config.DatabaseURI, "sqlite://"):
		dbConnection, err = ConnectSqlite(config.Config.DatabaseURI)
	case strings.HasPrefix(config.Config.DatabaseURI, "postgresql://"):
		dbConnection, err = ConnectPostgres(config.Config.DatabaseURI)
	default:
		err = errors.New("unsupported db")
	}
	if err != nil {
		log.Fatal().Msg("Failed to connect database")
	}

	return dbConnection

}
