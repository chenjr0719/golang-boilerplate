package db

import (
	"errors"
	"strings"

	"github.com/chenjr0719/golang-boilerplate/pkg/config"
	"github.com/chenjr0719/golang-boilerplate/pkg/log"
	"gorm.io/gorm"
)

func ConnectDatabase(conf *config.Config) (*gorm.DB, error) {
	var dbConnection *gorm.DB
	var err error

	databaseURI := conf.DatabaseURI
	switch {
	case strings.HasPrefix(databaseURI, "sqlite://"):
		dbConnection, err = ConnectSqlite(databaseURI)
	case strings.HasPrefix(databaseURI, "postgresql://"):
		dbConnection, err = ConnectPostgres(databaseURI)
	default:
		err = errors.New("unsupported db")
	}
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect database")
		return nil, err
	}

	return dbConnection, nil

}
