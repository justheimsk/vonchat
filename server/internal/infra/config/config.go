package config

import (
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	"github.com/spf13/viper"
)

type PostgresConfig struct {
	Host     string
	Port     string
	DB       string
	User     string
	Password string
}

type SqliteConfig struct {
	Path string
}

type Config struct {
	DatabaseDriver string
	Postgres       PostgresConfig
	Sqlite         SqliteConfig
	Port           string
	Debug          bool
}

func LoadConfig(log models.Logger) (*Config, error) {
	config := &Config{}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatalf("Failed to read configuration file: %w", err)
		}
	} else {
		log.Infof("Configuration file loaded: %s", viper.ConfigFileUsed())
	}

	viper.SetDefault("database.driver", "SQLITE")
	viper.SetDefault("debug", false)
	viper.SetDefault("port", "8080")

	config.Debug = viper.GetBool("debug")
	config.Port = viper.GetString("port")
	config.DatabaseDriver = viper.GetString("database.driver")

	if config.DatabaseDriver == "SQLITE" {
		path := viper.GetString("database.path")
		if path == "" {
			log.Fatalf("Using SQLITE database driver but missing database.path config.")
			return nil, models.ErrBadRequest
		} else {
			config.Sqlite.Path = path
		}
	} else if config.DatabaseDriver == "POSTGRES" {
		keys := []string{"host", "port", "user", "db", "password"}
		missing_config := false

		for _, key := range keys {
			value := viper.GetString("database." + key)

			if value == "" {
				log.Errorf("Using POSTGRES database driver but missing databse.%s config.", key)
				missing_config = true
				continue
			}

			switch key {
			case "host":
				config.Postgres.Host = value
			case "port":
				config.Postgres.Port = value
			case "db":
				config.Postgres.DB = value
			case "user":
				config.Postgres.User = value
			case "password":
				config.Postgres.Password = value
			}
		}

		if missing_config {
			return config, models.NewCustomError("malformed_config_input", "Broken config.")
		}
	} else {
		return config, models.NewCustomError("unknown_database_driver", "Unknown database driver: "+config.DatabaseDriver)
	}

	return config, nil
}
