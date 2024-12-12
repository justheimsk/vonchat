package database

import (
	"strings"

	"github.com/justheimsk/vonchat/server/internal/domain/models"
	domain_repo "github.com/justheimsk/vonchat/server/internal/domain/repository"
	"github.com/justheimsk/vonchat/server/internal/infra/config"
	"github.com/justheimsk/vonchat/server/internal/registry"
)

type DatabaseDriver interface {
	Open() error
	Close() error
	GetName() string
	GetRepository() *domain_repo.RepositoryAggregate
	Init(*config.Config, models.Logger)
}

var driverRegistry = registry.NewRegistry[string, DatabaseDriver]()

func NewDatabaseDriver(driverName string, config *config.Config, logger models.Logger) DatabaseDriver {
	driver, found := driverRegistry.Get(strings.ToUpper(driverName))
	if !found {
		logger.Errorf("Failed to find driver %s", driverName)
		return nil
	}

	driver.Init(config, logger)
	return driver
}

func GetDriverRegistry() *registry.Registry[string, DatabaseDriver] {
	return driverRegistry
}
