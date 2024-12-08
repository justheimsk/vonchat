package database

import (
	"database/sql"

	"github.com/justheimsk/vonchat/server/internal/domain/models"
	"github.com/justheimsk/vonchat/server/internal/infra/config"
)

type DatabaseDriver interface {
	Open() error
	Close() error
	GetDB() *sql.DB
	GetName() string
}

func NewDatabaseDriver(driverName string, config *config.Config, logger models.Logger) DatabaseDriver {
	switch driverName {
	case "POSTGRES":
		return NewPostgresDatabaseDriver(config, logger)
	case "SQLITE":
		return NewSQLiteDatabaseDriver(config, logger)
	}

	return nil
}
