package database

import (
	"database/sql"
	"fmt"

	"github.com/justheimsk/vonchat/server/internal/infra/config"
	"github.com/justheimsk/vonchat/server/internal/infra/logger"
	"github.com/justheimsk/vonchat/server/scripts"
	_ "github.com/lib/pq"
)

type PostgresDatabaseDriver struct {
  Host     string
  Port     string
  DB       string
  User     string
  Password string
  db *sql.DB
  logger *logger.Logger
}

func NewPostgresDatabaseDriver(config *config.Config) *PostgresDatabaseDriver {
  return &PostgresDatabaseDriver{
    Host: config.PostgresHost,
    Port: config.PostgresPort,
    DB: config.PostgresDB,
    User: config.PostgresUser,
    Password: config.PostgresPassword,
    logger: logger.NewLogger("DATABASE", config),
  }
}

func (self *PostgresDatabaseDriver) Open() (err error) {
	str := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", self.Host, self.Port, self.DB, self.User, self.Password)
  db, err := sql.Open("postgres", str)
	if err != nil {
		err = fmt.Errorf("Failed to open connection: %w", err)
		return
	}

	err = db.Ping()
	if err != nil {
		err = fmt.Errorf("Failed to connect to the database: %w", err)
    return
	}

  _, err = db.Exec(scripts.GetPGInitScript())
  if err != nil {
    self.logger.Warn("Failed to exec init script: ", err) 
  }

  self.db = db
	return
}

func (self *PostgresDatabaseDriver) GetDB() (*sql.DB) {
  return self.db
}

func (self *PostgresDatabaseDriver) GetName() string {
  return "POSTGRES"
}

func (self *PostgresDatabaseDriver) Close() {
  self.db.Close()
}