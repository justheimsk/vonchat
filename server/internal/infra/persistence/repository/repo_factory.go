package repository

import (
	domain "github.com/justheimsk/vonchat/server/internal/domain/repository"
	"github.com/justheimsk/vonchat/server/internal/infra/database"
	"github.com/justheimsk/vonchat/server/internal/infra/persistence/repository/pgsql"
)

func NewAuthRepository(driver database.DatabaseDriver) (repo domain.AuthRepository) {
  switch driver.GetName() {
  case "POSTGRES":
      repo = pgsql.NewAuthRepository(driver.GetDB())
  }

  return
}

func NewHealthRepository(driver database.DatabaseDriver) (repo domain.HealthRepository) {
  switch driver.GetName() {
  case "POSTGRES":
      repo = pgsql.NewHealthRepository(driver.GetDB())
  }

  return
}
