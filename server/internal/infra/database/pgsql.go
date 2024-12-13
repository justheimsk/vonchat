package database

import (
	"database/sql"
	"fmt"

	"github.com/justheimsk/vonchat/server/internal/domain/models"
	domain_repo "github.com/justheimsk/vonchat/server/internal/domain/repository"
	"github.com/justheimsk/vonchat/server/internal/infra/config"
	"github.com/justheimsk/vonchat/server/internal/infra/persistence/repository/pgsql"
	"github.com/justheimsk/vonchat/server/scripts"
	_ "github.com/lib/pq"
)

type PostgresDatabaseDriver struct {
	config *config.Config
	db     *sql.DB
	logger models.Logger
}

func NewPostgresDatabaseDriver(config *config.Config, logger models.Logger) *PostgresDatabaseDriver {
	return &PostgresDatabaseDriver{
		config: config,
		logger: logger,
	}
}

func init() {
	GetDriverRegistry().Register("POSTGRES", NewPostgresDatabaseDriver(nil, nil))
}

func (self *PostgresDatabaseDriver) Open() (err error) {
	if self.config == nil {
		return models.NewCustomError("config_not_set", "Config not set")
	}

	str := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", self.config.Postgres.Host, self.config.Postgres.Port, self.config.Postgres.DB, self.config.Postgres.User, self.config.Postgres.Password)
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
		self.logger.Warn("Failed to exec init script", "err", err)
	}

	self.db = db
	return
}

func (self *PostgresDatabaseDriver) Init(config *config.Config, logger models.Logger) {
	self.config = config
	self.logger = logger
}

func (self *PostgresDatabaseDriver) GetName() string {
	return "POSTGRES"
}

func (self *PostgresDatabaseDriver) Close() error {
	return self.db.Close()
}

func (self *PostgresDatabaseDriver) GetRepository() *domain_repo.RepositoryAggregate {
	return &domain_repo.RepositoryAggregate{
		Health: pgsql.NewHealthRepository(self.db),
		User:   pgsql.NewUserRepository(self.db),
		Auth:   pgsql.NewAuthRepository(self.db),
	}
}
