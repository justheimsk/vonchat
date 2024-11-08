package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/justheimsk/vonchat/server/internal/infra/logger"
)

type Config struct {
  DatabaseDriver   string
  SQLitePath       string
  PostgresHost     string
  PostgresPort     string
  PostgresDB       string
  PostgresUser     string
  PostgresPassword string
  Port             string
}

func LoadConfig(log *logger.Logger) (*Config, error) {
  config := &Config{}
  
  DBDriver := strings.ToUpper(os.Getenv("DATABASE_DRIVER"))
  if DBDriver == "" {
    DBDriver = "SQLITE"
  }

  if DBDriver != "SQLITE" && DBDriver != "POSTGRES" {
    return nil, fmt.Errorf("Invalid database driver.")
  }

  if DBDriver == "SQLITE" {
    path := os.Getenv("SQLITE_PATH")
    if path == "" {
      return nil, fmt.Errorf("Using SQLITE database driver but variable SQLITE_PATH is missing.")
    }

    config.SQLitePath = path;
  } else if DBDriver == "POSTGRES" {
    prefix := "POSTGRES_"
    keys := []string{"HOST", "PORT", "DB", "USER", "PASSWORD"}
    errorDet := false

    for _, k := range(keys) {
      value := os.Getenv(prefix + k)
      if value == "" {
        log.Error("Missing variable: ", prefix + k)
        errorDet = true
        continue
      }

      if k == "HOST" {
        config.PostgresHost = value
      } else if k == "PORT" {
        config.PostgresPort = value
      } else if k == "DB" {
        config.PostgresDB = value
      } else if k == "USER" {
        config.PostgresUser = value
      } else if k == "PASSWORD" {
        config.PostgresPassword = value
      }
    }

    if errorDet != false {
      return nil, fmt.Errorf("Using POSTGRES database driver but missing above variables.")
    }
  }
  
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  config.Port = port;
  config.DatabaseDriver = DBDriver
  return config, nil
}
