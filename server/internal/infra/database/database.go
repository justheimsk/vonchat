package database

import (
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	domain_repo "github.com/justheimsk/vonchat/server/internal/domain/repository"
	"github.com/justheimsk/vonchat/server/internal/infra/config"
)

type DatabaseDriver interface {
	Open() error
	Close() error
	GetName() string
	GetRepository() *domain_repo.RepositoryAggregate
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
