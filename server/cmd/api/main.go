package main

import (
	"github.com/justheimsk/vonchat/server/internal/database"
	Server "github.com/justheimsk/vonchat/server/internal/server"
	logger "github.com/justheimsk/vonchat/server/pkg/logger"
)

func main() {
	log := logger.GetLogger()

	log.Println("Opening database connection...")
	db, err := database.Open()
	if err != nil {
		log.Fatalf("Fatal error: %s", err)
	}

	log.Println("Connected to database.")
	server := Server.New(db, log)
	server.Init()
}
