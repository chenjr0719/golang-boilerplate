package db

import (
	"github.com/chenjr0719/golang-boilerplate/pkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres(databaseUri string) (*gorm.DB, error) {
	log.Info().Msg("Connect to PostgreSQL")

	db, err := gorm.Open(postgres.Open(databaseUri), &gorm.Config{})

	return db, err
}
