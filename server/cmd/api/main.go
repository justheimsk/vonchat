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

	log.Info("Opening database connection...")
	db, err := database.Open()
	if err != nil {
		log.Fatal("Fatal error: ", err)
    return
	}

	log.Info("Connected to the database.")
	defer db.Close()

	server := Server.New(db, log)
	server.CreateHTTPServer(config)
}
