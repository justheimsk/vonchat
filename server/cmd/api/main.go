package main

import (
	"github.com/justheimsk/vonchat/server/internal/database"
	Server "github.com/justheimsk/vonchat/server/internal/server"
	"github.com/justheimsk/vonchat/server/pkg/logger"
)

func main() {
	log := logger.NewLogger("CORE")

	log.Info("Opening database connection...")
	db, err := database.Open()
	if err != nil {
		log.Fatal("Fatal error: ", err)
	}

	log.Info("Connected to the database.")
	defer db.Close()

	server := Server.New(db, log)
	server.CreateHTTPServer()
}
