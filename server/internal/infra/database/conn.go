package database

import (
	"database/sql"
	"fmt"

	"github.com/justheimsk/vonchat/server/scripts"
	_ "github.com/lib/pq"
)

func Open() (db *sql.DB, err error) {
	host := "localhost"
	port := 5432
	dbname := "main"
	user := "root"
	password := "admin"

	str := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", host, port, dbname, user, password)
	db, err = sql.Open("postgres", str)
	if err != nil {
		err = fmt.Errorf("Failed to open connection: %w", err)
		return
	}

	err = db.Ping()
	if err != nil {
		err = fmt.Errorf("Failed to connect to the database: %w", err)
    return
	}

  _, err = db.Exec(scripts.GetPQInitScript())
  if err != nil {
    err = fmt.Errorf("Failed to run init script: %w", err)
    return
  }

	return
}
