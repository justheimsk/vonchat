package main

import (
	"github.com/justheimsk/vonchat/server/internal/infra/config"
	"github.com/justheimsk/vonchat/server/internal/infra/database"
	"github.com/justheimsk/vonchat/server/internal/infra/logger"
	Server "github.com/justheimsk/vonchat/server/internal/server"
)

func main() {
	log := logger.NewLogger("CORE")
  config, err := config.LoadConfig(logger.NewLogger("CONFIG"))
  if err != nil {
    log.Fatal("Failed to load config: ", err)
    return
  }

  var driver database.DatabaseDriver

  if config.DatabaseDriver == "POSTGRES" {
    driver = database.NewPostgresDatabaseDriver(config)
  } else if config.DatabaseDriver == "SQLITE" {
    driver = database.NewSQLiteDatabaseDriver(config)
  }

  log.Info("Using ", driver.GetName(), " database driver.")
	log.Info("Opening database connection...")
	err = driver.Open()
	if err != nil {
		log.Fatal("Fatal error: ", err)
    return
	}

	log.Info("Connected to the database.")
	defer driver.Close()

	server := Server.New(driver, log)
	server.CreateHTTPServer(config)
}
