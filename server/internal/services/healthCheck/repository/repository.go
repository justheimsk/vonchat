package repositories

import (
	"database/sql"
	"time"
)

type healthCheckRepo struct {
	db *sql.DB
}

func NewHealthRepo(db *sql.DB) *healthCheckRepo {
	return &healthCheckRepo{
		db,
	}
}

func (self *healthCheckRepo) GetPing() (time.Duration, error) {
	start := time.Now()

	err := self.db.Ping()
	if err != nil {
		return 0, err
	}

	return time.Since(start), nil
}
