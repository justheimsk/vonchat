package database

import (
	"database/sql"
	"fmt"

	"github.com/justheimsk/vonchat/server/internal/domain/models"
	domain_repo "github.com/justheimsk/vonchat/server/internal/domain/repository"
	"github.com/justheimsk/vonchat/server/internal/infra/config"
	"github.com/justheimsk/vonchat/server/internal/infra/persistence/repository/sqlite"
	"github.com/justheimsk/vonchat/server/scripts"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDatabaseDriver struct {
	config *config.Config
	db     *sql.DB
	logger models.Logger
}

func NewSQLiteDatabaseDriver(config *config.Config, logger models.Logger) *SQLiteDatabaseDriver {
	return &SQLiteDatabaseDriver{
		config: config,
		logger: logger,
	}
}

func init() {
	GetDriverRegistry().Register("SQLITE", NewSQLiteDatabaseDriver(nil, nil))
}

func (self *SQLiteDatabaseDriver) Open() error {
	if self.config == nil {
		return fmt.Errorf("Config not set")
	}

	db, err := sql.Open("sqlite3", self.config.Sqlite.Path)
	if err != nil {
		return fmt.Errorf("Failed to open connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %w", err)
	}

	_, err = db.Exec(scripts.GetSQLiteInitScript())
	if err != nil {
		self.logger.Warn("Failed to exec init script", "err", err)
	}

	self.db = db
	return nil
}

func (self *SQLiteDatabaseDriver) Init(config *config.Config, logger models.Logger) {
	self.config = config
	self.logger = logger
}

func (self *SQLiteDatabaseDriver) GetDB() *sql.DB {
	return self.db
}

func (self *SQLiteDatabaseDriver) Close() error {
	return self.db.Close()
}

func (self *SQLiteDatabaseDriver) GetName() string {
	return "SQLITE"
}

func (self *SQLiteDatabaseDriver) GetRepository() *domain_repo.RepositoryAggregate {
	return &domain_repo.RepositoryAggregate{
		Health: sqlite.NewHealthRepository(self.db),
		User:   sqlite.NewUserRepository(self.db),
		Auth:   sqlite.NewAuthRepository(self.db),
	}
}
