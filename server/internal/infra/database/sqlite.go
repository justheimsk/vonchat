package database

import (
	"database/sql"
	"fmt"

	"github.com/justheimsk/vonchat/server/internal/infra/config"
	"github.com/justheimsk/vonchat/server/internal/infra/logger"
	"github.com/justheimsk/vonchat/server/scripts"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDatabaseDriver struct {
  Path   string
  db     *sql.DB
  logger *logger.Logger
}

func NewSQLiteDatabaseDriver(config *config.Config) *SQLiteDatabaseDriver {
  return &SQLiteDatabaseDriver{
    Path: config.SQLitePath,
    logger: logger.NewLogger("DATABASE", config),
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
    self.logger.Warn("Failed to exec init script: ", err) 
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
