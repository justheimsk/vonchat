package main

import (
  "github.com/justheimsk/vonchat/server/internal/infra/config"
  "github.com/justheimsk/vonchat/server/internal/infra/database"
  http "github.com/justheimsk/vonchat/server/internal/infra/http"
  "github.com/justheimsk/vonchat/server/pkg/logger"
)

func main() {
  log := logger.NewLogger("CORE", nil, nil)
  config, err := config.LoadConfig(log.New("CONFIG"))
  log = logger.NewLogger("CORE", config, nil)

  if config.Debug {
    log.Debugf("Debug mode enabled.")
  }

  if err != nil {
    log.Fatalf("Failed to load config: %w", err)
    return
  }

  driver := database.NewDatabaseDriver(config.DatabaseDriver, config, log)
  log.Infof("Using %s database driver.", driver.GetName())
  log.Infof("Opening database connection...")
  err = driver.Open()

  if err != nil {
    log.Fatalf("Fatal error: %w", err)
    return
  }

  log.Infof("Connected to the database.")
  defer driver.Close()

  server := http.NewServer(driver, log)
  server.Serve(config)
}
