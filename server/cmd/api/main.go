package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/justheimsk/vonchat/server/internal/infra/config"
	"github.com/justheimsk/vonchat/server/internal/infra/database"
	"github.com/justheimsk/vonchat/server/internal/infra/http"
)

func setupLogger() *log.Logger {
	logger := log.New(os.Stdout)
	logger.SetReportCaller(true)
	logger.SetReportTimestamp(true)

	return logger
}

func main() {
	logger := setupLogger()

	config, err := config.LoadConfig(logger)
	if err != nil {
		log.Fatal("Failed to load config", "err", err)
		return
	}

	if config.Debug {
		logger.SetLevel(log.DebugLevel)
		logger.Debug("Debug mode enabled.")
	}

	driver := database.NewDatabaseDriver(config.DatabaseDriver, config, logger)
	defer driver.Close()

	logger.Infof("Using %s database driver.", driver.GetName())
	logger.Info("Opening database connection...")

	err = driver.Open()
	if err != nil {
		logger.Fatal("Failed to open database connection", "err", err)
		return
	}

	logger.Info("Connected to the database.")
	server := http.NewServer(driver, logger)
	server.Serve(config)
}
