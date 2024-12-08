package sqlite

import (
	"database/sql"
	"time"
)

type healthRepository struct {
	db *sql.DB
}

func NewHealthRepository(db *sql.DB) *healthRepository {
	return &healthRepository{
		db,
	}
}

func (self *healthRepository) GetPing() (time.Duration, error) {
	start := time.Now()

	err := self.db.Ping()
	if err != nil {
		return 0, err
	}

	return time.Since(start), nil
}
