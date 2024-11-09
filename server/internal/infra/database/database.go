package database

import "database/sql"

type DatabaseDriver interface {
  Open() error
  Close()
  GetDB() *sql.DB
  GetName() string
}
