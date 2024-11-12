package repository

import (
	domain "github.com/justheimsk/vonchat/server/internal/domain/repository"
	"github.com/justheimsk/vonchat/server/internal/infra/database"
	"github.com/justheimsk/vonchat/server/internal/infra/persistence/repository/pgsql"
	"github.com/justheimsk/vonchat/server/internal/infra/persistence/repository/sqlite"
)

func NewAuthRepository(driver database.DatabaseDriver) (repo domain.AuthRepository) {
  switch driver.GetName() {
  case "POSTGRES":
      repo = pgsql.NewAuthRepository(driver.GetDB())
  case "SQLITE":
      repo = sqlite.NewAuthRepository(driver.GetDB())
  }

  return
}

func NewHealthRepository(driver database.DatabaseDriver) (repo domain.HealthRepository) {
  switch driver.GetName() {
  case "POSTGRES":
      repo = pgsql.NewHealthRepository(driver.GetDB())
  case "SQLITE":
      repo = sqlite.NewHealthRepository(driver.GetDB())
  }

  return
}

func NewUserRepository(driver database.DatabaseDriver) (repo domain.UserRepository) {
  switch driver.GetName() {
  case "SQLITE":
    repo = sqlite.NewUserRepository(driver.GetDB())
  case "POSTGRES":
    repo = pgsql.NewUserRepository(driver.GetDB())
  }

  return
}
