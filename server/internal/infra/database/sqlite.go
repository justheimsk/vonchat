package database

import (
	"database/sql"
	"fmt"

	"github.com/justheimsk/vonchat/server/internal/domain/models"
	"github.com/justheimsk/vonchat/server/internal/infra/config"
	"github.com/justheimsk/vonchat/server/scripts"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDatabaseDriver struct {
  Path   string
  db     *sql.DB
  logger models.Logger
}

func NewSQLiteDatabaseDriver(config *config.Config, logger models.Logger) *SQLiteDatabaseDriver {
  return &SQLiteDatabaseDriver{
    Path: config.SQLitePath,
    logger: logger.New("DATABASE"),
  }
}

func (self *SQLiteDatabaseDriver) Open() error {
  db, err := sql.Open("sqlite3", self.Path)
  if err != nil {
    return fmt.Errorf("Failed to open connection: %w", err)
  }

  err = db.Ping()
  if err != nil {
    return fmt.Errorf("Failed to connect to database: %w", err)
  }

  _, err = db.Exec(scripts.GetSQLiteInitScript())
  if err != nil {
    self.logger.Warnf("Failed to exec init script: %w", err) 
  }

  self.db = db
  return nil
}

func (self *SQLiteDatabaseDriver) GetDB() *sql.DB {
  return self.db
}

func (self *SQLiteDatabaseDriver) Close() {
  self.db.Close()
}


func (self *SQLiteDatabaseDriver) GetName() string {
  return "SQLITE"
}
